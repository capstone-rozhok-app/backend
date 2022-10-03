package delivery

import transaksiporter "rozhok/features/transaksi_porter"

type Response struct {
	IdTransaksi   uint   `json:"id_transaksi"`
	TipeTransaksi string `json:"tipe_transaksi"`
	Status        string `json:"status"`
	GrandTotal    int64  `json:"grand_total"`

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

func toResponse(TransaksiPorterCore transaksiporter.Core) Response {
	TransaksiPorterResponse := Response{
		IdTransaksi:   TransaksiPorterCore.ID,
		TipeTransaksi: TransaksiPorterCore.TipeTransaksi,
		Status:        TransaksiPorterCore.Status,
		GrandTotal:    TransaksiPorterCore.GrandTotal,
		Client: Client{
			Name:      TransaksiPorterCore.Client.Username,
			NoTelp:    TransaksiPorterCore.Client.Telepon,
			Provinsi:  TransaksiPorterCore.Client.Provinsi,
			Kota:      TransaksiPorterCore.Client.Kota,
			Kecamatan: TransaksiPorterCore.Client.Kecamatan,
		},
	}

	barangRosokList := []BarangRosok{}
	for _, barangRosokCore := range TransaksiPorterCore.DetailTransaksiPorter {
		barangRosokList = append(barangRosokList, BarangRosok{
			Id:            barangRosokCore.Id,
			Kategori:      barangRosokCore.Nama,
			Subtotal:      barangRosokCore.Subtotal,
			Berat:         barangRosokCore.Berat,
			HargaKategori: barangRosokCore.HargaKategori,
		})
	}

	TransaksiPorterResponse.BarangRosok = barangRosokList

	return TransaksiPorterResponse
}
