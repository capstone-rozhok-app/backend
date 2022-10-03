package data

import (
	penjualanclient "rozhok/features/penjualan_client"

	"gorm.io/gorm"
)

type KeranjangRosok struct {
	gorm.Model
	ClientID        uint
	KategoriRosokID uint
	Berat           int
	Subtotal        int64

	KategoriRosok KategoriRosok
}

type KategoriRosok struct {
	gorm.Model
	NamaKategori string
	HargaMitra   int64
	HargaClient  int64
	Desc         string
}

func ToCore(keranjangRosok KeranjangRosok) penjualanclient.Core {
	return penjualanclient.Core{
		ID:           keranjangRosok.ID,
		KategoriID:   keranjangRosok.KategoriRosokID,
		ClientID:     keranjangRosok.ClientID,
		NamaKategori: keranjangRosok.KategoriRosok.NamaKategori,
	}
}

func FromCore(penjualanClientCore penjualanclient.Core) KeranjangRosok {
	keranjangRosok := KeranjangRosok{
		ClientID:        penjualanClientCore.ClientID,
		KategoriRosokID: penjualanClientCore.KategoriID,
	}

	keranjangRosok.ID = penjualanClientCore.ID
	return keranjangRosok
}
