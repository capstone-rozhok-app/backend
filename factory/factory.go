package factory

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	porterData "rozhok/features/porter/data"
	porterDelivery "rozhok/features/porter/delivery"
	porterUsecase "rozhok/features/porter/usecase"

	clientData "rozhok/features/client/data"
	clientDelivery "rozhok/features/client/delivery"
	clientUsecase "rozhok/features/client/usecase"

	loginData "rozhok/features/login/data"
	loginDelivery "rozhok/features/login/delivery"
	loginUsecase "rozhok/features/login/usecase"

	transaksiPorterData "rozhok/features/transaksi_porter/data"
	transaksiPorterDelivery "rozhok/features/transaksi_porter/delivery"
	transaksiPorterUsecase "rozhok/features/transaksi_porter/usecase"

	alamatData "rozhok/features/alamat/data"
	alamatDelivery "rozhok/features/alamat/delivery"
	alamatUsecase "rozhok/features/alamat/usecase"

	produkData "rozhok/features/produk/data"
	produkDelivery "rozhok/features/produk/delivery"
	produkUsecase "rozhok/features/produk/usecase"

	kategoriData "rozhok/features/kategori/data"
	kategoriDelivery "rozhok/features/kategori/delivery"
	kategoriUsecase "rozhok/features/kategori/usecase"

	junk_stationData "rozhok/features/junk_station/data"
	junk_stationDelivery "rozhok/features/junk_station/delivery"
	junk_stationUsecase "rozhok/features/junk_station/usecase"
	
	penjualanClientData "rozhok/features/penjualan_client/data"
	penjualanClientDelivery "rozhok/features/penjualan_client/delivery"
	penjualanClientUsecase "rozhok/features/penjualan_client/usecase"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	clientDataFactory := clientData.New(db)
	clientUsecaseFactory := clientUsecase.New(clientDataFactory)
	clientDelivery.New(e, clientUsecaseFactory)

	porterDataFactory := porterData.New(db)
	porterUsecaseFactory := porterUsecase.New(porterDataFactory)
	porterDelivery.New(e, porterUsecaseFactory)

	loginDataFactory := loginData.New(db)
	loginUsecaseFactory := loginUsecase.New(loginDataFactory)
	loginDelivery.New(e, loginUsecaseFactory)

	transaksiPorterDataFactory := transaksiPorterData.New(db)
	transaksiPorterUsecaseFactory := transaksiPorterUsecase.New(transaksiPorterDataFactory)
	transaksiPorterDelivery.New(e, transaksiPorterUsecaseFactory)

	alamatDataFactory := alamatData.New(db)
	alamatUsecaseFactory := alamatUsecase.New(alamatDataFactory)
	alamatDelivery.New(e, alamatUsecaseFactory)

	produkDataFactory := produkData.New(db)
	produkUsecaseFactory := produkUsecase.New(produkDataFactory)
	produkDelivery.New(e, produkUsecaseFactory)

	kategoriDataFactory := kategoriData.New(db)
	kategoriUsecaseFactory := kategoriUsecase.New(kategoriDataFactory)
	kategoriDelivery.New(e, kategoriUsecaseFactory)

	junk_stationDataFactory := junk_stationData.New(db)
	junk_stationcaseFactory := junk_stationUsecase.NewLogic(junk_stationDataFactory)
	junk_stationDelivery.NewHandller(e, junk_stationcaseFactory)
	
	penjualanClientDataFactory := penjualanClientData.New(db)
	penjualanClientUsecaseFactory := penjualanClientUsecase.New(penjualanClientDataFactory)
	penjualanClientDelivery.New(e, penjualanClientUsecaseFactory)
}
