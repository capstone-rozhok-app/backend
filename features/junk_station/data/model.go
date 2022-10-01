package data

import (
	"gorm.io/gorm"
	js "rozhok/features/junk_station"
)
type JunkStation struct{
	gorm.Model
	JunkStationName		string
	Status				string
	User				User
}

type User struct {
	gorm.Model
	Email    			string
	Password 			string
	Role    			string
	Username 			string
	StatusKemitraan   	string
	Foto 				string
	Provinsi 			string
	Kota 				string
	Kecamatan 			string
	Jalan 				string
	Telepon 			string
	JunkStation			[]JunkStation
}

func FromCore(dataCore js.Core) JunkStation {
	dataModel := JunkStation{
		JunkStationName:  dataCore.JunkStationName,
	}
	return dataModel
}

func (dataCore *JunkStation)ToCore() js.Core {
	return js.Core{
		JunkStationID: 		int(dataCore.ID),
		JunkStationName: 	dataCore.JunkStationName,
	}
}

func CoreList(dataCore []JunkStation) []js.Core {
	var data []js.Core
	for key := range dataCore {
		data = append(data, dataCore[key].ToCore())
	}
	return data
}