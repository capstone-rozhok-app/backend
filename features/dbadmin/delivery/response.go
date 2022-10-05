package delivery

import "rozhok/features/dbadmin"

type DboardResponse struct {
	TotalJS int `json:"total_junk_station" form:"total_junk_station"`
	TotalCL int `json:"total_client" form:"total_client"`
	Grafik  []dbadmin.GrafikData
}

func fromCore(dataCore dbadmin.ResponseCore) DboardResponse {
	return DboardResponse{
		TotalJS: dataCore.TotalJS,
		TotalCL: dataCore.TotalCL,
		Grafik:  dataCore.Grafik,
	}
}
