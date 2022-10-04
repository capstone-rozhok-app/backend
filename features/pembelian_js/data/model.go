package data

import (
	pjs "rozhok/features/pembelian_js"

	"gorm.io/gorm"
)

type KeranjangRosok struct {
	gorm.Model
	ClientID        uint
	KategoriRosokID uint
	Berat           int
	Subtotal        int64
	KategoriRosok   KategoriRosok
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
		ClientID:        uint(junkCore.JunkStationID),
		KategoriRosokID: uint(junkCore.Kategori),
		Berat:           junkCore.Berat,
		Subtotal:        int64(junkCore.Harga),
	}
	return junkModel
}

func ToCore(junkCore KeranjangRosok) pjs.PembelianCore {
	return pjs.PembelianCore{
		IDPembelian:   int(junkCore.ID),
		JunkStationID: int(junkCore.ClientID),
		NamaKategori:  junkCore.KategoriRosok.NamaKategori,
		Berat:         junkCore.Berat,
	}
}

func CoreList(junkCore []KeranjangRosok) []pjs.PembelianCore {
	var junk []pjs.PembelianCore
	for _, v := range junkCore {
		junk = append(junk, ToCore(v))
	}
	return junk
}
