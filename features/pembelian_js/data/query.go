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

func (q *DataPembelian) FindPembelian(JunkStationID int) ([]pjs.PembelianCore, error) {
	var data []KeranjangRosok
	tx := q.db.Where("client_id = ?", JunkStationID).Preload("KategoriRosok").Find(&data)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return CoreList(data), nil
}

func (q *DataPembelian) InsertPembelian(data pjs.PembelianCore) (int, error) {
	DpModel := FromCore(data)
	tx := q.db.Model(&KeranjangRosok{}).Create(&DpModel)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}

func (q *DataPembelian) UpdatePembelian(id int, data pjs.PembelianCore)(int, error) {
	tx := q.db.Model(&KeranjangRosok{}).Where("id = ?", id).Updates(FromCore(data))
	if tx.Error != nil{
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("failed to update Pembelian")
	}
	return int(tx.RowsAffected), nil
}

func (q *DataPembelian) DeletePembelian(id int, data pjs.PembelianCore)(int, error){
	var keranjangModel KeranjangRosok
	keranjangModel.ID = uint(id)
 	tx := q.db.Model(&KeranjangRosok{}).Delete(&keranjangModel)
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0{
		return 0, errors.New("failed to delete Pembelian")
	}
	return int(tx.RowsAffected), nil
}