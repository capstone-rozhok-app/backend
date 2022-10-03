package data

import (
	"errors"
	pjs "rozhok/features/pembelian_js"

	"gorm.io/gorm"
)
type DataPembelian struct{
	db *gorm.DB
}

func New(db *gorm.DB) pjs.DataInterface {
	return &DataPembelian{
		db: db,
	}
}

func (q *DataPembelian) FindPembelian() ([]pjs.PembelianCore, error) {
	var data []PembelianJS
	tx := q.db.Find(&data)
	if tx.Error != nil {
		return nil, tx.Error
	}
	DP := CoreList(data)
	return DP, nil
}

func (q *DataPembelian) InsertPembelian(data pjs.PembelianCore) (int, error) {
	DpModel := FromCore(data)
	tx := q.db.Model(DpModel)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}

func (q *DataPembelian) UpdatePembelian(id int, data pjs.PembelianCore)(int, error) {
	tx := q.db.Model(&PembelianJS{}).Where("id = ?", data.ID).Updates(FromCore(data))
	if tx.Error != nil{
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("Failed to update Pembelian")
	}
	return int(tx.RowsAffected), nil
}

func (q *DataPembelian) DeletePembelian(id int, data pjs.PembelianCore)(int, error){
 	tx := q.db.Model(&PembelianJS{}).Where("id = ?", data.ID). Updates(FromCore(data))
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0{
		return 0, errors.New("failed to delete Pembelian")
	}
	return int(tx.RowsAffected), nil
}