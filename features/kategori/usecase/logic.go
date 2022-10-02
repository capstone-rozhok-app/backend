package usecase

import (
	"errors"
	"rozhok/features/kategori"
)

type kategoriUsecase struct {
	kategoriData kategori.DataInterface
}

func New(dataKategori kategori.DataInterface) kategori.UsecaseInterface {
	return &kategoriUsecase{
		kategoriData: dataKategori,
	}
}

func (usecase *kategoriUsecase) CreateKategori(kategori kategori.Core) (int, error) {
	if kategori.Nama == "" || kategori.Harga_client == 0 || kategori.Harga_mitra == 0 || kategori.Desc == "" {
		return -1, errors.New("tidak boleh ada data yang kosong")
	}

	row, err := usecase.kategoriData.CreateKategori(kategori)
	return row, err
}

func (usecase *kategoriUsecase) UpdateKategori(newData kategori.Core, id int) (int, error) {

	row, err := usecase.kategoriData.UpdateKategori(newData, id)
	return row, err
}

func (usecase *kategoriUsecase) GetAllKategori() (data []kategori.Core, err error) {
	results, err := usecase.kategoriData.GetAllKategori()
	return results, err
}

func (usecase *kategoriUsecase) DeleteKategori(id int) (int, error) {
	row, err := usecase.kategoriData.DeleteKategori(id)
	return row, err
}
