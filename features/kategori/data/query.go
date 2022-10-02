package data

import (
	"errors"
	"rozhok/features/kategori"

	"gorm.io/gorm"
)

type kategoriData struct {
	db *gorm.DB
}

func New(db *gorm.DB) kategori.DataInterface {
	return &kategoriData{
		db: db,
	}
}

func (repo *kategoriData) CreateKategori(kategori kategori.Core) (int, error) {
	kategoriModel := fromCore(kategori)

	tx := repo.db.Create(&kategoriModel)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return int(tx.RowsAffected), nil
}

func (repo *kategoriData) UpdateKategori(data kategori.Core, id int) (row int, err error) {
	tx := repo.db.Model(&Kategori{}).Where("id = ?", id).Updates(fromCore(data))
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("gagal meperbarui data")
	}
	return int(tx.RowsAffected), nil
}

func (repo *kategoriData) GetAllKategori() ([]kategori.Core, error) {
	var allkategoriData []Kategori
	tx := repo.db.Find(&allkategoriData)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return toCoreList(allkategoriData), nil
}

func (repo *kategoriData) DeleteKategori(id int) (row int, err error) {
	tx := repo.db.Where("id = ?", id).Delete(&Kategori{})
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("gagal menghapus akun")
	}
	return int(tx.RowsAffected), nil
}
