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
		Email: 					dataCore.Email,
		Password: 				dataCore.Password,
		JunkStationName:  		dataCore.JunkStationName,
		Username: 				dataCore.JunkStationOwner,
		StatusKemitraan: 		dataCore.Status,
		Provinsi: 				dataCore.Provinsi,
		Kota:					dataCore.Kota,
		Kecamatan: 				dataCore.Kecamatan,
		Telepon: 				dataCore.Telp,
		Jalan: 					dataCore.Jalan,

	}
	return dataModel
}

func (dataCore *User)ToCore() js.Core {
	return js.Core{
		JunkStationID: 				int(dataCore.ID),
		Email: 					dataCore.Email,
		Password: 				dataCore.Password,
		JunkStationName:  		dataCore.JunkStationName,
		JunkStationOwner: 		dataCore.Username,
		Status: 				dataCore.StatusKemitraan,
		Provinsi: 				dataCore.Provinsi,
		Kota:					dataCore.Kota,
		Kecamatan: 				dataCore.Kecamatan,
		Telp: 					dataCore.Telepon,
		Jalan: 					dataCore.Jalan,
	}
}

func CoreList(dataCore []User) []js.Core {
	var data []js.Core
	for key := range dataCore {
		data = append(data, dataCore[key].ToCore())
	}
	return data
}