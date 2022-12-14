package data

import (
	"errors"
	pengambilanrosok "rozhok/features/pengambilan_rosok"

	"gorm.io/gorm"
)

type pengambilanRosokRepo struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *pengambilanRosokRepo {
	return &pengambilanRosokRepo{
		DB: db,
	}
}

func (repo *pengambilanRosokRepo) GetAll(TransaksiCore pengambilanrosok.Core) (rows []pengambilanrosok.Core, err error) {
	var pengambilanRosokModelList []TransaksiClient

	//ambil data porter
	porter := User{}
	if porerError := repo.DB.Model(&User{}).Where("id = ?", TransaksiCore.PorterID).First(&porter).Error; porerError != nil {
		return rows, porerError
	}

	//ambil data transaksi client dengan status belum_bayar dan kecamatan,kota,provinsi sesuai porter
	clientTransactionError := repo.DB.Model(&TransaksiClient{}).Where("tipe_transaksi = ?", "penjualan").Where("status = ?", "belum_bayar").Preload("UserClient.Alamat", func(db *gorm.DB) *gorm.DB {
		return db.Where("alamats.status", "utama").Where("alamats.provinsi", porter.Provinsi).Where("alamats.kota", porter.Kota).Where("alamats.kecamatan", porter.Kecamatan)
	}).Find(&pengambilanRosokModelList).Error

	if clientTransactionError != nil {
		return rows, clientTransactionError
	}

	var pengambilanRosokCoreList []pengambilanrosok.Core
	for _, transaksiRosokClient := range pengambilanRosokModelList {
		if len(transaksiRosokClient.UserClient.Alamat) > 0 {
			pengambilanRosokCoreList = append(pengambilanRosokCoreList, toCore(transaksiRosokClient))
		}
	}

	return pengambilanRosokCoreList, nil
}

func (repo *pengambilanRosokRepo) Get(TransaksiCore pengambilanrosok.Core) (row pengambilanrosok.Core, err error) {
	var transaksiClientModel TransaksiClient
	transaksiClientModel.ID = TransaksiCore.ID

	if transaksiClientError := repo.DB.Model(&TransaksiClient{}).Preload("UserClient.Alamat", func(db *gorm.DB) *gorm.DB {
		return db.Where("alamats.status", "utama")
	}).Preload("TransaksiClientDetail.KategoriRosok").First(&transaksiClientModel).Error; transaksiClientError != nil {
		return row, transaksiClientError
	}

	return toCore(transaksiClientModel), nil
}

// pengambilan rosok dari transaksi client ke dalam transaksi porter
func (repo *pengambilanRosokRepo) CreatePengambilanRosok(TransaksiCore pengambilanrosok.Core) (row int, err error) {
	// ambil data transaksi client
	var transaksiClientModel TransaksiClient
	transaksiClientModel.ID = TransaksiCore.ID

	if transaksiClientError := repo.DB.Model(&TransaksiClient{}).Preload("UserClient").Preload("TransaksiClientDetail.KategoriRosok").First(&transaksiClientModel).Error; transaksiClientError != nil {
		return row, transaksiClientError
	}

	// masukkan data transaksi client ke dalam transaksi porter dengan status "unpaid"
	transaksiPorterModel := TransaksiPorter{
		ClientID:          transaksiClientModel.ClientID,
		PorterID:          TransaksiCore.PorterID,
		TransaksiClientID: transaksiClientModel.ID,
		TipeTransaksi:     "pembelian",
		Status:            "belum_bayar",
	}

	txTransaksiPorter := repo.DB.Model(&TransaksiPorter{}).Create(&transaksiPorterModel)
	if txTransaksiPorter.Error != nil {
		return row, txTransaksiPorter.Error
	}

	if txTransaksiPorter.RowsAffected < 1 {
		return row, errors.New("failed create to transaksi porter")
	}

	transaksiPorterModelDetailList := []TransaksiPorterDetail{}
	for _, barangRosok := range transaksiClientModel.TransaksiClientDetail {
		transaksiPorterModelDetailList = append(transaksiPorterModelDetailList, TransaksiPorterDetail{
			TransaksiPorterID: transaksiPorterModel.ID,
			KategoriID:        barangRosok.KategoriID,
		})
	}

	txTransaksiPorterDetail := repo.DB.Model(&TransaksiPorterDetail{}).Create(&transaksiPorterModelDetailList)
	if txTransaksiPorterDetail.Error != nil {
		return row, txTransaksiPorterDetail.Error
	}

	if txTransaksiPorterDetail.RowsAffected < 1 {
		return row, errors.New("failed create to transaksi porter detail")
	}

	// update data transaksi client "pending"
	txTransaksiClient := repo.DB.Model(&TransaksiClient{}).Where("id = ?", TransaksiCore.ID).Update("status", "dalam_perjalanan")
	if txTransaksiClient.Error != nil {
		return row, txTransaksiClient.Error
	}

	if txTransaksiClient.RowsAffected < 1 {
		return row, errors.New("failed update to client transaction")
	}

	return 1, nil
}
