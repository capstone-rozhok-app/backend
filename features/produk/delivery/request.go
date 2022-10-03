package delivery

import "rozhok/features/produk"

type ProdukRequest struct {
	Nama      string `json:"nama_product" form:"nama_product" validate:"required"`
	Image_url string `json:"image_url" form:"image_url" validate:"required"`
	Stok      int    `json:"stok" form:"stok" validate:"required"`
	Harga     string `json:"harga" form:"harga" validate:"required"`
	Desc      string `json:"desc" form:"desc" validate:"required"`
}

func toCore(dataRequest ProdukRequest) produk.Core {
	return produk.Core{
		Nama:      dataRequest.Nama,
		Image_url: dataRequest.Image_url,
		Stok:      dataRequest.Stok,
		Harga:     dataRequest.Harga,
		Desc:      dataRequest.Desc,
	}
}
