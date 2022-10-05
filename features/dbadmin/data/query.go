package data

import (
	"rozhok/features/dbadmin"

	"gorm.io/gorm"
)

type data struct {
	db *gorm.DB
}

func New(db *gorm.DB) dbadmin.DataInterface {
	return &data{
		db: db,
	}
}

func (repo *data) GetUsers() (dbadmin.ResponseCore, error) {
	var userCores dbadmin.ResponseCore

	var client int64
	repo.db.Model(&User{}).Where("role = ?", "client").Count(&client)
	var mitra int64
	repo.db.Model(&User{}).Where("role = ?", "junk_station").Count(&mitra)
	userCores.TotalCL = int(client)
	userCores.TotalJS = int(mitra)

	var sliceGrafik []dbadmin.GrafikData
	repo.db.Model(&User{}).Select("month(created_at) as bulan, count(id) as jumlah_cl").Where("role = ?", "client").Group("month(created_at)").Find(&sliceGrafik)
	var sliceGrafik2 []dbadmin.GrafikData
	repo.db.Model(&User{}).Select("month(created_at) as bulan, count(id) as jumlah_js").Where("role = ?", "junk_station").Group("month(created_at)").Find(&sliceGrafik2)

	for i := range sliceGrafik {
		for j := range sliceGrafik2 {
			if sliceGrafik[i].Bulan == sliceGrafik2[j].Bulan {
				sliceGrafik[i].JumlahJS = sliceGrafik2[j].JumlahJS
			}
		}
	}

	userCores.Grafik = sliceGrafik

	return userCores, nil
}
