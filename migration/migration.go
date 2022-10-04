package migration

import (
	addresModel "rozhok/features/alamat/data"
	kategoriModel "rozhok/features/kategori/data"
	usermodel "rozhok/features/login/data"
	pengambilanrosokmodel "rozhok/features/pengambilan_rosok/data"
	penjualanclientmodel "rozhok/features/penjualan_client/data"
	produkModel "rozhok/features/produk/data"
	transaksiportermodel "rozhok/features/transaksi_porter/data"


	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(
		new(usermodel.User),
		new(addresModel.Alamat),
		new(produkModel.Produk),
		new(kategoriModel.KategoriRosok),
		new(penjualanclientmodel.KeranjangRosok),
		new(pengambilanrosokmodel.Alamat),
		new(pengambilanrosokmodel.TransaksiClient),
		new(pengambilanrosokmodel.TransaksiClientDetail),
		new(transaksiportermodel.TransaksiPorter),
		new(transaksiportermodel.TransaksiPorterDetail),
	)
}
