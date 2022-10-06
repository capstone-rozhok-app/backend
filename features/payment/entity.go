package payment

import "time"

type Core struct {
	IdTransaksi    uint
	Bank           string
	Kurir          string
	NoVA           string
	TipePembayaran string
	GrandTotal     int64
	ExpiredAt      time.Time

	Client Client
}

type Client struct {
	ID        uint
	Username  string
	Provinsi  string
	Kota      string
	Kecamatan string
}

type PaymentData interface {
	GetUserData(PaymentCore Core) (Core, error)
	GetTagihan(idTransaksi uint) (Core, error)
	Insert(PaymentData Core) (idTransaksi uint, err error)
}

type PaymentUsecase interface {
	Create(PaymentData Core) (Core, error)
}
