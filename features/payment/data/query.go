package data

import (
	"errors"
	"rozhok/features/payment"
	"rozhok/utils/helper"

	"gorm.io/gorm"
)

type PaymentRepo struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *PaymentRepo {
	return &PaymentRepo{
		DB: db,
	}
}

func (r *PaymentRepo) GetUserData(PaymentCore payment.Core) (payment.Core, error) {
	var client User
	tx := r.DB.Model(&User{}).Preload("Alamat", func(db *gorm.DB) {
		db.Where("status = ?", "utama")
	}).Where("id = ?", PaymentCore.Client.ID).First(&client)
	if tx.Error != nil {
		return payment.Core{}, tx.Error
	}

	if len(client.Alamat) < 1 {
		return payment.Core{}, errors.New("user must have alamat with status utama")
	}

	return payment.Core{
		Client: payment.Client{
			ID:        client.ID,
			Username:  client.Username,
			Provinsi:  client.Alamat[0].Provinsi,
			Kota:      client.Alamat[0].Kota,
			Kecamatan: client.Alamat[0].Kecamatan,
		},
	}, nil
}

func (r *PaymentRepo) Insert(PaymentData payment.Core) (int, error) {
	// ambil produk dari keranjang user dengan checklist true
	var keranjangBelanja []Cart
	tx := r.DB.Model(&Cart{}).Where("user_id = ?", PaymentData.Client.ID).Where("checklist = ?", 1).Find(&keranjangBelanja)
	if tx.Error != nil {
		return 0, tx.Error
	}

	// hitung grand total
	var grandTotal int64
	for _, keranjangbelanja := range keranjangBelanja {
		grandTotal += keranjangbelanja.Subtotal
	}

	// buat tagihan
	tagihan := Tagihan{
		NoVa:           "1241274012",
		TipePembayaran: "bank_transaksi",
		Bank:           "BCA",
		GrandTotal:     grandTotal,
	}
	tx = r.DB.Model(&Tagihan{}).Create(&tagihan)
	if tx.Error != nil {
		return 0, tx.Error
	}
	if tx.RowsAffected < 1 {
		return 0, errors.New("error row affected")
	}
	
	// masukkan kedalam transaksi master
	transaksi := TransaksiClient{
		ClientID:      PaymentData.Client.ID,
		TagihanID:     tagihan.ID,
		Kurir:         PaymentData.Kurir,
		TipeTransaksi: "pembelian",
		GrandTotal:    grandTotal,
		Status:        "belum_bayar",
		KodeTransaksi: helper.GenerateTF(int(PaymentData.Client.ID)),
	}
	tx = r.DB.Model(&TransaksiClient{}).Omit("porter_id").Create(&transaksi)
	if tx.Error != nil {
		return 0, tx.Error
	}
	if tx.RowsAffected < 1 {
		return 0, errors.New("error row affected")
	}
	// mapping dari keranjang user ke dalam transaksi detail client
	var transaksiDetail []TransaksiClientDetail
	for _, keranjangbelanja := range keranjangBelanja {
		grandTotal += keranjangbelanja.Subtotal
		transaksiDetail = append(transaksiDetail, TransaksiClientDetail{
			TransaksiClientID: transaksi.ID,
			ProdukID:          keranjangbelanja.ProdukId,
			Qty:               keranjangbelanja.Qty,
			Subtotal:          keranjangbelanja.Subtotal,
		})
	}

	// masukkan kedalam detail transaksi
	tx = r.DB.Model(&TransaksiClientDetail{}).Omit("kategori_id").Create(&transaksiDetail)
	if tx.Error != nil {
		return 0, tx.Error
	}
	if tx.RowsAffected < 1 {
		return 0, errors.New("error row affected")
	}

	// hapus data dari keranjang user dengan checklist true
	var cart Cart
	tx = r.DB.Model(&Cart{}).Where("user_id = ?", PaymentData.Client.ID).Where("checklist = ?", 1).Delete(&cart)
	if tx.Error != nil {
		return 0, tx.Error
	}
	if tx.RowsAffected < 1 {
		return 0, errors.New("error row affected")
	}

	return 1, nil
}
