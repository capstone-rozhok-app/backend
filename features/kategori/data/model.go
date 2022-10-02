package data

import (
	"rozhok/features/kategori"

	"gorm.io/gorm"
)

type Kategori struct {
	gorm.Model
	Nama         string
	Harga_mitra  int
	Harga_client int
	Desc         string
}

func fromCore(dataCore kategori.Core) Kategori {
	return Kategori{
		Nama:         dataCore.Nama,
		Harga_mitra:  dataCore.Harga_mitra,
		Harga_client: dataCore.Harga_client,
		Desc:         dataCore.Desc,
	}
}

func (dataCore *Kategori) toCore() kategori.Core {
	return kategori.Core{
		ID:           dataCore.ID,
		Nama:         dataCore.Nama,
		Harga_mitra:  dataCore.Harga_mitra,
		Harga_client: dataCore.Harga_client,
		Desc:         dataCore.Desc,
	}
}

func toCoreList(data []Kategori) []kategori.Core {
	var dataCore []kategori.Core

	for key := range data {
		dataCore = append(dataCore, data[key].toCore())

	}

	return dataCore

}
