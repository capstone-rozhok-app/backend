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
	var grandTotal int64
	for _, barangRosok := range transaksiPorterModel.TransaksiPorterDetail {
		grandTotal += int64(barangRosok.Berat) * barangRosok.KategoriRosok.HargaMitra
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
	tx := repo.DB.Model(&TransaksiPorter{}).Preload("TransaksiPorterDetail.KategoriRosok").First(&transaksiPorterModel)
	if tx.Error != nil {
		return row, tx.Error
	}

	if len(transaksiPorterModel.TransaksiPorterDetail) < 1 {
		return 0, errors.New("transaksi porter detail is empty")
	}

	//looping untuk kalkulasi subtotal (harga kategori client * berat)
	barangRosokList := []TransaksiPorterDetail{}
	var grandTotal int64
	for index, barangRosok := range transaksiPorterModel.TransaksiPorterDetail {
		barangRosokList = append(barangRosokList, TransaksiPorterDetail{
			TransaksiPorterID: barangRosok.TransaksiPorterID,
			KategoriID:        barangRosok.KategoriID,
			Berat:             uint(TransaksiCore.DetailTransaksiPorter[index].Berat),
			Subtotal:          barangRosok.KategoriRosok.HargaClient * int64(TransaksiCore.DetailTransaksiPorter[index].Berat),
		})
		grandTotal += barangRosok.KategoriRosok.HargaClient * int64(TransaksiCore.DetailTransaksiPorter[index].Berat)
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

	//buat ulang barang rosok by transaksi porter id
	transaksiPorterModel.GrandTotal = grandTotal
	transaksiPorterModel.Status = "dibayar"
	tx4 := repo.DB.Model(&TransaksiPorter{}).Where("id =  ?", transaksiPorterModel.ID).Updates(&transaksiPorterModel)
	if tx4.Error != nil {
		return row, tx.Error
	}

	if tx4.RowsAffected < 1 {
		return int(tx4.RowsAffected), errors.New("failed to insert data")
	}

	// update transaksi client dengan status sudah_bayar
	txTransaksiClient := repo.DB.Model(&TransaksiClient{}).Where("id = ?", transaksiPorterModel.TransaksiClientID).Update("status", "dibayar")
	if txTransaksiClient.Error != nil {
		return row, txTransaksiClient.Error
	}

	return int(tx4.RowsAffected), nil
}
