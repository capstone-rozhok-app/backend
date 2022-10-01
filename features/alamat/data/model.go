package data

import (
	"rozhok/features/client"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
	Role     string
	Username string
	Foto     string
	Telepon  string
	Bonus    int64
}

func fromCore(dataCore client.Core) User {
	return User{
		Email:    dataCore.Email,
		Password: dataCore.Password,
		Role:     dataCore.Role,
		Username: dataCore.Username,
		Foto:     dataCore.Foto,
		Telepon:  dataCore.Telepon,
		Bonus:    dataCore.Bonus,
	}
}
