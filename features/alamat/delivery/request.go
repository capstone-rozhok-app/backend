package delivery

import "rozhok/features/alamat"

type AlamatRequest struct {
	Provinsi  string `json:"provinsi" form:"provinsi"`
	Kota      string `json:"kota" form:"kota"`
	Kecamatan string `json:"kecamatan" form:"kecamatan"`
	Jalan     string `json:"jalan" form:"jalan"`
	Status    string `json:"status" form:"status"`
	UserId    uint
}

func toCore(dataRequest AlamatRequest) alamat.Core {
	return alamat.Core{
		Provinsi:  dataRequest.Provinsi,
		Kota:      dataRequest.Kota,
		Kecamatan: dataRequest.Kecamatan,
		Jalan:     dataRequest.Jalan,
		Status:    dataRequest.Status,
		UserId:    dataRequest.UserId,
	}
}
