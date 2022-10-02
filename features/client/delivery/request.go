package delivery

import "rozhok/features/client"

type ClientRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Username string `json:"username" form:"username"`
	Telp     string `json:"telepon" form:"telepon"`
	Role     string
}

func toCore(dataRequest ClientRequest) client.Core {
	return client.Core{
		Email:    dataRequest.Email,
		Password: dataRequest.Password,
		Username: dataRequest.Username,
		Telepon:  dataRequest.Telp,
		Role:     dataRequest.Role,
	}
}
