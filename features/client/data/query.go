package data

import (
	"rozhok/features/client"

	"gorm.io/gorm"
)

type clientData struct {
	db *gorm.DB
}

func New(db *gorm.DB) client.DataInterface {
	return &clientData{
		db: db,
	}
}

func (repo *clientData) InsertClient(client client.Core) (int, error) {
	clientModel := fromCore(client)

	tx := repo.db.Create(&clientModel)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return int(tx.RowsAffected), nil
}
