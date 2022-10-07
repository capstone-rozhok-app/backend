package data

import (
	"rozhok/features/payment"

	"gorm.io/gorm"
)

type TransaksiClient struct {
	gorm.Model
	ClientID              uint
	PorterID              uint
	TagihanID             uint
	Kurir                 string
	TipeTransaksi         string
	GrandTotal            int64
	Status                string
	KodeTransaksi         string
	DetailTransaksiClient []TransaksiClientDetail

	Tagihan Tagihan
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

type Tagihan struct {
	gorm.Model
	NoVa           string
	TipePembayaran string
	Bank           string
	GrandTotal     int64
}

type Cart struct {
	gorm.Model
	Subtotal  int64
	Qty       uint
	Checklist bool
	UserId    uint
	ProdukId  uint
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

	Alamat []Alamat
}

type Produk struct {
	gorm.Model
	Nama      string
	Image_url string
	Stok      int
	Harga     int64
	Desc      string
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

func ToCore(transaksiClient TransaksiClient) payment.Core {
	return payment.Core{
		Bank:           transaksiClient.Tagihan.Bank,
		NoVA:           transaksiClient.Tagihan.NoVa,
		TipePembayaran: transaksiClient.Tagihan.TipePembayaran,
		GrandTotal:     transaksiClient.Tagihan.GrandTotal,
		Kurir:          transaksiClient.Kurir,
	}
}
