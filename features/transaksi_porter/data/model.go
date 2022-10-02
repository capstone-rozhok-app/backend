package data

import (
	transaksiporter "rozhok/features/transaksi_porter"

	"gorm.io/gorm"
)

type TransaksiPorter struct {
	gorm.Model
	PorterID      uint
	ClientID      uint
	TipeTransaksi string
	GrandTotal    int64
	Status        string

	UserClient            User `gorm:"foreignKey:ClientID"`
	TransaksiPorterDetail []TransaksiPorterDetail
}

type TransaksiPorterDetail struct {
	gorm.Model
	TransaksiPorterID uint
	KategoriID        uint
	Berat             uint
	Subtotal          int64

	KategoriRosok   KategoriRosok   `gorm:"foreignKey:KategoriID"`
	TransaksiPorter TransaksiPorter `gorm:"foreignKey:TransaksiPorterID"`
}

type KategoriRosok struct {
	gorm.Model
	NamaKategori string
	HargaMitra   int64
	HargaClient  int64
	Desc         string
}

type User struct {
	gorm.Model
	Email           string
	Password        string
	Role            string
	Username        string
	StatusKemitraan string
	JunkStationName string
	Foto            string
	Provinsi        string
	Kota            string
	Kecamatan       string
	Jalan           string
	Telepon         string
	Bonus           int64
}

func FromCore(transaksiPorterCore transaksiporter.Core) TransaksiPorter {
	transaksiPorterModel := TransaksiPorter{
		ClientID:      transaksiPorterCore.Client.ID,
		PorterID:      transaksiPorterCore.PorterID,
		TipeTransaksi: transaksiPorterCore.TipeTransaksi,
		GrandTotal:    transaksiPorterCore.GrandTotal,
		Status:        transaksiPorterCore.Status,
	}
	transaksiPorterModel.ID = transaksiPorterCore.ID

	transaksiPorterDetailModelList := []TransaksiPorterDetail{}
	for _, detailTransaksiCore := range transaksiPorterCore.DetailTransaksiPorter {
		barangRosok := TransaksiPorterDetail{
			Berat:    uint(detailTransaksiCore.Berat),
			Subtotal: detailTransaksiCore.Subtotal,
		}
		barangRosok.ID = detailTransaksiCore.Id
		transaksiPorterDetailModelList = append(transaksiPorterDetailModelList, barangRosok)
	}

	transaksiPorterModel.TransaksiPorterDetail = transaksiPorterDetailModelList

	return transaksiPorterModel
}

func toCore(transaksiPorterModel TransaksiPorter) transaksiporter.Core {
	transaksiPorterCore := transaksiporter.Core{
		PorterID:      transaksiPorterModel.PorterID,
		TipeTransaksi: transaksiPorterModel.TipeTransaksi,
		Status:        transaksiPorterModel.Status,
		GrandTotal:    transaksiPorterModel.GrandTotal,
		Client: transaksiporter.User{
			Username:  transaksiPorterModel.UserClient.Username,
			Provinsi:  transaksiPorterModel.UserClient.Provinsi,
			Kota:      transaksiPorterModel.UserClient.Kota,
			Jalan:     transaksiPorterModel.UserClient.Jalan,
			Telepon:   transaksiPorterModel.UserClient.Telepon,
			Kecamatan: transaksiPorterModel.UserClient.Kecamatan,
		},
	}
	transaksiPorterCore.ID = transaksiPorterModel.ID

	transaksiDetailPorterCoreList := []transaksiporter.DetailTransaksiPorter{}
	for _, detailTransaksiModel := range transaksiPorterModel.TransaksiPorterDetail {
		barangRosok := transaksiporter.DetailTransaksiPorter{
			Nama:          detailTransaksiModel.KategoriRosok.NamaKategori,
			Berat:         int64(detailTransaksiModel.Berat),
			Subtotal:      detailTransaksiModel.Subtotal,
			HargaKategori: detailTransaksiModel.KategoriRosok.HargaClient,
		}
		barangRosok.Id = detailTransaksiModel.ID
		transaksiDetailPorterCoreList = append(transaksiDetailPorterCoreList, barangRosok)
	}

	transaksiPorterCore.DetailTransaksiPorter = transaksiDetailPorterCoreList

	return transaksiPorterCore
}
