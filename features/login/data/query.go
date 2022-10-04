package data

import (
	"fmt"
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

func (repo *data) GetUsers() (login.ResponseCore, error) {
	var userCores login.ResponseCore

	var client int64
	repo.db.Model(&User{}).Where("role = ?", "client").Count(&client)
	var mitra int64
	repo.db.Model(&User{}).Where("role = ?", "mitra").Count(&mitra)
	userCores.TotalCL = int(client)
	userCores.TotalJS = int(mitra)

	var userModel []User
	repo.db.Model(&User{}).Where("role = ?", "client").Or("role = ?", "mitra").Find(&userModel)
	// userCores.Grafik =
	fmt.Println(userModel)

	return userCores, nil
}
