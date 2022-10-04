package delivery

import transaksijunkstation "rozhok/features/transaksi_junk_station"

type Response struct {
	IdTransaksi    uint   `json:"id_transaksi,omitempty"`
	KodeTF         string `json:"kode_tf,omitempty"`
	HargaTransaksi int64  `json:"harga_transaksi,omitempty"`
	GrandTotal     int64  `json:"grand_total,omitempty"`
	TanggalDibuat  string `json:"tanggal_dibuat,omitempty"`

	BarangRosok []BarangRosok `json:"barang_rosok,omitempty"`
}

type BarangRosok struct {
	ID       uint   `json:"id,omitempty"`
	Kategori string `json:"kategori,omitempty"`
	Berat    int64  `json:"berat,omitempty"`
	Harga    int64  `json:"harga,omitempty"`
}

func FromCore(Core transaksijunkstation.Core) Response {
	response := Response{
		IdTransaksi:    Core.ID,
		KodeTF:         Core.KodeTF,
		HargaTransaksi: Core.GrandTotal,
		GrandTotal:     Core.GrandTotal,
		TanggalDibuat:  Core.CreatedAt,
	}

	barangRosok := []BarangRosok{}
	for _, barangrosok := range Core.BarangRosok {
		barangRosok = append(barangRosok, BarangRosok{
			ID:       barangrosok.ID,
			Kategori: barangrosok.Kategori,
			Berat:    int64(barangrosok.Berat),
			Harga:    barangrosok.Subtotal,
		})
	}

	response.BarangRosok = barangRosok

	return response
}
