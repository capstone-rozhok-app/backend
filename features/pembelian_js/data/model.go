package data

import (
	"gorm.io/gorm"
	pjs "rozhok/features/pembelian_js"
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
	KeranjangRosok	[]KeranjangRosok 
}

type KeranjangRosok struct {
	gorm.Model
	JunkStationID   uint
	KategoriRosokID uint
	Berat           int
	Subtotal        int64
	User			User `gorm:"foreignKey:JunkStationID"`
	KategoriRosok KategoriRosok
}

type KategoriRosok struct {
	gorm.Model
	NamaKategori string
	HargaMitra   int64
	HargaClient  int64
	Desc         string
}

func FromCore(junkCore pjs.PembelianCore) KeranjangRosok {
	junkModel := KeranjangRosok{
		Berat: junkCore.Berat,
		Subtotal: int64(junkCore.Harga),
	}
	return junkModel
}

func (junkCore *KeranjangRosok) ToCore() pjs.PembelianCore {
	return pjs.PembelianCore{
		ID: 	int(junkCore.ID),	
		Berat: junkCore.Berat,
		Harga: int(junkCore.Subtotal),
	}
}

func CoreList(junkCore []KeranjangRosok) []pjs.PembelianCore {
	var junk []pjs.PembelianCore
	for key := range junkCore {
		junk = append(junk, junkCore[key].ToCore())
	}
	return junk
}