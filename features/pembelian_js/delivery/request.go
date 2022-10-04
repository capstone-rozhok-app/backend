package delivery

import (
	pjs "rozhok/features/pembelian_js"
)

type PembelianRequest struct{
	Kategori	int		`json:"kategori" form:"kategori" validate:"required"`
	Berat		int64		`json:"berat" form:"berat" validate:"required"`
	Harga		int64		`json:"harga" form:"harga" validate:"required"`
}

func FromCoreReq(r PembelianRequest) pjs.PembelianCore {
	return pjs.PembelianCore{
		Kategori: r.Kategori,
		Berat: int(r.Berat),
		Harga: int(r.Harga),
	}
}

func ToCore(r PembelianRequest) pjs.PembelianCore {
	return pjs.PembelianCore{
		Kategori: r.Kategori,
		Berat: int(r.Berat),
		Harga: int(r.Harga),
	}
}