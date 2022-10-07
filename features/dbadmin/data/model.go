package data

import (
	"rozhok/features/dbadmin"

	"gorm.io/gorm"
)

type TransaksiClient struct {
	gorm.Model
	ClientID      uint
	PorterID      uint
	TagihanID     uint
	Kurir         string
	TipeTransaksi string
	GrandTotal    int64
	Status        string
	KodeTransaksi string `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	Client                User                    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	DetailTransaksiClient []TransaksiClientDetail `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type User struct {
	gorm.Model
	Email           string `gorm:"unique"`
	Password        string
	Role            string
	Username        string
	StatusKemitraan string
	JunkStationName string
	ImageUrl        string
	Provinsi        string
	Kota            string
	Kecamatan       string
	Jalan           string
	Telepon         string
	Bonus           int64
	AlamatID        uint

	Alamat []Alamat `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type Alamat struct {
	gorm.Model
	Provinsi  string
	Kota      string
	Kecamatan string
	Jalan     string
	UserId    uint
	Status    string
}

type TransaksiClientDetail struct {
	gorm.Model
	TransaksiClientID uint
	KategoriID        uint
	ProdukID          uint
	Berat             int64
	Qty               uint
	Subtotal          int64

	Produk Produk
}

type Produk struct {
	gorm.Model
	Nama      string
	Image_url string
	Stok      int
	Harga     int64
	Desc      string
}

func ToCore(TransaksiClientModel TransaksiClient) dbadmin.TransaksiCore {
	transaksiClient := dbadmin.TransaksiCore{
		ID:            TransaksiClientModel.ID,
		GrandTotal:    TransaksiClientModel.GrandTotal,
		TipeTransaksi: TransaksiClientModel.TipeTransaksi,
		Status:        TransaksiClientModel.Status,
		KodeTransaksi: TransaksiClientModel.KodeTransaksi,
		Kurir:         TransaksiClientModel.Kurir,
	}

	var clientCore dbadmin.User
	if len(TransaksiClientModel.Client.Alamat) > 0 {
		clientCore = dbadmin.User{
			Name:      TransaksiClientModel.Client.Username,
			NoTelp:    TransaksiClientModel.Client.Telepon,
			Provinsi:  TransaksiClientModel.Client.Alamat[0].Provinsi,
			Kota:      TransaksiClientModel.Client.Alamat[0].Kota,
			Kecamatan: TransaksiClientModel.Client.Alamat[0].Kecamatan,
		}
	}

	productsCoreList := []dbadmin.Product{}
	for _, product := range TransaksiClientModel.DetailTransaksiClient {
		if product.Qty > 0 {
			productsCoreList = append(productsCoreList, dbadmin.Product{
				ImageUrl:    product.Produk.Image_url,
				ProductName: product.Produk.Nama,
				Qty:         product.Qty,
				Subtotal:    product.Subtotal,
			})
		}
	}

	transaksiClient.Product = productsCoreList
	transaksiClient.Client = clientCore

	return transaksiClient
}
