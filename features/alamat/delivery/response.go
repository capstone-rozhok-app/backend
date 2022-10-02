package delivery

import "rozhok/features/alamat"

type AlamatResponse struct {
	Id        uint   `json:"id" form:"id"`
	Provinsi  string `json:"provinsi" form:"provinsi"`
	Kota      string `json:"kota" form:"kota"`
	Kecamatan string `json:"kecamatan" form:"kecamatan"`
	Jalan     string `json:"jalan" form:"jalan"`
	Status    string `json:"status" form:"status"`
	User      string `json:"user" form:"user"`
}

func fromCore(dataCore alamat.ResponseCore) AlamatResponse {
	return AlamatResponse{
		Id:        dataCore.ID,
		Provinsi:  dataCore.Provinsi,
		Kota:      dataCore.Kota,
		Kecamatan: dataCore.Kecamatan,
		Jalan:     dataCore.Jalan,
		Status:    dataCore.Status,
		User:      dataCore.User,
	}
}

func fromCoreList(dataCore []alamat.ResponseCore) []AlamatResponse {
	var dataResponse []AlamatResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
