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

func ToCore(transaksiClient TransaksiClient) payment.Core {
	return payment.Core{
		Bank:           transaksiClient.Tagihan.Bank,
		NoVA:           transaksiClient.Tagihan.NoVa,
		TipePembayaran: transaksiClient.Tagihan.TipePembayaran,
		GrandTotal:     transaksiClient.Tagihan.GrandTotal,
		Kurir:          transaksiClient.Kurir,
	}
}
