package delivery

type LoginResponse struct {
	Token    string `json:"token" form:"token"`
	Role     string `json:"role" form:"role"`
	Username string `json:"username" form:"username"`
}
