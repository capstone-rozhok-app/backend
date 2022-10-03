package data

import (
	"errors"
	penjualanclient "rozhok/features/penjualan_client"

	"gorm.io/gorm"
)

type PenjualanClient struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *PenjualanClient {
	return &PenjualanClient{
		DB: db,
	}
}

func (r *PenjualanClient) GetAll(PenjualanClientCore penjualanclient.Core) ([]penjualanclient.Core, error) {
	KeranjangRosokList := []KeranjangRosok{}

	if err := r.DB.Model(&KeranjangRosok{}).Where("client_id = ?", PenjualanClientCore.ClientID).Preload("KeranjangRosok").Find(&KeranjangRosokList).Error; err != nil {
		return []penjualanclient.Core{}, err
	}

	penjualanClient := []penjualanclient.Core{}
	for _, rosok := range KeranjangRosokList {
		penjualanClient = append(penjualanClient, ToCore(rosok))
	}

	return penjualanClient, nil
}

func (r *PenjualanClient) Update(PenjualanClientCore penjualanclient.Core) (int, error) {
	keranjangRosok := FromCore(PenjualanClientCore)

	tx := r.DB.Model(&KeranjangRosok{}).Where("id = ?", PenjualanClientCore.ID).Updates(&keranjangRosok)

	if tx.Error != nil {
		return 0, tx.Error
	}

	if tx.RowsAffected < 1 {
		return 0, errors.New("failed to update data")
	}

	return 1, nil
}

func (r *PenjualanClient) Insert(PenjualanClientCore penjualanclient.Core) (int, error) {
	keranjangRosok := FromCore(PenjualanClientCore)

	tx := r.DB.Model(&KeranjangRosok{}).Create(&keranjangRosok)

	if tx.Error != nil {
		return 0, tx.Error
	}

	if tx.RowsAffected < 1 {
		return 0, errors.New("failed to create data")
	}

	return 1, nil
}

func (r *PenjualanClient) Delete(PenjualanClientCore penjualanclient.Core) (int, error) {
	keranjangRosok := FromCore(PenjualanClientCore)

	tx := r.DB.Model(&KeranjangRosok{}).Delete(&keranjangRosok)

	if tx.Error != nil {
		return 0, tx.Error
	}

	if tx.RowsAffected < 1 {
		return 0, errors.New("failed to delete data")
	}

	return 1, nil
}
