package data

import (
	"errors"
	"rozhok/features/produk"
	tsModel "rozhok/features/transaksi_client/data"

	"gorm.io/gorm"
)

type produkData struct {
	db *gorm.DB
}

func New(db *gorm.DB) produk.DataInterface {
	return &produkData{
		db: db,
	}
}

func (repo *produkData) CreateProduk(produk produk.Core) (int, error) {
	produkModel := fromCore(produk)

	tx := repo.db.Create(&produkModel)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return int(tx.RowsAffected), nil
}

func (repo *produkData) UpdateProduk(data produk.Core, id int) (row int, err error) {
	tx := repo.db.Model(&Produk{}).Where("id = ?", id).Updates(fromCore(data))
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("gagal meperbarui data")
	}
	return int(tx.RowsAffected), nil
}

func (repo *produkData) GetAllProduk() ([]produk.Core, error) {
	var allprodukData []Produk
	tx := repo.db.Find(&allprodukData)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return toCoreList(allprodukData), nil
}

func (repo *produkData) GetProduk(id int) (produk.Core, error) {
	var produkData Produk

	tx := repo.db.First(&produkData, id)

	if tx.Error != nil {
		return produk.Core{}, tx.Error
	}
	return produkData.toCore(), nil
}

func (repo *produkData) DeleteProduk(id int) (row int, err error) {
	tx := repo.db.Where("id = ?", id).Delete(&Produk{})
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("gagal menghapus akun")
	}
	return int(tx.RowsAffected), nil
}

func (repo *produkData) GetFavorite() (data []produk.Core, err error) {
	var slice [8]produk.Core
	repo.db.Model(&tsModel.TransaksiClientDetail{}).Joins("Produk").Select("produk_id as id, sum(qty) as total", "nama as nama, image_url as image_url, stok as stok, harga as harga, descr as descr").Group("id").Order("total desc").Find(&slice)

	return slice[:], nil
}
