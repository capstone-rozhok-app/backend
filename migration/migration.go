package migration

import (
	addresModel "rozhok/features/alamat/data"
	kategoriModel "rozhok/features/kategori/data"
	usermodel "rozhok/features/login/data"
	produkModel "rozhok/features/produk/data"
	transaksiportermodel "rozhok/features/transaksi_porter/data"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(
		new(usermodel.User),
		new(transaksiportermodel.TransaksiPorter),
		new(transaksiportermodel.TransaksiPorterDetail),
		new(addresModel.Alamat),
		new(produkModel.Produk),
		new(kategoriModel.Kategori),
	)
}
