package migration

import (
	clientModel "rozhok/features/client/data"
	// userModel "rozhok/features/login/data"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&clientModel.Client{})
	// db.AutoMigrate(&userModel.User{})
}
