package data

import (
	"rozhok/features/produk"

	"gorm.io/gorm"
)

type Produk struct {
	gorm.Model
	Nama      string
	Image_url string
	Stok      int
	Harga     int64
	Desc      string
}

func fromCore(dataCore produk.Core) Produk {
	return Produk{
		Nama:      dataCore.Nama,
		Image_url: dataCore.Image_url,
		Stok:      dataCore.Stok,
		Harga:     dataCore.Harga,
		Desc:      dataCore.Descr,
	}
}

func (dataCore *Produk) toCore() produk.Core {
	return produk.Core{
		ID:        dataCore.ID,
		Nama:      dataCore.Nama,
		Image_url: dataCore.Image_url,
		Stok:      dataCore.Stok,
		Harga:     dataCore.Harga,
		Descr:     dataCore.Desc,
	}
}

func toCoreList(data []Produk) []produk.Core {
	var dataCore []produk.Core

	for key := range data {
		dataCore = append(dataCore, data[key].toCore())

	}

	return dataCore

}
