package migration

import (
	addresModel "rozhok/features/alamat/data"
	userModel "rozhok/features/login/data"
	produkModel "rozhok/features/produk/data"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&userModel.User{})
	db.AutoMigrate(&addresModel.Alamat{})
	db.AutoMigrate(&produkModel.Produk{})
}
