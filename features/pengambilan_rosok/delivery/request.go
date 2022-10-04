package delivery

import transaksiporter "rozhok/features/transaksi_porter"

type Request struct {
	IdBarangRosok uint  `json:"id_barang_rosok" form:"id_barang_rosok" validate:"dive"`
	Berat         uint  `json:"berat" form:"berat" validate:"dive"`
	Subtotal      int64 `json:"subtotal" form:"subtotal" validate:"dive"`
}

func ToCore(transaksiRequest Request) transaksiporter.DetailTransaksiPorter {
	return transaksiporter.DetailTransaksiPorter{
		Id:       transaksiRequest.IdBarangRosok,
		Berat:    int64(transaksiRequest.Berat),
		Subtotal: transaksiRequest.Subtotal,
	}
}
