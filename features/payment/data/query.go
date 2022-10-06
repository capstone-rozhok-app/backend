package data

import (
	"errors"
	"rozhok/features/payment"

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

func (r *PaymentRepo) GetUserData(PaymentCore payment.Core) (payment.Client, error) {
	var client User
	tx := r.DB.Model(&User{}).Preload("Alamat", func(db *gorm.DB) *gorm.DB {
		return db.Where("alamats.status", "utama")
	}).Where("id = ?", PaymentCore.Client.ID).First(&client)
	if tx.Error != nil {
		return payment.Client{}, tx.Error
	}

	if len(client.Alamat) < 1 {
		return payment.Client{}, errors.New("user must have alamat with status utama")
	}

	return payment.Client{
		ID:        client.ID,
		Username:  client.Username,
		Provinsi:  client.Alamat[0].Provinsi,
		Kota:      client.Alamat[0].Kota,
		Kecamatan: client.Alamat[0].Kecamatan,
	}, nil
}

func (r *PaymentRepo) GetGrandTotal(PaymentCore payment.Core) (grandTotal int64, err error) {
	// ambil produk dari keranjang user dengan checklist true
	var keranjangBelanja []Cart
	tx := r.DB.Model(&Cart{}).Where("user_id = ?", PaymentCore.Client.ID).Where("checklist = ?", 1).Find(&keranjangBelanja)
	if tx.Error != nil {
		return 0, tx.Error
	}

	// hitung grand total
	for _, keranjangbelanja := range keranjangBelanja {
		grandTotal += keranjangbelanja.Subtotal
	}

	return
}

func (r *PaymentRepo) InsertTransaksi(PaymentData payment.Core) error {
	// ambil produk dari keranjang user dengan checklist true
	var keranjangBelanja []Cart
	tx := r.DB.Model(&Cart{}).Where("user_id = ?", PaymentData.Client.ID).Where("checklist = ?", 1).Find(&keranjangBelanja)
	if tx.Error != nil {
		return tx.Error
	}

	// masukkan kedalam transaksi master
	transaksi := TransaksiClient{
		ClientID:      PaymentData.Client.ID,
		TagihanID:     PaymentData.IdTagihan,
		Kurir:         PaymentData.Kurir,
		TipeTransaksi: "pembelian",
		GrandTotal:    PaymentData.GrandTotal,
		Status:        "belum_bayar",
		KodeTransaksi: PaymentData.KodeTransaksi,
	}
	tx = r.DB.Model(&TransaksiClient{}).Omit("porter_id").Create(&transaksi)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected < 1 {
		return errors.New("error row affected")
	}

	// mapping dari keranjang user ke dalam transaksi detail client
	var transaksiDetail []TransaksiClientDetail
	for _, keranjangbelanja := range keranjangBelanja {
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
		return tx.Error
	}
	if tx.RowsAffected < 1 {
		return errors.New("error row affected")
	}

	// hapus data dari keranjang user dengan checklist true
	var cart Cart
	tx = r.DB.Model(&Cart{}).Where("user_id = ?", PaymentData.Client.ID).Where("checklist = ?", 1).Unscoped().Delete(&cart)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected < 1 {
		return errors.New("error row affected")
	}

	return nil
}

func (r *PaymentRepo) InsertTagihan(PaymentData payment.Core) (idTagihan uint, err error) {
	// buat tagihan
	tagihan := Tagihan{
		NoVa:           PaymentData.NoVA,
		TipePembayaran: PaymentData.TipePembayaran,
		Bank:           PaymentData.Bank,
		GrandTotal:     PaymentData.GrandTotal,
	}

	tx := r.DB.Model(&Tagihan{}).Create(&tagihan)
	if tx.Error != nil {
		return 0, tx.Error
	}
	if tx.RowsAffected < 1 {
		return 0, errors.New("error row affected")
	}

	return tagihan.ID, nil
}
