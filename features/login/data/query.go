package data

import (
	"fmt"
	"rozhok/features/login"

	"gorm.io/gorm"
)

type data struct {
	db *gorm.DB
}

func New(db *gorm.DB) login.DataInterface {
	return &data{
		db: db,
	}
}

func (repo *data) LoginUser(email string) (login.Core, error) {

	var data User
	txEmail := repo.db.Model(&User{}).Where("email = ?", email).First(&data)
	if txEmail.Error != nil {
		return login.Core{}, txEmail.Error
	}

	if txEmail.RowsAffected != 1 {
		return login.Core{}, txEmail.Error
	}

	var dataUser = toCore(data)

	return dataUser, nil

}

func (repo *data) GetUsers() (login.ResponseCore, error) {
	var userCores login.ResponseCore

	var client int64
	repo.db.Model(&User{}).Where("role = ?", "client").Count(&client)
	var mitra int64
	repo.db.Model(&User{}).Where("role = ?", "mitra").Count(&mitra)
	userCores.TotalCL = int(client)
	userCores.TotalJS = int(mitra)

	type GrafikData struct {
		Bulan    int
		JumlahCL int
		JumlahJS int
	}

	var sliceGrafik []GrafikData
	repo.db.Model(&User{}).Select("month(created_at) as bulan, count(id) as jumlah_cl").Where("role = ?", "client").Group("month(created_at)").Find(&sliceGrafik)
	var sliceGrafik2 []GrafikData
	repo.db.Model(&User{}).Select("month(created_at) as bulan, count(id) as jumlah_js").Where("role = ?", "mitra").Group("month(created_at)").Find(&sliceGrafik2)

	for i := range sliceGrafik {
		for j := range sliceGrafik2 {
			if sliceGrafik[i].Bulan == sliceGrafik2[j].Bulan {
				sliceGrafik[i].JumlahJS = sliceGrafik2[j].JumlahJS
			}
		}
	}

	fmt.Println(sliceGrafik)

	return userCores, nil
}
