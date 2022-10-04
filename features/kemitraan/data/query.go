package data

import (
	"errors"
	mitra "rozhok/features/kemitraan"

	"gorm.io/gorm"
)
type Kemitraan struct{
	db *gorm.DB
}

func New(db *gorm.DB) mitra.DataInterface {
	return &Kemitraan{
		db: db,
	}
}

func (q *Kemitraan) UpdateKemitraan(id int, data mitra.MitraCore) (int, error) {
	tx := q.db.Model(&User{}).Where("id = ?", id).Updates(FromCore(data))
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("failed update kemitraan")
	}
	return int(tx.RowsAffected), nil
}