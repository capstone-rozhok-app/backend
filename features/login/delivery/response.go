package delivery

import "rozhok/features/login"

type LoginResponse struct {
	Token        string `json:"token" form:"token"`
	Role         string `json:"role" form:"role"`
	Username     string `json:"username" form:"username"`
	Status_Mitra string `json:"status_kemitraan" form:"status_kemitraan"`
}

type DboardResponse struct {
	TotalJS int          `json:"total_junk_station" form:"total_junk_station"`
	TotalCL int          `json:"total_client" form:"total_client"`
	Grafik  []login.Core `json:"grafik" form:"grafik"`
}

// type GrafikResponse struct {
// 	bulan         int `json:"bulan" form:"bulan"`
// 	jumlah_client int `json:"jumlah_client" form:"jumlah_client"`
// 	jumlah_js     int `json:"jumlah_js" form:"jumlah_js"`
// }

func FromCore(token, role, username, status_Mitra string) LoginResponse {
	return LoginResponse{
		Token:        token,
		Role:         role,
		Username:     username,
		Status_Mitra: status_Mitra,
	}

}

func fromCore(dataCore login.ResponseCore) DboardResponse {
	return DboardResponse{
		TotalJS: dataCore.TotalJS,
		TotalCL: dataCore.TotalCL,
		Grafik:  dataCore.Grafik,
	}
}

// func fromCoreList(dataCore []login.ResponseCore) []GrafikResponse {
// 	var dataResponse []GrafikResponse
// 	for _, v := range dataCore {
// 		dataResponse = append(dataResponse, fromCore(v))
// 	}
// 	return dataResponse
// }
