package data

import (
	"errors"
	js "rozhok/features/junk_station"

	"gorm.io/gorm"
)

type DataJS struct {
	db *gorm.DB
}

func New(db *gorm.DB) js.DataInterface {
	return &DataJS{
		db: db,
	}
}

func (junk *DataJS) FindJunkStationAll(dataCore js.Core) (row []js.Core, err error) {
	var data []User
	tx := junk.db.Model(&User{}).Where("role = ?", "junk_station").Where("status_kemitraan = ?", "terverifikasi")
	if dataCore.Provinsi != "" {
		tx.Where("provinsi = ?", dataCore.Provinsi)
	}
	if dataCore.Kota != "" {
		tx.Where("kota = ?", dataCore.Kota)
	}
	if dataCore.Kecamatan != "" {
		tx.Where("kecamatan = ?", dataCore.Kecamatan)
	}
	tx.Find(&data)

	if tx.Error != nil {
		return row, tx.Error
	}

	for _, v := range data {
		row = append(row, ToCore(v))
	}
	return row, nil
}

func (junk *DataJS) FindJunkStationById(id int) (js.Core, error) {
	var data User
	tx := junk.db.Where("role = ?", "junk_station").First(&data, id)
	if tx.Error != nil {
		return js.Core{}, tx.Error
	}
	return ToCore(data), nil
}

func (junk *DataJS) InsertJunkStation(data js.Core) (int, error) {
	dataModel := FromCore(data)
	dataModel.StatusKemitraan = "belum_terverifikasi"
	dataModel.Role = "junk_station"
	tx := junk.db.Create(&dataModel)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}

func (junk *DataJS) UpdateJunkStation(id int, data js.Core) (int, error) {
	tx := junk.db.Model(&User{}).Where("id = ?", id).Updates(FromCore(data))
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("failed update junk station")
	}
	return int(tx.RowsAffected), nil
}
