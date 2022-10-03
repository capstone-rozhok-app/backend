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

func (u *Usecase) GetPembelian()([]pjs.PembelianCore, error) {
	result, err := u.pembelianJS.FindPembelian()
	return result, err
}

func (u *Usecase) CreatePembelian(data pjs.PembelianCore, token int) (int, error) {
	if data.Kategori == "" || data.Harga == 0 || data.Berat == 0 {
		return -1, errors.New("data must be filled")
	}
	result, err := u.pembelianJS.InsertPembelian(data)
	if err != nil{
		return 0, errors.New("failed to create pembelian")
	}
	return result, nil
}

func (u *Usecase) PutPembelian(id int, data pjs.PembelianCore)(int, error) {
	pjsMap := make(map[string]interface{})
	if data.Kategori != ""{
		pjsMap["kategori"] = &data.Kategori
	}
	if data.Berat != 0{
		pjsMap["berat"] = &data.Berat
	}
	if data.Harga != 0{
		pjsMap["harga"] = &data.Harga
	}
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