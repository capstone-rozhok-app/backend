package data

import (
	"gorm.io/gorm"
	js "rozhok/features/junk_station"
)

type User struct {
	gorm.Model
	Email           string `gorm:"unique"`
	Password        string
	Role            string
	Username        string
	JunkStationName string
	StatusKemitraan string
	ImageUrl        string
	Provinsi        string
	Kota            string
	Kecamatan       string
	Jalan           string
	Telepon         string
}

func FromCore(dataCore js.Core) User {
	dataModel := User{
		Email:           dataCore.Email,
		Password:        dataCore.Password,
		JunkStationName: dataCore.JunkStationName,
		Username:        dataCore.JunkStationOwner,
		StatusKemitraan: dataCore.StatusKemitraan,
		Provinsi:        dataCore.Provinsi,
		Kota:            dataCore.Kota,
		Kecamatan:       dataCore.Kecamatan,
		Telepon:         dataCore.Telp,
		Jalan:           dataCore.Jalan,
	}
	return dataModel
}

func ToCore(junkStation User) js.Core {
	return js.Core{
		JunkStationID:    int(junkStation.ID),
		Email:            junkStation.Email,
		Password:         junkStation.Password,
		JunkStationName:  junkStation.JunkStationName,
		JunkStationOwner: junkStation.Username,
		StatusKemitraan:  junkStation.StatusKemitraan,
		Provinsi:         junkStation.Provinsi,
		Kota:             junkStation.Kota,
		Kecamatan:        junkStation.Kecamatan,
		Telp:             junkStation.Telepon,
		Image_url:        junkStation.ImageUrl,
		Jalan:            junkStation.Jalan,
	}
}

func CoreList(dataCore []User) []js.Core {
	var data []js.Core
	for _, v := range dataCore {
		data = append(data, ToCore(v))
	}
	return data
}
