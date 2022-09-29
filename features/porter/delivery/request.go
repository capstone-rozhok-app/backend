package delivery

import "rozhok/features/porter"

type PorterRequest struct {
	Nama      string `json:"nama" form:"nama"`
	Email     string `json:"email" form:"email"`
	Password  string `json:"password" form:"password"`
	Username  string `json:"username" form:"username"`
	Telp      string `json:"no.telp" form:"no.telp"`
	Role      string
	Provinsi  string `json:"provinsi" form:"provinsi"`
	Kota      string `json:"kota" form:"kota"`
	Kecamatan string `json:"kecamatan" form:"kecamatan"`
	Jalan     string `json:"jalan" form:"jalan"`
}

type AuthRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func toCore(dataRequest PorterRequest) porter.Core {
	return porter.Core{
		Nama:      dataRequest.Nama,
		Email:     dataRequest.Email,
		Password:  dataRequest.Password,
		Username:  dataRequest.Username,
		Telp:      dataRequest.Telp,
		Role:      dataRequest.Role,
		Provinsi:  dataRequest.Provinsi,
		Kota:      dataRequest.Kota,
		Kecamatan: dataRequest.Kecamatan,
		Jalan:     dataRequest.Jalan,
	}
}
