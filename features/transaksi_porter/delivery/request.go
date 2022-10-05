package delivery

import transaksiporter "rozhok/features/transaksi_porter"

type Request struct {
	IdBarangRosok uint  `json:"id_barang_rosok" form:"id_barang_rosok"`
	Berat         uint  `json:"berat" form:"berat"`
	Subtotal      int64 `json:"subtotal" form:"subtotal"`
}

type ArrayOfReq struct {
	BarangRosok []Request `json:"barang_rosok"`
}

func toCore(transaksiRequest Request) transaksiporter.DetailTransaksiPorter {
	return transaksiporter.DetailTransaksiPorter{
		Id:       transaksiRequest.IdBarangRosok,
		Berat:    int64(transaksiRequest.Berat),
		Subtotal: transaksiRequest.Subtotal,
	}
}
