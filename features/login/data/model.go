package data

import (
	modelClient "rozhok/features/client/data"
	"rozhok/features/login"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string
	Password string
	Role     string
	Username string
	Status   string
	ClientId uint
	Client   modelClient.Client `gorm:"foreignKey:ClientId"`
}

func toCore(user User) login.Core {

	core := login.Core{
		ID:       int(user.ID),
		Email:    user.Email,
		Password: user.Password,
		ClientId: int(user.ClientId),
	}
	return core
}
