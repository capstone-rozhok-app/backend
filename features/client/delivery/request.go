package delivery

import "rozhok/features/client"

type ClientRequest struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
	Username string `json:"username" form:"username" validate:"required"`
	Telp     string `json:"telepon" form:"telepon" validate:"required"`
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
