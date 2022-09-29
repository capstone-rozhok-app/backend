package delivery

import "rozhok/features/client"

type ClientRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Username string `json:"username" form:"username"`
	Telp     string `json:"no.telp" form:"no.telp"`
	Role     string
	AlamatId int
}

type AuthRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func toCore(dataRequest ClientRequest) client.Core {
	return client.Core{
		Email:    dataRequest.Email,
		Password: dataRequest.Password,
		Username: dataRequest.Username,
		Telp:     dataRequest.Telp,
		Role:     dataRequest.Role,
		AlamatId: dataRequest.AlamatId,
	}
}
