package data

import (
	pengambilanrosok "rozhok/features/pengambilan_rosok"

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

	KategoriRosok   KategoriRosok   `gorm:"constraint:OnUpdate:SET NULL,OnDelete:SET NULL;foreignKey:KategoriID"`
	TransaksiPorter TransaksiPorter `gorm:"foreignKey:TransaksiPorterID"`
}

type TransaksiClient struct {
	gorm.Model
	PorterID      uint
	ClientID      uint
	TagihanID     uint
	Kurir         string
	TipeTransaksi string
	GrandTotal    int64
	Status        string

	UserClient            User `gorm:"foreignKey:ClientID"`
	TransaksiClientDetail []TransaksiClientDetail
}

type TransaksiClientDetail struct {
	gorm.Model
	TransaksiClientID uint
	KategoriID        uint
	Berat             uint
	Subtotal          int64

	KategoriRosok   KategoriRosok   `gorm:"constraint:OnUpdate:SET NULL,OnDelete:SET NULL;foreignKey:KategoriID"`
	TransaksiClient TransaksiClient `gorm:"foreignKey:TransaksiClientID"`
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
	Alamat          []Alamat
}

type Alamat struct {
	gorm.Model
	UserID    uint
	Provinsi  string
	Kota      string
	Kecamatan string
	Status    string
	Jalan     string

	User User `gorm:"foreignKey:UserID"`
}

func FromCore(transaksiClientCore pengambilanrosok.Core) TransaksiClient {
	transaksiClientModel := TransaksiClient{
		ClientID: transaksiClientCore.Client.ID,
		PorterID: transaksiClientCore.PorterID,
		Status:   transaksiClientCore.Status,
	}
	transaksiClientModel.ID = transaksiClientCore.ID

	transaksiClientDetailModelList := []TransaksiClientDetail{}
	for _, detailTransaksiCore := range transaksiClientCore.DetailTransaksiClient {
		barangRosok := TransaksiClientDetail{
			TransaksiClientID: transaksiClientModel.ID,
			KategoriID:        detailTransaksiCore.IdKategori,
		}
		barangRosok.ID = detailTransaksiCore.Id
		transaksiClientDetailModelList = append(transaksiClientDetailModelList, barangRosok)
	}

	transaksiClientModel.TransaksiClientDetail = transaksiClientDetailModelList

	return transaksiClientModel
}

func toCore(transaksiClientModel TransaksiClient) pengambilanrosok.Core {
	transaksiClientCore := pengambilanrosok.Core{
		PorterID:      transaksiClientModel.PorterID,
		TipeTransaksi: transaksiClientModel.TipeTransaksi,
		Status:        transaksiClientModel.Status,
		Client: pengambilanrosok.User{
			Username:  transaksiClientModel.UserClient.Username,
			Provinsi:  transaksiClientModel.UserClient.Alamat[0].Provinsi,
			Kota:      transaksiClientModel.UserClient.Alamat[0].Kota,
			Jalan:     transaksiClientModel.UserClient.Alamat[0].Jalan,
			Telepon:   transaksiClientModel.UserClient.Telepon,
			Kecamatan: transaksiClientModel.UserClient.Alamat[0].Kecamatan,
		},
	}

	transaksiClientCoreDetail := []pengambilanrosok.DetailTransaksiClient{}
	for _, barangRosok := range transaksiClientModel.TransaksiClientDetail {
		transaksiClientCoreDetail = append(transaksiClientCoreDetail, pengambilanrosok.DetailTransaksiClient{
			Id:         barangRosok.ID,
			IdKategori: barangRosok.KategoriID,
			Nama:       barangRosok.KategoriRosok.NamaKategori,
		})
	}

	transaksiClientCore.ID = transaksiClientModel.ID
	transaksiClientCore.DetailTransaksiClient = transaksiClientCoreDetail
	return transaksiClientCore
}
