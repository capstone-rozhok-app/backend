package delivery

import "rozhok/features/produk"

type ProdukRequest struct {
	Nama      string `json:"nama_product" form:"nama_product"`
	Image_url string `json:"image_url" form:"image_url"`
	Stok      int    `json:"stok" form:"stok"`
	Harga     string `json:"harga" form:"harga"`
	Desc      string `json:"desc" form:"desc"`
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
