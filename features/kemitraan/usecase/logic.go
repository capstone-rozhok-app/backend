package usecase

import (
	"errors"
	mitra "rozhok/features/kemitraan"
)

type Usecase struct{
	statusMitra mitra.DataInterface
}

func New(data mitra.DataInterface) mitra.UsecaseInterface {
	return &Usecase{
		statusMitra: data,
	}
}

func (u *Usecase) PutKemitraan(id int, data mitra.MitraCore)(int, error) {
	result, err := u.statusMitra.UpdateKemitraan(id, data)
	if err != nil {
		return -1, errors.New("failed to update kemitraan")
	}
	return result, err
}