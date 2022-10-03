package delivery

import (
	pjs "rozhok/features/pembelian_js"
)

type PembelianRequest struct{
	ID			int		`json:"id" form:"id" validate:"required"`
	Kategori	string	`json:"kategori" form:"kategori" validate:"required"`
	Berat		int		`json:"berat" form:"berat" validate:"required"`
	Harga		int		`json:"harga" form:"harga" validate:"required"`
}

func FromCoreReq(data PembelianRequest) pjs.PembelianCore {
	return pjs.PembelianCore{
		ID: data.ID,
		Kategori: data.Kategori,
		Berat: data.Berat,
		Harga: data.Harga,
	}
}