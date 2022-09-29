package delivery

type LoginResponse struct {
	Token    string `json:"token" form:"token"`
	Role     string `json:"role" form:"role"`
	Username string `json:"username" form:"username"`
}

func FromLoginCore(token, role, username string) LoginResponse {
	return LoginResponse{
		Token:    token,
		Role:     role,
		Username: username,
	}
}
