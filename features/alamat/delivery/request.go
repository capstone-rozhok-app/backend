package delivery

import "rozhok/features/alamat"

type AlamatRequest struct {
	Provinsi  string `json:"provinsi" form:"provinsi" validate:"required`
	Kota      string `json:"kota" form:"kota" validate:"required`
	Kecamatan string `json:"kecamatan" form:"kecamatan" validate:"required`
	Jalan     string `json:"jalan" form:"jalan" validate:"required`
	Status    string `json:"status" form:"status" validate:"required`
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
