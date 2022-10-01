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

	for _, model := range TransaksiPorterModel {
		rows = append(rows, toCore(model))
	}

	return rows, nil
}

func (repo *transaksiPorterRepo) Get(TransaksiCore transaksiporter.Core) (row transaksiporter.Core, err error) {
	var transaksiPorterModel TransaksiPorter
	transaksiPorterModel.ID = TransaksiCore.ID
	tx := repo.DB.Model(&TransaksiPorter{}).Preload("UserClient").Preload("TransaksiPorterDetail").First(&transaksiPorterModel)
	if tx.Error != nil {
		return row, tx.Error
	}

	return toCore(transaksiPorterModel), nil
}

func (repo *transaksiPorterRepo) CreateTransaksiPenjualan(TransaksiCore transaksiporter.Core) (row int, err error) {
	var transaksiPorterModel TransaksiPorter
	transaksiPorterModel.ID = TransaksiCore.ID

	tx := repo.DB.Model(&TransaksiPorter{}).Preload("TransaksiPorterDetail.KategoriRosok").First(&transaksiPorterModel)
	if tx.Error != nil {
		return row, tx.Error
	}

	var grandTotal int64
	for _, barangRosok := range transaksiPorterModel.TransaksiPorterDetail {
		grandTotal += int64(barangRosok.Berat) * barangRosok.KategoriRosok.HargaMitra
	}

	transaksiPorterModel.ID = 0
	transaksiPorterModel.GrandTotal = grandTotal
	transaksiPorterModel.TipeTransaksi = "penjualan"
	transaksiPorterModel.Status = "terjual"

	transaksiPorterModelBefore := transaksiPorterModel
	tx1 := repo.DB.Model(&TransaksiPorter{}).Create(&transaksiPorterModel)
	if tx1.Error != nil {
		return row, tx.Error
	}

	if tx1.RowsAffected < 1 {
		return int(tx1.RowsAffected), errors.New("failed to insert data")
	}

	barangRosokList := []TransaksiPorterDetail{}
	for _, barangRosok := range transaksiPorterModelBefore.TransaksiPorterDetail {
		barangRosokList = append(barangRosokList, TransaksiPorterDetail{
			TransaksiPorterID: transaksiPorterModel.ID,
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

func (repo *transaksiPorterRepo) UpdateTransaksiPembelian(TransaksiCore transaksiporter.Core) (row int, err error) {
	var transaksiPorterModel TransaksiPorter
	transaksiPorterModel.ID = TransaksiCore.ID

	//get transaksi porter sebelum dibayar dengan berelasi dengan barang rosok dan kategori rosok
	tx := repo.DB.Model(&TransaksiPorter{}).Preload("TransaksiPorterDetail.KategoriRosok").First(&transaksiPorterModel)
	if tx.Error != nil {
		return row, tx.Error
	}

	//looping untuk kalkulasi subtotal (harga kategori client * berat)
	barangRosokList := []TransaksiPorterDetail{}
	for _, barangRosok := range transaksiPorterModel.TransaksiPorterDetail {
		barangRosokList = append(barangRosokList, TransaksiPorterDetail{
			TransaksiPorterID: barangRosok.TransaksiPorterID,
			KategoriID:        barangRosok.KategoriID,
			Berat:             barangRosok.Berat,
			Subtotal:          barangRosok.KategoriRosok.HargaClient * int64(barangRosok.Berat),
		})
	}

	//delete barang rosok by transaksi porter id
	tx2 := repo.DB.Model(&TransaksiPorterDetail{}).Where("transaksi_porter_id = ?", transaksiPorterModel.ID).Delete(TransaksiPorterDetail{})
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

	return int(tx3.RowsAffected), nil
}
