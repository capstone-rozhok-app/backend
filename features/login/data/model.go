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

func toCore(userModel User) login.Core {

	core := login.Core{
		ID:       int(userModel.ID),
		Email:    userModel.Email,
		Password: userModel.Password,
		Username: userModel.Username,
		Role: userModel.Role,
		Status: userModel.Status,
		ClientId: int(userModel.ClientId),
	}
	return core
}
