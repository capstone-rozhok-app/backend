package delivery

import "rozhok/features/porter"

type PorterResponse struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	NoTelp    string `json:"no_telp"`
	Provinsi  string `json:"provinsi"`
	Kota      string `json:"kota"`
	Kecamatan string `json:"kecamatan"`
	Jalan     string `json:"jalan"`
}

func fromCore(porterCore porter.Core) PorterResponse {
	return PorterResponse{
		Name:      porterCore.Username,
		Email:     porterCore.Email,
		NoTelp:    porterCore.Telp,
		Provinsi:  porterCore.Provinsi,
		Kota:      porterCore.Kota,
		Kecamatan: porterCore.Kecamatan,
		Jalan:     porterCore.Jalan,
	}
}
