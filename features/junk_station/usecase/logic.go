package usecase

import (
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

func (u *Usecase) CreateJunkStation(junkCreate js.Core) (int, error) {
	if junkCreate.JunkStationName == "" || junkCreate.JunkStationOwner == ""{
		
	}
}