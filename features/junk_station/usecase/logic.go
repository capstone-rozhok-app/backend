package usecase

import (
	"errors"
	js "rozhok/features/junk_station"

	"golang.org/x/crypto/bcrypt"
)

type Usecase struct {
	jsData js.DataInterface
}

func NewLogic(data js.DataInterface) js.UsecaseInterface {
	return &Usecase{
		jsData: data,
	}
}

func (u *Usecase) CreateJunkStation(junkCreate js.Core) (int, error) {
	passBcyrpt := []byte(junkCreate.Password)
	hash, errHash := bcrypt.GenerateFromPassword(passBcyrpt, bcrypt.DefaultCost)
	if errHash != nil {
		return -2, errors.New("failed to hashing password")
	}

	junkCreate.Password = string(hash)
	result, err := u.jsData.InsertJunkStation(junkCreate)
	if err != nil {
		return 0, err
	}

	return result, nil
}

func (u *Usecase) GetJunkStationAll(dataCore js.Core) ([]js.Core, error) {
	return u.jsData.FindJunkStationAll(dataCore)
}

func (u *Usecase) GetJunkStationById(id int) (data js.Core, err error) {
	result, err := u.jsData.FindJunkStationById(id)
	return result, err
}

func (u *Usecase) PutJunkStation(id int, data js.Core) (int, error) {
	result, err := u.jsData.UpdateJunkStation(id, data)
	if err != nil {
		return 0, errors.New("JunkStation failed to update")
	}
	return result, err
}

func (u *Usecase) PutKemitraan(id int, data js.Core)(int, error) {
	result, err := u.jsData.UpdateKemitraan(id, data)
	if err != nil {
		return -1, errors.New("failed to update kemitraan")
	}
	return result, err
}