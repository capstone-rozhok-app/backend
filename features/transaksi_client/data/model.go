package data

import (
	transaksiclient "rozhok/features/transaksi_client"

	"gorm.io/gorm"
)

type KeranjangRosok struct {
	gorm.Model
	ClientID        uint
	KategoriRosokID uint
	Berat           int
	Subtotal        int64

	KategoriRosok KategoriRosok `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

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
	Porter                User                    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Tagihan               Tagihan                 `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
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

type Tagihan struct {
	gorm.Model
	NoVa           string
	TipePembayaran string
	Bank           string
	GrandTotal     int64
}

type TransaksiClientDetail struct {
	gorm.Model
	TransaksiClientID uint
	KategoriID        uint
	ProdukID          uint
	Berat             int64
	Qty               uint
	Subtotal          int64

	Produk        Produk
	KategoriRosok KategoriRosok `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:KategoriID"`
}

type Produk struct {
	gorm.Model
	Nama      string
	Image_url string
	Stok      int
	Harga     int64
	Descr     string
}

type KategoriRosok struct {
	gorm.Model
	NamaKategori string
	HargaMitra   int64
	HargaClient  int64
	Desc         string
}

func ToCore(TransaksiClientModel TransaksiClient) transaksiclient.Core {
	transaksiClient := transaksiclient.Core{
		ID:            TransaksiClientModel.ID,
		GrandTotal:    TransaksiClientModel.GrandTotal,
		TipeTransaksi: TransaksiClientModel.TipeTransaksi,
		Status:        TransaksiClientModel.Status,
		KodeTransaksi: TransaksiClientModel.KodeTransaksi,
		Kurir:         TransaksiClientModel.Kurir,
		Porter: transaksiclient.User{
			Name:   TransaksiClientModel.Porter.Telepon,
			NoTelp: TransaksiClientModel.Porter.Telepon,
		},
	}

	var clientCore transaksiclient.User
	if len(TransaksiClientModel.Client.Alamat) > 0 {
		clientCore = transaksiclient.User{
			Name:      TransaksiClientModel.Client.Username,
			NoTelp:    TransaksiClientModel.Client.Telepon,
			Provinsi:  TransaksiClientModel.Client.Alamat[0].Provinsi,
			Kota:      TransaksiClientModel.Client.Alamat[0].Kota,
			Kecamatan: TransaksiClientModel.Client.Alamat[0].Kecamatan,
		}
	}

	productsCoreList := []transaksiclient.Product{}
	for _, product := range TransaksiClientModel.DetailTransaksiClient {
		if product.Qty > 0 {
			productsCoreList = append(productsCoreList, transaksiclient.Product{
				ImageUrl:    product.Produk.Image_url,
				ProductName: product.Produk.Nama,
				Qty:         product.Qty,
				Subtotal:    product.Subtotal,
			})
		}
	}

	barangRosokCoreList := []transaksiclient.BarangRosok{}
	for _, barangRosok := range TransaksiClientModel.DetailTransaksiClient {
		if barangRosok.KategoriID > 0 {
			barangRosokCoreList = append(barangRosokCoreList, transaksiclient.BarangRosok{
				Kategori: barangRosok.KategoriRosok.NamaKategori,
				Berat:    barangRosok.Berat,
				Harga:    barangRosok.Subtotal,
			})
		}
	}

	transaksiClient.Product = productsCoreList
	transaksiClient.BarangRosok = barangRosokCoreList
	transaksiClient.Client = clientCore

	return transaksiClient
}
