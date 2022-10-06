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
	var count int64

	if address.Status == "utama" {
		txUtama := repo.db.Model(&Alamat{}).Where("user_id = ?", address.UserId).Where("status = ?", "utama").Count(&count)
		if txUtama.Error != nil {
			if !errors.Is(txUtama.Error, gorm.ErrRecordNotFound) {
				return 0, txUtama.Error
			}
		}
		if count > 0 {
			return 0, errors.New("status utama has used")
		}
	}

	tx := repo.db.Create(&addressModel)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return int(tx.RowsAffected), nil
}

func (repo *addressData) UpdateAdress(data alamat.Core, id, userId int) (row int, err error) {
	var addressModel Alamat

	if data.Status == "utama" {
		txUtama := repo.db.Model(&Alamat{}).Where("user_id = ?", userId).Where("status = ?", "utama").First(&addressModel)
		if txUtama.Error != nil {
			return 0, txUtama.Error
		}
		if addressModel.ID > 0 {
			txToCadangan := repo.db.Model(&Alamat{}).Where("id = ?", addressModel.ID).Update("status", "cadangan")
			if txToCadangan.Error != nil {
				return 0, txToCadangan.Error
			}
		}
	}

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
	tx := repo.db.Where("user_id = ?", userId).Preload("User").Find(&allAddressData)

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
