package data

import (
	"errors"
	"rozhok/features/dbadmin"

	"gorm.io/gorm"
)

type data struct {
	db *gorm.DB
}

func New(db *gorm.DB) dbadmin.DataInterface {
	return &data{
		db: db,
	}
}

func (repo *data) GetUsers() (dbadmin.ResponseCore, error) {
	var userCores dbadmin.ResponseCore

	var client int64
	repo.db.Model(&User{}).Where("role = ?", "client").Count(&client)
	var mitra int64
	repo.db.Model(&User{}).Where("role = ?", "junk_station").Count(&mitra)
	userCores.TotalCL = int(client)
	userCores.TotalJS = int(mitra)

	var sliceGrafik []dbadmin.GrafikData
	repo.db.Model(&User{}).Select("month(created_at) as bulan, count(id) as jumlah_cl").Where("role = ?", "client").Group("month(created_at)").Find(&sliceGrafik)
	var sliceGrafik2 []dbadmin.GrafikData
	repo.db.Model(&User{}).Select("month(created_at) as bulan, count(id) as jumlah_js").Where("role = ?", "junk_station").Group("month(created_at)").Find(&sliceGrafik2)

	for i := range sliceGrafik {
		for j := range sliceGrafik2 {
			if sliceGrafik[i].Bulan == sliceGrafik2[j].Bulan {
				sliceGrafik[i].JumlahJS = sliceGrafik2[j].JumlahJS
			}
		}
	}

	userCores.Grafik = sliceGrafik

	return userCores, nil
}

func (r *data) GetTransaksiDetail(TransaksiData dbadmin.TransaksiCore) (dbadmin.TransaksiCore, error) {
	transaksiClientModel := TransaksiClient{}
	transaksiClientModel.ID = TransaksiData.ID

	tx := r.db.Model(&TransaksiClient{}).Preload("Client.Alamat", func(db *gorm.DB) *gorm.DB {
		return db.Where("alamats.status", "utama")
	})
	tx.Preload("DetailTransaksiClient.Produk")

	tx.First(&transaksiClientModel)

	if tx.Error != nil {
		return dbadmin.TransaksiCore{}, tx.Error
	}

	return ToCore(transaksiClientModel), nil
}

func (r *data) GetTransaksi(TransaksiData dbadmin.TransaksiCore) ([]dbadmin.TransaksiCore, error) {
	transaksiClientModelList := []TransaksiClient{}

	tx := r.db.Model(&TransaksiClient{})

	if TransaksiData.StartDate != "" {
		tx.Where("created_at >=", TransaksiData.StartDate)
	}

	if TransaksiData.EndDate != "" {
		tx.Where("created_at <=", TransaksiData.EndDate)
	}

	if TransaksiData.TipeTransaksi != "" {
		tx.Where("tipe_transaksi = ?", TransaksiData.TipeTransaksi)
	}

	tx.Find(&transaksiClientModelList)

	if tx.Error != nil {
		return []dbadmin.TransaksiCore{}, tx.Error
	}

	transaksiClientCoreList := []dbadmin.TransaksiCore{}
	for _, transaksiClient := range transaksiClientModelList {
		transaksiClientCoreList = append(transaksiClientCoreList, ToCore(transaksiClient))
	}

	return transaksiClientCoreList, nil
}

func (repo *data) UpdateTransaksi(TransaksiData dbadmin.TransaksiCore) error {
	tx := repo.db.Model(&TransaksiClient{}).Where("id = ?", TransaksiData.ID).Where("status = ?", "dibayar").Update("status", "dikirim")
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected < 1 {
		return errors.New("error affected row")
	}

	return nil
}
