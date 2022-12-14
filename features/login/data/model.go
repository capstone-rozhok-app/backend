package data

import (
	"rozhok/features/login"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email           string `gorm:"unique"`
	Password        string
	Role            string
	Username        string
	StatusKemitraan string
	JunkStationName string
	ImageUrl        string
	Provinsi        string
	Kota            string
	Kecamatan       string
	Jalan           string
	Telepon         string
	Bonus           int64
}

func toCore(userModel User) login.Core {

	core := login.Core{
		ID:       int(userModel.ID),
		Email:    userModel.Email,
		Password: userModel.Password,
		Username: userModel.Username,
		Role:     userModel.Role,
		Status:   userModel.StatusKemitraan,
	}
	return core
}
