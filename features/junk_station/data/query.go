package data

import (
	"gorm.io/gorm"
	js "rozhok/features/junk_station"
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
	var data []JunkStation
	tx := junk.db.Find(&data)
	if tx.Error != nil {
		return nil, tx.Error
	}
	JS := CoreList(data)
	return JS, nil
}

func (junk *DataJS) FindJunkStationById	(id int) (js.Core, error){
	var data JunkStation
	tx := junk.db.First(&data, id)
	if tx.Error != nil{
		return js.Core{}, tx.Error
	}
	dataCore := FromCore(js.Core{})
	return dataCore.ToCore(), nil
}