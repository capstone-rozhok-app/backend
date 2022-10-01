package data

import (
	"rozhok/features/porter"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email           string `gorm:"unique"`
	Password        string
	Role            string
	Username        string `gorm:"unique"`
	StatusKemitraan string
	Foto            string
	Provinsi        string
	Kota            string
	Kecamatan       string
	Jalan           string
	Telepon         string `gorm:"unique"`
	Bonus           int64
}

func fromCore(dataCore porter.Core) User {
	return User{
		Username:  dataCore.Username,
		Telepon:   dataCore.Telp,
		Email:     dataCore.Email,
		Password:  dataCore.Password,
		Role:      "porter",
		Provinsi:  dataCore.Provinsi,
		Kota:      dataCore.Kota,
		Kecamatan: dataCore.Kecamatan,
		Jalan:     dataCore.Jalan,
	}
}

func toCore(data User) porter.Core {
	return porter.Core{
		ID:        data.ID,
		Email:     data.Email,
		Password:  data.Password,
		Username:  data.Username,
		Telp:      data.Telepon,
		Provinsi:  data.Provinsi,
		Kota:      data.Kota,
		Kecamatan: data.Kecamatan,
		Jalan:     data.Jalan,
	}
}
