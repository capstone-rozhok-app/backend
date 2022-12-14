package payment

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type Core struct {
	IdTransaksi     uint
	IdTagihan       uint
	Bank            string
	Kurir           string
	NoVA            string
	TipePembayaran  string
	GrandTotal      int64
	KodeTransaksi   string
	StatusTransaksi string

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
	GetUserData(PaymentCore Core) (Client, error)
	GetGrandTotal(PaymentCore Core) (grandTotal int64, err error)
	InsertTransaksi(PaymentData Core) error
	InsertTagihan(PaymentData Core) (idTagihan uint, err error)
	UpdateStokProduct(PaymentData Core) error
	UpdateTransaksi(PaymentData Core) error
}

type PaymentUsecase interface {
	Create(PaymentData Core) (Core, error)
	PaymentWebHook(OrderID, status string) error
}

func ToMidtransCore(PaymentCore Core) *coreapi.ChargeReq {
	return &coreapi.ChargeReq{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  PaymentCore.KodeTransaksi,
			GrossAmt: PaymentCore.GrandTotal,
		},
	}
}
