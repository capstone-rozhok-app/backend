package usecase

import (
	"errors"
	"rozhok/features/produk"
)

type produkUsecase struct {
	produkData produk.DataInterface
}

func New(dataProduk produk.DataInterface) produk.UsecaseInterface {
	return &produkUsecase{
		produkData: dataProduk,
	}
}

func (usecase *produkUsecase) CreateProduk(produk produk.Core) (int, error) {
	if produk.Nama == "" || produk.Harga == "" || produk.Image_url == "" || produk.Desc == "" {
		return -1, errors.New("tidak boleh ada data yang kosong")
	}

	row, err := usecase.produkData.CreateProduk(produk) //produkData.InsertAddress(address)
	return row, err
}

func (usecase *produkUsecase) UpdateProduk(newData produk.Core, id int) (int, error) {

	row, err := usecase.produkData.UpdateProduk(newData, id)
	return row, err
}

func (usecase *produkUsecase) GetAllProduk() (data []produk.Core, err error) {
	results, err := usecase.produkData.GetAllProduk()
	return results, err
}

func (usecase *produkUsecase) GetProduk(id int) (produk.Core, error) {
	result, err := usecase.produkData.GetProduk(id)
	if err != nil {
		return produk.Core{}, err
	}
	return result, nil
}

func (usecase *produkUsecase) DeleteProduk(id int) (int, error) {
	row, err := usecase.produkData.DeleteProduk(id)
	return row, err
}
