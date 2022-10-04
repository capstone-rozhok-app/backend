package usecase

import (
	"errors"
	pjs "rozhok/features/pembelian_js"
)

type Usecase struct{
	pembelianJS pjs.DataInterface
}

func New(data pjs.DataInterface) pjs.UsecaseInterface {
	return &Usecase{
		pembelianJS: data,
	}
}

func (u *Usecase) GetPembelian(JunkStationID int)([]pjs.PembelianCore, error) {
	result, err := u.pembelianJS.FindPembelian(JunkStationID)
	return result, err
}

func (u *Usecase) CreatePembelian(data pjs.PembelianCore) (int, error) {
	result, err := u.pembelianJS.InsertPembelian(data)
	if err != nil{
		return 0, errors.New("failed to create pembelian")
	}
	return result, nil
}

func (u *Usecase) PutPembelian(id int, data pjs.PembelianCore)(int, error) {
	result, err := u.pembelianJS.UpdatePembelian(id, data)
	if err != nil{
		return 0, errors.New("failed to update pembelian")
	}
	return result, err
}

func (u *Usecase) DeletePembelian(id int, data pjs.PembelianCore) (int, error) {
	result, err:= u.pembelianJS.DeletePembelian(id, data)
	if err != nil {
		return -1, errors.New("failed to delete pembelian")
	}
	return result, err
}