package delivery

import "rozhok/features/kategori"

type Request struct {
	Nama         string `json:"nama_kategori" form:"nama_kategori"`
	Harga_mitra  int    `json:"harga_mitra" form:"harga_mitra"`
	Harga_client int    `json:"harga_client" form:"harga_client"`
	Desc         string `json:"desc" form:"desc"`
}

func toCore(dataRequest Request) kategori.Core {
	return kategori.Core{
		Nama:         dataRequest.Nama,
		Harga_mitra:  dataRequest.Harga_mitra,
		Harga_client: dataRequest.Harga_client,
		Desc:         dataRequest.Desc,
	}
}
