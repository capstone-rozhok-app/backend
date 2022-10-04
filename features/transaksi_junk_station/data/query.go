package data

import (
	"errors"
	transaksijunkstation "rozhok/features/transaksi_junk_station"
	"rozhok/utils/helper"

	"gorm.io/gorm"
)

type TransaksiJunkStationRepo struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *TransaksiJunkStationRepo {
	return &TransaksiJunkStationRepo{
		DB: db,
	}
}

func (r *TransaksiJunkStationRepo) GetAll(TransaksiJunkStationCore transaksijunkstation.Core) ([]transaksijunkstation.Core, error) {
	transaksiModelList := []TransaksiJunkStation{}

	tx := r.DB.Model(&TransaksiJunkStation{}).Where("user_id = ?", TransaksiJunkStationCore.UserID).Preload("TransaksiJunkStationDetail.KategoriRosok")

	if TransaksiJunkStationCore.StartDate != "" {
		tx.Where("created_at >=", TransaksiJunkStationCore.StartDate)
	}

	if TransaksiJunkStationCore.EndDate != "" {
		tx.Where("created_at <=", TransaksiJunkStationCore.EndDate)
	}

	tx.Find(&transaksiModelList)

	if tx.Error != nil {
		return []transaksijunkstation.Core{}, tx.Error
	}

	transaksiJunkStationCoreList := []transaksijunkstation.Core{}
	for _, transaksi := range transaksiModelList {
		transaksiJunkStationCoreList = append(transaksiJunkStationCoreList, ToCore(transaksi))
	}

	return transaksiJunkStationCoreList, nil
}

func (r *TransaksiJunkStationRepo) Get(TransaksiJunkStationCore transaksijunkstation.Core) (transaksijunkstation.Core, error) {
	transaksiModel := TransaksiJunkStation{}
	transaksiModel.ID = TransaksiJunkStationCore.ID

	tx := r.DB.Model(&TransaksiJunkStation{}).Preload("TransaksiJunkStationDetail.KategoriRosok").First(&transaksiModel)
	if tx.Error != nil {
		return transaksijunkstation.Core{}, tx.Error
	}

	return ToCore(transaksiModel), nil
}

func (r *TransaksiJunkStationRepo) Insert(TransaksiJunkStationCore transaksijunkstation.Core) (int, error) {
	// ambil data dari keranjang rosok
	KeranjangRosokList := []KeranjangRosok{}
	txKeranjangRosok := r.DB.Model(&KeranjangRosok{}).Where("client_id = ?", TransaksiJunkStationCore.UserID).Find(&KeranjangRosokList)
	if txKeranjangRosok.Error != nil {
		return 0, errors.New("errors get data keranjang rosok")
	}

	//loop untuk mendapatkan grandtotal
	var grandTotal int64
	for _, keranjangrosok := range KeranjangRosokList {
		grandTotal += keranjangrosok.Subtotal
	}

	// insert data transaksi
	transaksiModel := TransaksiJunkStation{
		UserID:     TransaksiJunkStationCore.UserID,
		GrandTotal: grandTotal,
		KodeTf:     helper.GenerateTF(int(TransaksiJunkStationCore.UserID)),
	}

	txTransaksi := r.DB.Model(&TransaksiJunkStation{}).Create(&transaksiModel)
	if txTransaksi.Error != nil {
		return 0, txTransaksi.Error
	}

	if txTransaksi.RowsAffected < 1 {
		return 0, errors.New("error rows affected")
	}

	// mapping data keranjang rosok ke detail transaksi
	transaksiModelDetail := []TransaksiJunkStationDetail{}
	for _, transaksi := range KeranjangRosokList {
		transaksiModel := TransaksiJunkStationDetail{
			TransaksiJunkStationID: transaksiModel.ID,
			KategoriRosokID:        transaksi.KategoriRosokID,
			Berat:                  transaksi.Berat,
			Subtotal:               transaksi.Subtotal,
		}
		transaksiModelDetail = append(transaksiModelDetail, transaksiModel)
	}

	// insert data detail transaksi
	txTransaksiDetail := r.DB.Model(&TransaksiJunkStationDetail{}).Create(&transaksiModelDetail)
	if txTransaksiDetail.Error != nil {
		return 0, txTransaksiDetail.Error
	}
	if txTransaksiDetail.RowsAffected < 1 {
		return 0, errors.New("error rows affected")
	}

	// hapus semua data di keranjang rosok berdasarkan user
	keranjangRosok := KeranjangRosok{}
	txDeleteKeranjangRosok := r.DB.Model(&KeranjangRosok{}).Where("client_id = ?", TransaksiJunkStationCore.UserID).Delete(&keranjangRosok)
	if txTransaksiDetail != nil {
		return 0, txDeleteKeranjangRosok.Error
	}
	if txDeleteKeranjangRosok.RowsAffected < 1 {
		return 0, errors.New("error rows affected")
	}

	return 1, nil
}
