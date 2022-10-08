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
	tx := junk.db.Model(&User{}).Where("role = ?", "junk_station")
	if dataCore.Provinsi != "" {
		tx.Where("provinsi = ?", dataCore.Provinsi)
	}
	if dataCore.Kota != "" {
		tx.Where("kota = ?", dataCore.Kota)
	}
	if dataCore.Kecamatan != "" {
		tx.Where("kecamatan = ?", dataCore.Kecamatan)
	}
	if dataCore.StatusKemitraan != "" {
		tx.Where("status_kemitraan = ?", dataCore.StatusKemitraan)
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

func (q *DataJS) UpdateKemitraan(id int) (int, error) {
	tx := q.db.Model(&User{}).Where("id = ?", id).Update("status_kemitraan", "terverifikasi")
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("failed update kemitraan")
	}
	return int(tx.RowsAffected), nil
}

func (junk *DataJS) Dashboard(data js.Core) (int64, error) {
	var grandTotal int64
	var tx = junk.db

	if data.StartDate != "" && data.EndDate != "" {
		tx = tx.Raw("SELECT SUM(grand_total) grand_total FROM transaksi_junk_stations WHERE created_at >= ? AND created_at <= ? AND user_id = ", data.StartDate, data.EndDate, data.JunkStationID).Scan(&grandTotal)
	} else {
		tx = tx.Raw("SELECT SUM(grand_total) grand_total FROM transaksi_junk_stations WHERE user_id = ?", data.JunkStationID).Scan(&grandTotal)
	}

	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrInvalidValue) {
			return 0, tx.Error
		}
	}

	return grandTotal, nil
}
