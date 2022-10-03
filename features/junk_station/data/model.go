package data

import (
	"gorm.io/gorm"
	js "rozhok/features/junk_station"
)

type User struct {
	gorm.Model
	Email    			string
	Password 			string
	Role    			string
	Username 			string
	JunkStationName		string
	StatusKemitraan   	string
	Foto 				string
	Provinsi 			string
	Kota 				string
	Kecamatan 			string
	Jalan 				string
	Telepon 			string
}

func FromCore(dataCore js.Core) User {
	dataModel := User{
		JunkStationName:  		dataCore.JunkStationName,
		StatusKemitraan: 		dataCore.Status,
	}
	return dataModel
}

func (dataCore *User)ToCore() js.Core {
	return js.Core{
		JunkStationID: 				int(dataCore.ID),
		JunkStationName: 	dataCore.JunkStationName,
	}
}

func CoreList(dataCore []User) []js.Core {
	var data []js.Core
	for key := range dataCore {
		data = append(data, dataCore[key].ToCore())
	}
	return data
}
