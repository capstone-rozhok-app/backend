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

func FromCoreReq(r PembelianRequest) pjs.PembelianCore {
	return pjs.PembelianCore{
		ID: r.ID,
		Kategori: r.Kategori,
		Berat: r.Berat,
		Harga: r.Harga,
	}
}

func ToCore(r PembelianRequest) pjs.PembelianCore {
	return pjs.PembelianCore{
		Kategori: r.Kategori,
		Berat: r.Berat,
		Harga: r.Harga,
	}
}