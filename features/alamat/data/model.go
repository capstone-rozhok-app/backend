package data

import (
	"rozhok/features/alamat"

	"gorm.io/gorm"
)

type Alamat struct {
	gorm.Model
	Provinsi  string
	Kota      string
	Kecamatan string
	Jalan     string
	UserId    uint
	Status    string
	User      User
}

type User struct {
	gorm.Model
	Username string
	Telepon  string
	Alamat   []Alamat
}

func fromCore(dataCore alamat.Core) Alamat {
	return Alamat{
		Provinsi:  dataCore.Provinsi,
		Kota:      dataCore.Kota,
		Kecamatan: dataCore.Kecamatan,
		Jalan:     dataCore.Jalan,
		Status:    dataCore.Status,
		UserId:    dataCore.UserId,
	}
}

func (dataAlamat *Alamat) toCore() alamat.ResponseCore {
	return alamat.ResponseCore{
		ID:        dataAlamat.ID,
		Provinsi:  dataAlamat.Provinsi,
		Kota:      dataAlamat.Kota,
		Kecamatan: dataAlamat.Kecamatan,
		Jalan:     dataAlamat.Jalan,
		Status:    dataAlamat.Status,
		User:      dataAlamat.User.Username,
	}
}

func toCoreList(dataAlamat []Alamat) []alamat.ResponseCore {
	var dataCore []alamat.ResponseCore

	for key := range dataAlamat {
		dataCore = append(dataCore, dataAlamat[key].toCore())

	}

	return dataCore

}
