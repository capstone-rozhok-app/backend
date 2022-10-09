package data

import (
	"errors"
	"rozhok/features/client"
	transaksiModel "rozhok/features/pengambilan_rosok/data"

	"gorm.io/gorm"
)

type clientData struct {
	db *gorm.DB
}

func New(db *gorm.DB) client.DataInterface {
	return &clientData{
		db: db,
	}
}

func (repo *clientData) InsertClient(client client.Core) (int, error) {
	clientModel := fromCore(client)

	tx := repo.db.Create(&clientModel)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return int(tx.RowsAffected), nil
}

func (repo *clientData) UpdateClient(data client.Core, id int) (row int, err error) {
	tx := repo.db.Model(&User{}).Where("id = ?", id).Updates(fromCore(data))
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("gagal meperbarui data")
	}
	return int(tx.RowsAffected), nil
}

func (repo *clientData) DeleteDataClient(id int) (row int, err error) {
	tx := repo.db.Unscoped().Delete(&User{}, id)
	//  repo.db.Delete(&User{}, id)
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("gagal menghapus akun")
	}
	return int(tx.RowsAffected), nil
}

func (repo *clientData) GetClient(id int) (data client.Core, err error) {
	var jual int64
	var bonus int64
	var username string
	tx := repo.db.Model(&transaksiModel.TransaksiClient{}).Select("sum(grand_total)").Where("client_id = ?", id).Where("tipe_transaksi = ?", "penjualan").Where("status = ?", "dibayar").Find(&jual)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrInvalidData) {
			return client.Core{}, nil
		}
	}
	tx = repo.db.Model(&User{}).Select("bonus").Where("id = ?", id).Find(&bonus)
	if tx.Error != nil {
		return client.Core{}, nil
	}
	tx = repo.db.Model(&User{}).Select("username").Where("id = ?", id).First(&username)
	if tx.Error != nil {
		return client.Core{}, nil
	}
	data.TotalJual = jual
	data.Bonus = bonus
	data.Username = username

	return data, nil
}
