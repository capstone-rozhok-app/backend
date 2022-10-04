package delivery

import "rozhok/features/kategori"

type Response struct {
	Id           uint   `json:"id" form:"id"`
	Nama         string `json:"nama" form:"nama"`
	Harga_mitra  int    `json:"harga_mitra" form:"harga_mitra"`
	Harga_client int    `json:"harga_client" form:"harga_client"`
	Desc         string `json:"desc" form:"desc"`
}

func fromCore(dataCore kategori.Core) Response {
	return Response{
		Id:           dataCore.ID,
		Nama:         dataCore.Nama,
		Harga_mitra:  dataCore.Harga_mitra,
		Harga_client: dataCore.Harga_client,
		Desc:         dataCore.Desc,
	}
}

func fromCoreList(dataCore []kategori.Core) []Response {
	var dataResponse []Response
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
