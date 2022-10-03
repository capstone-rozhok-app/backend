package migration

import (
	usermodel "rozhok/features/login/data"
	transaksiportermodel "rozhok/features/transaksi_porter/data"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(
		new(usermodel.User),
		new(transaksiportermodel.KategoriRosok),
		new(transaksiportermodel.TransaksiPorter),
		new(transaksiportermodel.TransaksiPorterDetail),
	)
}
