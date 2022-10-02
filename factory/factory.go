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

}
