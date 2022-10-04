package data

import (
	mitra "rozhok/features/kemitraan"

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

func FromCore(userCore mitra.MitraCore) User {
	userModel:= User{
		StatusKemitraan: userCore.StatusKemitraan,
	}
	return userModel
}

func ToCore(userCore User) mitra.MitraCore {
	return mitra.MitraCore{
		StatusKemitraan: userCore.StatusKemitraan,
	}
}

func CoreList(userCore []User) []mitra.MitraCore {
	var kemitraan []mitra.MitraCore
	for _, v := range userCore {
		kemitraan = append(kemitraan, ToCore(v))
	}
	return kemitraan
}