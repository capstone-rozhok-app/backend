package data

import (
	transaksiclient "rozhok/features/transaksi_client"

	"gorm.io/gorm"
)

type TransaksiClientData struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *TransaksiClientData {
	return &TransaksiClientData{
		DB: db,
	}
}

func (r *TransaksiClientData) GetAll(TransaksiClientCore transaksiclient.Core) ([]transaksiclient.Core, error) {
	transaksiClientModelList := []TransaksiClient{}

	tx := r.DB.Model(&TransaksiClient{})

	if TransaksiClientCore.StartDate != "" {
		tx.Where("created_at >=", TransaksiClientCore.StartDate)
	}

	if TransaksiClientCore.EndDate != "" {
		tx.Where("created_at <=", TransaksiClientCore.EndDate)
	}

	if TransaksiClientCore.TipeTransaksi != "" {
		tx.Where("tipe_transaksi = ?", TransaksiClientCore.TipeTransaksi)
	}

	if TransaksiClientCore.Status != "" {
		tx.Where("status = ?", TransaksiClientCore.Status)
	}

	tx.Find(&transaksiClientModelList)

	if tx.Error != nil {
		return []transaksiclient.Core{}, tx.Error
	}

	transaksiClientCoreList := []transaksiclient.Core{}
	for _, transaksiClient := range transaksiClientModelList {
		transaksiClientCoreList = append(transaksiClientCoreList, ToCore(transaksiClient))
	}

	return transaksiClientCoreList, nil
}

func (r *TransaksiClientData) Get(TransaksiClientCore transaksiclient.Core) (transaksiclient.Core, error) {
	transaksiClientModel := TransaksiClient{}
	transaksiClientModel.ID = transaksiClientModel.ID

	tx := r.DB.Model(&TransaksiClient{}).
}

func (r *TransaksiClientData) Insert(TransaksiClientCore transaksiclient.Core) (int, error) {

}

func (r *TransaksiClientData) Update(TransaksiClientCore transaksiclient.Core) (int, error) {

}
