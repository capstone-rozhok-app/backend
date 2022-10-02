package delivery

import "rozhok/features/produk"

type Response struct {
	Id        uint   `json:"id" form:"id"`
	Nama      string `json:"nama_product" form:"nama_product"`
	Image_url string `json:"image_url" form:"image_url"`
	Stok      int    `json:"stok" form:"stok"`
	Harga     string `json:"harga" form:"harga"`
	Desc      string `json:"desc" form:"desc"`
}

func fromCore(dataCore produk.Core) Response {
	return Response{
		Id:        dataCore.ID,
		Nama:      dataCore.Nama,
		Image_url: dataCore.Image_url,
		Stok:      dataCore.Stok,
		Harga:     dataCore.Harga,
		Desc:      dataCore.Desc,
	}
}

func fromCoreList(dataCore []produk.Core) []Response {
	var dataResponse []Response
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
