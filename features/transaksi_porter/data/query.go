package data

import (
	"errors"
	transaksiporter "rozhok/features/transaksi_porter"

	"gorm.io/gorm"
)

type transaksiPorterRepo struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *transaksiPorterRepo {
	return &transaksiPorterRepo{
		DB: db,
	}
}

func (repo *transaksiPorterRepo) GetAll(TransaksiCore transaksiporter.Core) (rows []transaksiporter.Core, err error) {
	var TransaksiPorterModel []TransaksiPorter

	tx := repo.DB.Model(&TransaksiPorter{}).Where("porter_id = ?", TransaksiCore.PorterID)

	if TransaksiCore.StartDate != "" {
		tx.Where("created_at >= ?", TransaksiCore.StartDate)
	}

	if TransaksiCore.EndDate != "" {
		tx.Where("created_at <= ?", TransaksiCore.EndDate)
	}

	if TransaksiCore.TipeTransaksi != "" {
		tx.Where("tipe_transaksi = ?", TransaksiCore.TipeTransaksi)
	}

	if TransaksiCore.Status != "" {
		tx.Where("status = ?", TransaksiCore.Status)
	}

	tx.Find(&TransaksiPorterModel)
	if tx.Error != nil {
		return rows, tx.Error
	}

	if len(TransaksiPorterModel) < 1 {
		return rows, errors.New("not found")
	}

	for _, model := range TransaksiPorterModel {
		rows = append(rows, toCore(model))
	}

	return rows, nil
}

func (repo *transaksiPorterRepo) Get(TransaksiCore transaksiporter.Core) (row transaksiporter.Core, err error) {
	var transaksiPorterModel TransaksiPorter
	transaksiPorterModel.ID = TransaksiCore.ID
	tx := repo.DB.Model(&TransaksiPorter{}).Preload("UserClient.Alamat", func(db *gorm.DB) *gorm.DB {
		return db.Where("alamats.status", "utama")
	}).Preload("TransaksiPorterDetail.KategoriRosok").Where("tipe_transaksi = ?", TransaksiCore.TipeTransaksi).First(&transaksiPorterModel)
	if tx.Error != nil {
		return row, tx.Error
	}

	return toCore(transaksiPorterModel), nil
}

// jual rosok ke junk-station
func (repo *transaksiPorterRepo) CreateTransaksiPenjualan(TransaksiCore transaksiporter.Core) (row int, err error) {
	var transaksiPorterModel TransaksiPorter

	transaksiPorterModel.ID = TransaksiCore.ID

	// ambil data transaksi porter by id transaksi
	tx := repo.DB.Model(&TransaksiPorter{}).Preload("TransaksiPorterDetail.KategoriRosok").First(&transaksiPorterModel)
	if tx.Error != nil {
		return row, tx.Error
	}

	if transaksiPorterModel.Status != "dibayar" {
		return 0, errors.New("failed status not sudah_dibayar")
	}

	// kalkulasi harga mitra
	var grandTotal float64
	for _, barangRosok := range transaksiPorterModel.TransaksiPorterDetail {
		grandTotal += float64(barangRosok.Berat) * float64(barangRosok.KategoriRosok.HargaMitra)
	}

	// buat ulang data transaksi kemudian ganti dengan status terjual
	transaksiPorterModelAfter := TransaksiPorter{
		PorterID:      transaksiPorterModel.PorterID,
		ClientID:      transaksiPorterModel.ClientID,
		GrandTotal:    grandTotal,
		TipeTransaksi: "penjualan",
		Status:        "terjual",
	}

	tx1 := repo.DB.Model(&TransaksiPorter{}).Create(&transaksiPorterModelAfter)
	if tx1.Error != nil {
		return row, tx.Error
	}

	if tx1.RowsAffected < 1 {
		return int(tx1.RowsAffected), errors.New("failed to insert data")
	}

	// insert ulang juga data detail transaksi yang berkaitan, dan kalkulasi ulang dengan harga mitra.
	barangRosokList := []TransaksiPorterDetail{}
	for _, barangRosok := range transaksiPorterModel.TransaksiPorterDetail {
		barangRosokList = append(barangRosokList, TransaksiPorterDetail{
			TransaksiPorterID: transaksiPorterModelAfter.ID,
			KategoriID:        barangRosok.KategoriID,
			Berat:             barangRosok.Berat,
			Subtotal:          barangRosok.KategoriRosok.HargaMitra * int64(barangRosok.Berat),
		})
	}

	tx3 := repo.DB.Model(&TransaksiPorterDetail{}).Create(&barangRosokList)
	if tx3.Error != nil {
		return row, tx.Error
	}

	if tx3.RowsAffected < 1 {
		return int(tx3.RowsAffected), errors.New("failed to insert data")
	}

	return int(tx3.RowsAffected), nil
}

// update transaksi porter input berat dan harga ke transaksi porter
func (repo *transaksiPorterRepo) UpdateTransaksiPembelian(TransaksiCore transaksiporter.Core) (row int, err error) {
	var transaksiPorterModel TransaksiPorter
	transaksiPorterModel.ID = TransaksiCore.ID

	//get transaksi porter sebelum dibayar dengan berelasi dengan barang rosok dan kategori rosok
	tx := repo.DB.Model(&TransaksiPorter{}).Preload("TransaksiPorterDetail.KategoriRosok").Where("status", "belum_bayar").First(&transaksiPorterModel)
	if tx.Error != nil {
		return row, tx.Error
	}

	if len(transaksiPorterModel.TransaksiPorterDetail) < 1 {
		return 0, errors.New("transaksi porter detail is empty")
	}

	//looping untuk kalkulasi subtotal (harga kategori client * berat)
	barangRosokList := []TransaksiPorterDetail{}
	var grandTotal float64
	var totalBerat int64
	for index, barangRosok := range transaksiPorterModel.TransaksiPorterDetail {
		barangRosokList = append(barangRosokList, TransaksiPorterDetail{
			TransaksiPorterID: barangRosok.TransaksiPorterID,
			KategoriID:        barangRosok.KategoriID,
			Berat:             uint(TransaksiCore.DetailTransaksiPorter[index].Berat),
			Subtotal:          barangRosok.KategoriRosok.HargaClient * int64(TransaksiCore.DetailTransaksiPorter[index].Berat),
		})
		totalBerat += TransaksiCore.DetailTransaksiPorter[index].Berat
		grandTotal += float64(barangRosok.KategoriRosok.HargaClient) * float64(TransaksiCore.DetailTransaksiPorter[index].Berat)
	}

	//delete barang rosok by transaksi porter id
	tx2 := repo.DB.Model(&TransaksiPorterDetail{}).Where("transaksi_porter_id = ?", transaksiPorterModel.ID).Unscoped().Delete(barangRosokList)
	if tx2.Error != nil {
		return row, tx.Error
	}

	//buat ulang barang rosok by transaksi porter id
	tx3 := repo.DB.Model(&TransaksiPorterDetail{}).Create(&barangRosokList)
	if tx3.Error != nil {
		return row, tx.Error
	}

	if tx3.RowsAffected < 1 {
		return int(tx3.RowsAffected), errors.New("failed to insert data")
	}

	// get client dan ambil bonus
	var Client User

	txClient := repo.DB.Model(&User{}).Where("id = ?", transaksiPorterModel.ClientID).First(&Client)
	if txClient.Error != nil {
		return row, tx.Error
	}

	// check bonus jika lebih dari 100kg maka reset dari 0 dan tambahkan bonus 5% dari harga awal
	prebonus := Client.Bonus + totalBerat
	Client.Bonus = prebonus
	var bonusHarga float64
	if prebonus >= 100 {
		// jika prebonus melebihi 100
		Client.Bonus = prebonus % 100
		bonusHarga = float64(prebonus/100) * float64(grandTotal) * 0.5

		// insert bonus ke log bonus
		logBonus := LogBonus{
			PorterID:   transaksiPorterModel.PorterID,
			ClientID:   transaksiPorterModel.ClientID,
			TotalBerat: totalBerat,
			BonusHarga: bonusHarga,
		}
		txLog := repo.DB.Model(&LogBonus{}).Create(&logBonus)
		if txLog.Error != nil {
			return 0, txLog.Error
		}
	}

	// update bonus di client
	txUpdate := repo.DB.Model(&User{}).Select("bonus").Where("id = ?", transaksiPorterModel.ClientID).Update("bonus", Client.Bonus)
	if txUpdate.Error != nil {
		return 0, txUpdate.Error
	}

	if txUpdate.RowsAffected < 1 {
		return 0, errors.New("error affected row")
	}

	//update transaksi porter
	transaksiPorterModel.GrandTotal = grandTotal + bonusHarga
	transaksiPorterModel.Status = "dibayar"
	tx4 := repo.DB.Model(&TransaksiPorter{}).Where("id =  ?", transaksiPorterModel.ID).Updates(&transaksiPorterModel)
	if tx4.Error != nil {
		return row, tx.Error
	}

	if tx4.RowsAffected < 1 {
		return int(tx4.RowsAffected), errors.New("failed to insert data")
	}

	// update transaksi client
	transaksiClient := TransaksiClient{
		GrandTotal: transaksiPorterModel.GrandTotal,
		Status:     "dibayar",
	}
	txTransaksiClient := repo.DB.Model(&TransaksiClient{}).Where("id = ?", transaksiPorterModel.TransaksiClientID).Select("status", "grand_total").Updates(&transaksiClient)
	if txTransaksiClient.Error != nil {
		return row, txTransaksiClient.Error
	}

	// hapus semua transaksi detail client
	var TransaksiClientDetailModelList []TransaksiClientDetail
	txDelTransaksiClient := repo.DB.Model(&TransaksiClientDetail{}).Where("transaksi_client_id", transaksiPorterModel.TransaksiClientID).Unscoped().Delete(&TransaksiClientDetailModelList)
	if txDelTransaksiClient.Error != nil {
		return 0, txDelTransaksiClient.Error
	}

	// create transaksi detail client
	var NewTransaksiDetailList []TransaksiClientDetail
	for _, barangrosok := range barangRosokList {
		NewTransaksiDetailList = append(NewTransaksiDetailList, TransaksiClientDetail{
			TransaksiClientID: transaksiPorterModel.TransaksiClientID,
			KategoriID:        barangrosok.KategoriID,
			Berat:             int64(barangrosok.Berat),
			Subtotal:          barangrosok.Subtotal,
		})
	}

	txCreateTransaksiClient := repo.DB.Model(&TransaksiClientDetail{}).Omit("produk_id").Create(&NewTransaksiDetailList)
	if txCreateTransaksiClient.Error != nil {
		return 0, txCreateTransaksiClient.Error
	}

	return int(txCreateTransaksiClient.RowsAffected), nil
}
