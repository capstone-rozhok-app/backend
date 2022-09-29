package data

import (
	"rozhok/features/login"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string
	Password string
	Role     string
	Username string
	StatusKemitraan   string
	Foto string
	Provinsi string
	Kota string
	Kecamatan string
	Jalan string
	Telepon string
}

func toCore(userModel User) login.Core {

	core := login.Core{
		ID:       int(userModel.ID),
		Email:    userModel.Email,
		Password: userModel.Password,
		Username: userModel.Username,
		Role: userModel.Role,
		Status: userModel.StatusKemitraan,
	}
	return core
}
