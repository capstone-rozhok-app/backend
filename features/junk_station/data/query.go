package data

import (
	"errors"
	js "rozhok/features/junk_station"

	"gorm.io/gorm"
)
type DataJS struct{
	db *gorm.DB
}

func New(db *gorm.DB) js.DataInterface {
	return &DataJS{
		db: db,
	}
}

func (junk *DataJS) FindJunkStationAll() ([]js.Core, error) {
	var data []User
	tx := junk.db.Find(&data)
	if tx.Error != nil {
		return nil, tx.Error
	}
	JS := CoreList(data)
	return JS, nil
}

func (junk *DataJS) FindJunkStationById(id int) (js.Core, error){
	var data User
	tx := junk.db.First(&data, id)
	if tx.Error != nil{
		return js.Core{}, tx.Error
	}
	dataCore := FromCore(js.Core{})
	return dataCore.ToCore(), nil
}

func (junk *DataJS) InsertJunkStation(data js.Core)(int , error) {
	dataModel := FromCore(data)
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