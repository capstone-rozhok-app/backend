package data

import (
	"errors"
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

func (repo *clientData) LoginClient(email string) (client.Core, error) {

	var data Client
	txEmail := repo.db.Where("email = ?", email).First(&data)
	if txEmail.Error != nil {
		return client.Core{}, txEmail.Error
	}

	if txEmail.RowsAffected != 1 {
		return client.Core{}, txEmail.Error
	}

	var dataUser = toCore(data)

	return dataUser, nil

}

func (repo *clientData) UpdateClient(data client.Core, id int) (row int, err error) {
	tx := repo.db.Model(&Client{}).Where("id = ?", id).Updates(fromCore(data))
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("gagal meperbarui data")
	}
	return int(tx.RowsAffected), nil
}

func (repo *clientData) DeleteDataClient(id int) (row int, err error) {
	tx := repo.db.Delete(&Client{}, id)
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("gagal menghapus akun")
	}
	return int(tx.RowsAffected), nil
}
