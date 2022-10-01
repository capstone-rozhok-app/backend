package delivery

import "rozhok/features/porter"

type PorterRequest struct {
	Email     string `json:"email" form:"email" validate:"required"`
	Password  string `json:"password" form:"password" validate:"required"`
	Name  string `json:"name" form:"name" validate:"required"`
	NoTelp    string `json:"no_telp" form:"no_telp" validate:"required"`
	Provinsi  string `json:"provinsi" form:"provinsi" validate:"required"`
	Kota      string `json:"kota" form:"kota" validate:"required"`
	Kecamatan string `json:"kecamatan" form:"kecamatan" validate:"required"`
	Jalan     string `json:"jalan" form:"jalan" validate:"required"`
}

func toCore(dataRequest PorterRequest) porter.Core {
	return porter.Core{
		Email:     dataRequest.Email,
		Password:  dataRequest.Password,
		Username:  dataRequest.Name,
		Telp:      dataRequest.NoTelp,
		Provinsi:  dataRequest.Provinsi,
		Kota:      dataRequest.Kota,
		Kecamatan: dataRequest.Kecamatan,
		Jalan:     dataRequest.Jalan,
	}
}
