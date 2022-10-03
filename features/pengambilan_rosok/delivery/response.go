package delivery

import (
	pengambilanrosok "rozhok/features/pengambilan_rosok"
)

type Response struct {
	IdPenjualan uint          `json:"id_penjualan"`
	Client      Client        `json:"client"`
	BarangRosok []BarangRosok `json:"barang_rosok"`
}

type Client struct {
	Name      string `json:"name"`
	NoTelp    string `json:"no_telp"`
	Provinsi  string `json:"provinsi"`
	Kota      string `json:"kota"`
	Kecamatan string `json:"kecamatan"`
}

type BarangRosok struct {
	Id            uint   `json:"id"`
	Kategori      string `json:"kategori"`
	Berat         int64  `json:"berat"`
	Subtotal      int64  `json:"subtotal"`
	HargaKategori int64  `json:"harga_kategori"`
}

func toResponse(PengambilanRosokCore pengambilanrosok.Core) Response {
	TransaksiPorterResponse := Response{
		IdPenjualan: PengambilanRosokCore.ID,
		Client: Client{
			Name:      PengambilanRosokCore.Client.Username,
			NoTelp:    PengambilanRosokCore.Client.Telepon,
			Provinsi:  PengambilanRosokCore.Client.Provinsi,
			Kota:      PengambilanRosokCore.Client.Kota,
			Kecamatan: PengambilanRosokCore.Client.Kecamatan,
		},
	}

	barangRosokList := []BarangRosok{}
	for _, barangRosokCore := range PengambilanRosokCore.DetailTransaksiClient {
		barangRosokList = append(barangRosokList, BarangRosok{
			Kategori: barangRosokCore.Nama,
		})
	}

	TransaksiPorterResponse.BarangRosok = barangRosokList
	return TransaksiPorterResponse
}
