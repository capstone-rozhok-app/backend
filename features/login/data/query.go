package data

import (
	"rozhok/features/login"

	"gorm.io/gorm"
)

type data struct {
	db *gorm.DB
}

func New(db *gorm.DB) login.DataInterface {
	return &data{
		db: db,
	}
}

func (repo *data) LoginUser(email string) (login.Core, error) {

	var data User
	txEmail := repo.db.Model(&User{}).Where("email = ?", email).First(&data)
	if txEmail.Error != nil {
		return login.Core{}, txEmail.Error
	}

	if txEmail.RowsAffected != 1 {
		return login.Core{}, txEmail.Error
	}

	var dataUser = toCore(data)

	return dataUser, nil

}
