package usecase

import (
	"errors"
	js "rozhok/features/junk_station"
)

type Usecase struct{
	jsData js.DataInterface
}


func NewLogic(data js.DataInterface) js.UsecaseInterface {
	return &Usecase{
		jsData: data,
	}
}

func (u *Usecase) CreateJunkStation(junkCreate js.Core, id int) (int, error) {
	if junkCreate.JunkStationName == "" || junkCreate.Status == ""{
		return -1, errors.New("your data not fuel field")
	}

	result, err := u.jsData.InsertJunkStation(junkCreate)
	if err != nil{
		return 0, errors.New("failed insert data")
	}

	return result, nil
}

func (u *Usecase) GetJunkStationAll() ([]js.Core, error) {
	result, err := u.jsData.FindJunkStationAll()
	return result, err
}

func (u *Usecase) GetJunkStationById(id, token int) (data js.Core, err error) {
	if data.JunkStationID == 0{
		return data, errors.New("id tidak ditemukan")
	}
	result, err:= u.jsData.FindJunkStationById(id)
	return result, err
}

func (u *Usecase) PutJunkStation(id int, data js.Core) (int, error){
	junkMap := make(map[string]interface{})
	if data.JunkStationName != "" {
		junkMap["js_name"] = &data.JunkStationName
	}
	if data.JunkStationOwner != "" {
		junkMap["js_owner"]  = &data.JunkStationOwner
	}
	if data.Status != "" {
		junkMap["status"] = &data.Status
	}
	if data.Provinsi != "" {
		junkMap["provinsi"] = &data.Provinsi
	}
	if data.Kota != "" {
		junkMap["kota"] = &data.Kota
	}
	if data.Kecamatan != "" {
		junkMap["kecamatan"] = &data.Kecamatan
	}
	if data.Telp != "" {
		junkMap["telp"] = &data.Telp
	}
	if data.Jalan != "" {
		junkMap["jalan"] = &data.Jalan
	}
	result, err := u.jsData.UpdateJunkStation(id, data)
	if err != nil{
		return 0, errors.New("JunkStation failed to update")
	}
	return result, err
}