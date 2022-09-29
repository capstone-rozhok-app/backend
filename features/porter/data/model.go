package data

import (
	"rozhok/features/porter"

	"gorm.io/gorm"
)

type Porter struct {
	gorm.Model
	Nama      string
	Email     string `gorm:unique`
	Password  string
	Username  string
	Telp      string
	Role      string
	Provinsi  string
	Kota      string
	Kecamatan string
	Jalan     string
}

func fromCore(dataCore porter.Core) Porter {
	return Porter{
		Nama:      dataCore.Nama,
		Email:     dataCore.Email,
		Password:  dataCore.Password,
		Username:  dataCore.Username,
		Telp:      dataCore.Telp,
		Role:      dataCore.Role,
		Provinsi:  dataCore.Provinsi,
		Kota:      dataCore.Kota,
		Kecamatan: dataCore.Kecamatan,
		Jalan:     dataCore.Jalan,
	}
}

// func toCore(data Porter) porter.Core {
// 	return porter.Core{
// 		ID:       int(data.ID),
// 		Email:    data.Email,
// 		Password: data.Password,
// 		Role:     data.Role,
// 		Username: data.Username,
// 	}
// }
