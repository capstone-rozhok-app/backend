package delivery

type LoginResponse struct {
	Token        string `json:"token" form:"token"`
	Role         string `json:"role" form:"role"`
	Username     string `json:"username" form:"username"`
	Status_Mitra string `json:"status_kemitraan" form:"status_kemitraan"`
}

func FromCore(token, role, username, status_Mitra string) LoginResponse {
	return LoginResponse{
		Token:        token,
		Role:         role,
		Username:     username,
		Status_Mitra: status_Mitra,
	}

}
