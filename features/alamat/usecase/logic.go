package usecase

import (
	"errors"
	"rozhok/features/alamat"
)

type addressUsecase struct {
	addressData alamat.DataInterface
}

func New(dataAddress alamat.DataInterface) alamat.UsecaseInterface {
	return &addressUsecase{
		addressData: dataAddress,
	}
}

func (usecase *addressUsecase) CreateAddress(address alamat.Core) (int, error) {
	if address.Provinsi == "" || address.Kota == "" || address.Kecamatan == "" || address.Jalan == "" {
		return -1, errors.New("tidak boleh ada data yang kosong")
	}

	row, err := usecase.addressData.InsertAddress(address)
	return row, err
}

func (usecase *addressUsecase) PutAddress(newData alamat.Core, id, userId int) (int, error) {

	row, err := usecase.addressData.UpdateAdress(newData, id, userId)
	return row, err
}

func (usecase *addressUsecase) DeleteAddress(id, userId int) (int, error) {
	row, err := usecase.addressData.DeleteDataAddress(id, userId)
	return row, err
}

func (usecase *addressUsecase) GetAllAddress(userId int) (data []alamat.ResponseCore, err error) {
	results, err := usecase.addressData.GetAllAddress(userId)
	return results, err
}

func (usecase *addressUsecase) GetAddress(id int) (alamat.ResponseCore, error) {
	result, err := usecase.addressData.GetAddress(id)
	if err != nil {
		return alamat.ResponseCore{}, err
	}
	return result, nil
}
