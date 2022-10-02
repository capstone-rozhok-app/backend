package data

import (
	"rozhok/features/porter"

	"gorm.io/gorm"
)

type porterData struct {
	db *gorm.DB
}

func New(db *gorm.DB) porter.DataInterface {
	return &porterData{
		db: db,
	}
}

func (repo *porterData) InsertPorter(porter porter.Core) (int, error) {
	porterModel := fromCore(porter)

	tx := repo.db.Model(&User{}).Create(&porterModel)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return int(tx.RowsAffected), nil
}

func (repo *porterData) UpdatePorter(porter porter.Core, id uint) (row int, err error) {
	porterModel := fromCore(porter)

	tx := repo.db.Model(&User{}).Where("id = ?", id).Updates(&porterModel)

	if tx.Error != nil {
		return 0, tx.Error
	}

	return int(tx.RowsAffected), nil
}

func (repo *porterData) DeletePorter(id uint) (row int, err error) {
	porterModel := User{}
	porterModel.ID = id
	tx := repo.db.Model(&User{}).Delete(&porterModel)

	if tx.Error != nil {
		return 0, tx.Error
	}

	return int(tx.RowsAffected), nil
}

func (repo *porterData) GetAll() (rows []porter.Core, err error) {
	porterModels := []User{}

	tx := repo.db.Model(&User{}).Where("role = ?", "porter").Find(&porterModels)
	if tx.Error != nil {
		return []porter.Core{}, tx.Error
	}

	porterCores := []porter.Core{}
	for _, porter := range porterModels {
		porterCores = append(porterCores, toCore(porter))
	}

	return porterCores, nil
}

func (repo *porterData) GetPendapatan(porter porter.Core) (row porter.Core, err error) {
	reportPorter := []struct {
		TipeTransaksi string `json:"tipe_transaksi"`
		GrandTotal    int64  `json:"grand_total"`
	}{}
	var tx = repo.db

	if porter.StartDate != "" && porter.EndDate != "" {
		tx = repo.db.Raw("SELECT tipe_transaksi, SUM(grand_total) grand_total FROM transaksi_porters tp WHERE created_at >= ? AND created_at <= ? AND porter_id  = ? GROUP BY tipe_transaksi", porter.StartDate, porter.EndDate, porter.ID).Scan(&reportPorter)
	} else {
		tx = repo.db.Raw("SELECT tipe_transaksi, SUM(grand_total) grand_total FROM transaksi_porters tp WHERE porter_id  = ? GROUP BY tipe_transaksi", porter.ID).Scan(&reportPorter)
	}

	if tx.Error != nil {
		return row, tx.Error
	}

	for _, report := range reportPorter {
		if report.TipeTransaksi == "pembelian" {
			row.TotalPembelian = report.GrandTotal
		} else {
			row.TotalPenjualan = report.GrandTotal
		}
	}

	row.Laba = row.TotalPenjualan - row.TotalPembelian

	return row, nil
}

func (repo *porterData) Get(id uint) (row porter.Core, err error) {
	porterModel := User{}
	porterModel.ID = id
	tx := repo.db.Model(&User{}).Where("role = ?", "porter").First(&porterModel)
	if tx.Error != nil {
		return porter.Core{}, tx.Error
	}

	return toCore(porterModel), nil
}
