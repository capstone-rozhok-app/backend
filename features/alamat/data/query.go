package data

import (
	"errors"
	"rozhok/features/alamat"

	"gorm.io/gorm"
)

type addressData struct {
	db *gorm.DB
}

func New(db *gorm.DB) alamat.DataInterface {
	return &addressData{
		db: db,
	}
}

func (repo *addressData) InsertAddress(address alamat.Core) (int, error) {
	addressModel := fromCore(address)

	tx := repo.db.Create(&addressModel)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return int(tx.RowsAffected), nil
}

func (repo *addressData) UpdateAdress(data alamat.Core, id, userId int) (row int, err error) {
	tx := repo.db.Model(&Alamat{}).Where("id = ?", id).Where("user_id = ?", userId).Updates(fromCore(data))
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("gagal meperbarui data")
	}
	return int(tx.RowsAffected), nil
}

func (repo *addressData) DeleteDataAddress(id, userId int) (row int, err error) {
	tx := repo.db.Where("id = ?", id).Where("user_id = ?", userId).Delete(&Alamat{})
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("gagal menghapus akun")
	}
	return int(tx.RowsAffected), nil
}

func (repo *addressData) GetAllAddress(userId int) ([]alamat.ResponseCore, error) {
	var allAddressData []Alamat
	tx := repo.db.Where("user_id = ?", userId).Preload("User").Find(&allAddressData) //.Where("userId = ?", userId)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return toCoreList(allAddressData), nil
}

func (repo *addressData) GetAddress(id int) (alamat.ResponseCore, error) {
	var addressData Alamat

	tx := repo.db.Preload("User").First(&addressData, id)

	if tx.Error != nil {
		return alamat.ResponseCore{}, tx.Error
	}
	return addressData.toCore(), nil
}
