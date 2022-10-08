package migration

import (
	addresModel "rozhok/features/alamat/data"
	cartModel "rozhok/features/cart/data"
	kategoriModel "rozhok/features/kategori/data"
	usermodel "rozhok/features/login/data"
	paymentModel "rozhok/features/payment/data"
	pengambilanrosokmodel "rozhok/features/pengambilan_rosok/data"
	penjualanclientmodel "rozhok/features/penjualan_client/data"
	produkModel "rozhok/features/produk/data"
	transaksiclientmodel "rozhok/features/transaksi_client/data"
	transaksijunkstationModel "rozhok/features/transaksi_junk_station/data"
	transaksiportermodel "rozhok/features/transaksi_porter/data"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(
		new(usermodel.User),
		new(paymentModel.Tagihan),
		new(addresModel.Alamat),
		new(produkModel.Produk),
		new(kategoriModel.KategoriRosok),
		new(penjualanclientmodel.KeranjangRosok),
		new(pengambilanrosokmodel.Alamat),
		new(transaksiclientmodel.TransaksiClient),
		new(transaksiclientmodel.TransaksiClientDetail),
		new(transaksiportermodel.TransaksiPorter),
		new(transaksiportermodel.TransaksiPorterDetail),
		new(transaksiportermodel.LogBonus),
		new(cartModel.Cart),
		new(transaksijunkstationModel.TransaksiJunkStation),
		new(transaksijunkstationModel.TransaksiJunkStationDetail),
	)
}
