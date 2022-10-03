package data

import (
	"gorm.io/gorm"
	pjs "rozhok/features/pembelian_js"
)

type PembelianJS struct{
	gorm.Model
	Kategori		string
	Berat			int
	Harga			int
	User			User
}

type User struct {
	gorm.Model
	Email    			string
	Password 			string
	Role    			string
	Username 			string
	JunkStationName		string
	StatusKemitraan   	string
	Foto 				string
	Provinsi 			string
	Kota 				string
	Kecamatan 			string
	Jalan 				string
	Telepon 			string
	PembelianJS			[]PembelianJS
}

func FromCore(junkCore pjs.PembelianCore) PembelianJS {
	junkModel := PembelianJS{
		Kategori: junkCore.Kategori,
		Berat: junkCore.Berat,
		Harga:	junkCore.Harga,
	}
	return junkModel
}

func (junkCore *PembelianJS) ToCore() pjs.PembelianCore {
	return pjs.PembelianCore{
		ID: 	int(junkCore.ID),	
		Kategori: junkCore.Kategori,
		Berat: junkCore.Berat,
		Harga: junkCore.Harga,
	}
}

func CoreList(junkCore []PembelianJS) []pjs.PembelianCore {
	var junk []pjs.PembelianCore
	for key := range junkCore {
		junk = append(junk, junkCore[key].ToCore())
	}
	return junk
}