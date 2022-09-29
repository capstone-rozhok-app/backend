package data

import (
	"rozhok/features/porter"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email           string
	Password        string
	Role            string
	Username        string
	StatusKemitraan string
	Foto            string
	Provinsi        string
	Kota            string
	Kecamatan       string
	Jalan           string
	Telepon         string
}

func fromCore(dataCore porter.Core) User {
	return User{
		Username:  dataCore.Username,
		Telepon:   dataCore.Telp,
		Email:     dataCore.Email,
		Password:  dataCore.Password,
		Role:      dataCore.Role,
		Provinsi:  dataCore.Provinsi,
		Kota:      dataCore.Kota,
		Kecamatan: dataCore.Kecamatan,
		Jalan:     dataCore.Jalan,
	}
}

func toCore(data User) porter.Core {
	return porter.Core{
		ID:       int(data.ID),
		Email:    data.Email,
		Password: data.Password,
		Role:     data.Role,
		Username: data.Username,
	}
}
