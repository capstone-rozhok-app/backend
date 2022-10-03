package delivery

import (
	penjualanclient "rozhok/features/penjualan_client"
)

type Response struct {
	PenjualanID uint   `json:"id_penjualan"`
	Kategori    string `json:"kategori"`
}

func ToResponse(penjualanRosokCore penjualanclient.Core) Response {
	return Response{
		PenjualanID: penjualanRosokCore.ID,
		Kategori:    penjualanRosokCore.NamaKategori,
	}
}
