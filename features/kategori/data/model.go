package data

import (
	"rozhok/features/kategori"

	"gorm.io/gorm"
)

type KategoriRosok struct {
	gorm.Model
	NamaKategori string
	HargaMitra   int64
	HargaClient  int64
	Desc         string
}

func fromCore(dataCore kategori.Core) KategoriRosok {
	return KategoriRosok{
		NamaKategori: dataCore.Nama,
		HargaMitra:   int64(dataCore.Harga_mitra),
		HargaClient:  int64(dataCore.Harga_client),
		Desc:         dataCore.Desc,
	}
}

func (dataCore *KategoriRosok) toCore() kategori.Core {
	return kategori.Core{
		ID:           dataCore.ID,
		Nama:         dataCore.NamaKategori,
		Harga_mitra:  int(dataCore.HargaMitra),
		Harga_client: int(dataCore.HargaClient),
		Desc:         dataCore.Desc,
	}
}

func toCoreList(data []KategoriRosok) []kategori.Core {
	var dataCore []kategori.Core

	for key := range data {
		dataCore = append(dataCore, data[key].toCore())

	}

	return dataCore

}
