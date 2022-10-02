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

func (u *Usecase) GetJunkStationById(id int) (js.Core, error) {
	getID, err := u.jsData.FindJunkStationById(id)
	return getID, err
}

func (u *Usecase) PutJunkStation(id int, data js.Core) (int, error){
	junkMap := make(map[string]interface{})
	if data.JunkStationName != "" {
		junkMap["JunkStationName"] = &data.JunkStationName
	}
	if data.Status != "" {
		junkMap["Status"] = &data.Status
	}
	result, err := u.jsData.UpdateJunkStation(id, data)
	if err != nil{
		return 0, errors.New("JunkStation failed to update")
	}
	return result, err
}