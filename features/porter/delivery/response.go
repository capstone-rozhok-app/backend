package delivery

import "rozhok/features/porter"

type PorterResponse struct {
	ID		  int	`json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	NoTelp    string `json:"no_telp"`
	Provinsi  string `json:"provinsi"`
	Kota      string `json:"kota"`
	Kecamatan string `json:"kecamatan"`
	Jalan     string `json:"jalan"`
}

type PorterResponseReport struct {
	TotalLaba      int64 `json:"total_laba"`
	TotalPembelian int64 `json:"total_pembelian"`
	TotalPenjualan int64 `json:"total_penjualan"`
}

func fromCore(porterCore porter.Core) PorterResponse {
	return PorterResponse{
		ID: 	   int(porterCore.ID),
		Name:      porterCore.Username,
		Email:     porterCore.Email,
		NoTelp:    porterCore.Telp,
		Provinsi:  porterCore.Provinsi,
		Kota:      porterCore.Kota,
		Kecamatan: porterCore.Kecamatan,
		Jalan:     porterCore.Jalan,
	}
}
