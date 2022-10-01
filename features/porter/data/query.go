package data

import (
	"rozhok/features/porter"

	"gorm.io/gorm"
)

type porterData struct {
	db *gorm.DB
}

func New(db *gorm.DB) porter.DataInterface {
	return &porterData{
		db: db,
	}
}

func (repo *porterData) InsertPorter(porter porter.Core) (int, error) {
	porterModel := fromCore(porter)

	tx := repo.db.Model(&User{}).Create(&porterModel)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return int(tx.RowsAffected), nil
}

func (repo *porterData) UpdatePorter(porter porter.Core, id uint) (row int, err error) {
	porterModel := fromCore(porter)

	tx := repo.db.Model(&User{}).Where("id = ?", id).Updates(&porterModel)

	if tx.Error != nil {
		return 0, tx.Error
	}

	return int(tx.RowsAffected), nil
}

func (repo *porterData) DeletePorter(id uint) (row int, err error) {
	porterModel := User{}
	porterModel.ID = id
	tx := repo.db.Model(&User{}).Delete(&porterModel)

	if tx.Error != nil {
		return 0, tx.Error
	}

	return int(tx.RowsAffected), nil
}

func (repo *porterData) GetAll() (rows []porter.Core, err error) {
	porterModels := []User{}

	tx := repo.db.Model(&User{}).Where("role = ?", "porter").Find(&porterModels)
	if tx.Error != nil {
		return []porter.Core{}, tx.Error
	}

	porterCores := []porter.Core{}
	for _, porter := range porterModels {
		porterCores = append(porterCores, toCore(porter))
	}

	return porterCores, nil
}

func (repo *porterData) Get(id uint) (row porter.Core, err error) {
	porterModel := User{}

	tx := repo.db.Model(&User{}).Where("role = ?", "porter").First(&porterModel)
	if tx.Error != nil {
		return porter.Core{}, tx.Error
	}

	return toCore(porterModel), nil
}
