package data

import (
	"rozhok/features/client"

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

type Alamat struct {
	gorm.Model
	UserID uint
	Provinsi  string
	Kota      string
	Kecamatan string
	Jalan     string
	Status    string
	User User
}

func fromCore(dataCore client.Core) User {
	return User{
		Telepon:     dataCore.Telp,
	}
}

func toCore(data User) client.Core {
	return client.Core{
		ID:       int(data.ID),
	}
}
