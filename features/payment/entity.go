package payment

import "time"

type Core struct {
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
	Insert(PaymentData Core) (int, error)
}

type PaymentUsecase interface {
	Create(PaymentData Core) (int, error)
}
