package delivery

import "rozhok/features/payment"

type Response struct {
	NoVa           string `json:"no_va"`
	TipePembayaran string `json:"tipe_pembayaran"`
	TotalHarga     int64  `json:"total_harga"`
	Bank           string `json:"bank"`
	ExpireAt       string `json:"batas_pembayaran"`
}

func FromCore(paymentCore payment.Core) Response {
	return Response{
		NoVa:           paymentCore.NoVA,
		TipePembayaran: paymentCore.TipePembayaran,
		TotalHarga:     paymentCore.GrandTotal,
		Bank:           paymentCore.Bank,
		ExpireAt:       paymentCore.ExpiredAt,
	}
}
