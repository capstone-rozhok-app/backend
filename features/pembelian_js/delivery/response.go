package delivery

import pjs "rozhok/features/pembelian_js"

type PembelianResponse struct{
	ID			int			`json:"id"`
	Kategori	string		`json:"kategori"`
	Berat		int			`json:"berat"`
	Harga		int			`json:"harga"`
}

func FromCore(data pjs.PembelianCore) PembelianResponse {
	return PembelianResponse{
		ID: 	data.ID,
		Kategori: data.Kategori,
		Berat: data.Berat,
		Harga: data.Harga,
	}
}

func CoreList(data []pjs.PembelianCore) []PembelianResponse {
	var response []PembelianResponse
	for _, v := range data {
		response = append(response, FromCore(v))
	}
	return response
}

func FromCoreToResponse(data pjs.PembelianCore) PembelianResponse {
	dataResponse := PembelianResponse{
		Kategori: data.Kategori,
		Berat: data.Berat,
		Harga: data.Harga,
	}
	return dataResponse
}