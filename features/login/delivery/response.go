package delivery

import "rozhok/features/login"

type LoginResponse struct {
	Token        string `json:"token" form:"token"`
	Role         string `json:"role" form:"role"`
	Username     string `json:"username" form:"username"`
	Status_Mitra string `json:"status_kemitraan" form:"status_kemitraan"`
}

type DboardResponse struct {
	TotalJS int                `json:"total_junk_station" form:"total_junk_station"`
	TotalCL int                `json:"total_client" form:"total_client"`
	Grafik  []login.GrafikData `json:"grafik" form:"grafik"`
}

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
