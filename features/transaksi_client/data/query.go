package data

import (
	"errors"
	transaksiclient "rozhok/features/transaksi_client"
	"rozhok/utils/helper"

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

	tx := r.DB.Model(&TransaksiClient{}).Where("client_id = ?", TransaksiClientCore.Client.ID)

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
	transaksiClientModel.ID = TransaksiClientCore.ID

	tx := r.DB.Model(&TransaksiClient{}).Preload("Client.Alamat", func(db *gorm.DB) *gorm.DB {
		return db.Where("alamats.status", "utama")
	})

	if TransaksiClientCore.TipeTransaksi == "penjualan" {
		tx.Preload("Porter").Preload("DetailTransaksiClient.KategoriRosok")
	} else {
		tx.Preload("DetailTransaksiClient.Produk").Preload("Tagihan")
	}

	tx.First(&transaksiClientModel)

	if tx.Error != nil {
		return transaksiclient.Core{}, tx.Error
	}

	return ToCore(transaksiClientModel), nil
}

// insert penjualan client dari keranjang_rosoks ke transaksi_clients
func (r *TransaksiClientData) Insert(TransaksiClientCore transaksiclient.Core) (int, error) {
	//cek user punya alamat atau tidak
	var alamatUtama int64
	tx := r.DB.Model(&User{}).Joins("JOIN alamats ON alamats.user_id = users.id AND alamats.status = ?", "utama").Count(&alamatUtama)
	if tx.Error != nil {
		return 0, tx.Error
	}

	if alamatUtama < 1 {
		return 0, errors.New("user must have alamat with status utama")
	}

	// ambil barang rosok dari keranjang rosok dari id client
	KeranjangRosokList := []KeranjangRosok{}
	tx = r.DB.Model(&KeranjangRosok{}).Where("client_id = ?", TransaksiClientCore.Client.ID).Find(&KeranjangRosokList)
	if tx.Error != nil {
		return 0, tx.Error
	}

	// buat transaksi client dengan status belum_bayar dan insert
	transaksiClientModel := TransaksiClient{
		ClientID:      TransaksiClientCore.Client.ID,
		Status:        "belum_bayar",
		TipeTransaksi: "penjualan",
		KodeTransaksi: helper.GenerateTF(int(TransaksiClientCore.Client.ID)),
	}
	tx = r.DB.Model(&TransaksiClient{}).Omit("porter_id", "tagihan_id").Create(&transaksiClientModel)
	if tx.Error != nil {
		return 0, tx.Error
	}

	if tx.RowsAffected < 1 {
		return 0, errors.New("affected rows error")
	}

	// buat detail transaksi client dan insert
	detailTransaksiModel := []TransaksiClientDetail{}
	for _, barangrosok := range KeranjangRosokList {
		detailTransaksiModel = append(detailTransaksiModel, TransaksiClientDetail{
			TransaksiClientID: transaksiClientModel.ID,
			KategoriID:        barangrosok.KategoriRosokID,
		})
	}

	tx = r.DB.Model(&TransaksiClientDetail{}).Omit("produk_id").Create(&detailTransaksiModel)
	if tx.Error != nil {
		return 0, tx.Error
	}

	if tx.RowsAffected < 1 {
		return 0, errors.New("affected rows error")
	}

	// hapus keranjang_rosok dengan id client
	keranjangRosok := &KeranjangRosok{}
	tx = r.DB.Model(&KeranjangRosok{}).Where("client_id = ?", TransaksiClientCore.Client.ID).Unscoped().Delete(&keranjangRosok)
	if tx.Error != nil {
		return 0, tx.Error
	}

	if tx.RowsAffected < 1 {
		return 0, errors.New("affected rows error")
	}

	return 1, nil
}

// update ketika pembelian dengan status dikirim ke diterima
func (r *TransaksiClientData) Update(TransaksiClientCore transaksiclient.Core) (int, error) {
	tx := r.DB.Model(&TransaksiClient{}).Where("id = ?", TransaksiClientCore.ID).Update("status", "diterima")
	if tx.Error != nil {
		return 0, tx.Error
	}

	if tx.RowsAffected < 1 {
		return 0, errors.New("affected rows error")
	}

	return 0, nil
}
