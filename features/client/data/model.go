package data

import (
	"rozhok/features/client"

	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	Nama     string
	Email    string `gorm:unique`
	Password string
	Username string
	Telp     string
	Role     string
	AlamatId int
	Alamat   Alamat
}

type Alamat struct {
	gorm.Model
	Provinsi  string
	Kota      string
	Kecamatan string
	Jalan     string
	Status    string
	Client    []Client
}

func fromCore(dataCore client.Core) Client {
	return Client{
		Nama:     dataCore.Nama,
		Email:    dataCore.Email,
		Password: dataCore.Password,
		Username: dataCore.Username,
		Telp:     dataCore.Telp,
		Role:     dataCore.Role,
		AlamatId: dataCore.AlamatId,
	}
}
