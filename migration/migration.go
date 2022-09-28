package migration

import (
	clientModel "rozhok/features/client/data"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&clientModel.Client{})
}
