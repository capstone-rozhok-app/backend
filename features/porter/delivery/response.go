package delivery

type LoginResponse struct {
	Token    string `json:"token" form:"token"`
	Role     string `json:"role" form:"role"`
	Username string `json:"username" form:"username"`
}

func fromLoginCore(token, role, username string) LoginResponse {
	return LoginResponse{
		Token:    token,
		Role:     role,
		Username: username,
	}

}
