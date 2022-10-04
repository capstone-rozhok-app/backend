package delivery

import pjs "rozhok/features/pembelian_js"

type PembelianResponse struct{
	ID			int			`json:"id_pembelian"`
	Kategori	string		`json:"kategori"`
}

func FromCore(data pjs.PembelianCore) PembelianResponse {
	return PembelianResponse{
		ID: 	data.JunkStationID,
		Kategori: data.NamaKategori,
	}
}

func ToResponse(data pjs.PembelianCore) PembelianResponse {
	return PembelianResponse{
		ID:		data.IDPembelian,
		Kategori: data.NamaKategori,
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
		Kategori: data.NamaKategori,
	}
	return dataResponse
}