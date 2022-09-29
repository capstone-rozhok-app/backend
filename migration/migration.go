package migration

import (
	clientModel "rozhok/features/client/data"
	porterModel "rozhok/features/porter/data"

	userModel "rozhok/features/login/data"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&clientModel.Client{})
	db.AutoMigrate(&porterModel.Porter{})
	db.AutoMigrate(&userModel.User{})
}
