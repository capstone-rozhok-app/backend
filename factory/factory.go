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

	JSData "rozhok/features/junk_station/data"
	JSDelivery "rozhok/features/junk_station/delivery"
	JSUsecase "rozhok/features/junk_station/usecase"

	PJSData "rozhok/features/pembelian_js/data"
	PJSDelivery "rozhok/features/pembelian_js/delivery"
	PJSUsecase "rozhok/features/pembelian_js/usecase"

	kategoriData "rozhok/features/kategori/data"
	kategoriDelivery "rozhok/features/kategori/delivery"
	kategoriUsecase "rozhok/features/kategori/usecase"

	penjualanClientData "rozhok/features/penjualan_client/data"
	penjualanClientDelivery "rozhok/features/penjualan_client/delivery"
	penjualanClientUsecase "rozhok/features/penjualan_client/usecase"

	pengambilanRosok "rozhok/features/pengambilan_rosok/data"
	pengambilanRosokDelivery "rozhok/features/pengambilan_rosok/delivery"
	pengambilanRosokUsecase "rozhok/features/pengambilan_rosok/usecase"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	clientDataFactory := clientData.New(db)
	clientUsecaseFactory := clientUsecase.New(clientDataFactory)
	clientDelivery.New(e, clientUsecaseFactory)

	kategoriDataFactory := kategoriData.New(db)
	kategoriUsecaseFactory := kategoriUsecase.New(kategoriDataFactory)
	kategoriDelivery.New(e, kategoriUsecaseFactory)

	porterDataFactory := porterData.New(db)
	porterUsecaseFactory := porterUsecase.New(porterDataFactory)
	porterDelivery.New(e, porterUsecaseFactory)

	loginDataFactory := loginData.New(db)
	loginUsecaseFactory := loginUsecase.New(loginDataFactory)
	loginDelivery.New(e, loginUsecaseFactory)

	transaksiPorterDataFactory := transaksiPorterData.New(db)
	transaksiPorterUsecaseFactory := transaksiPorterUsecase.New(transaksiPorterDataFactory)
	transaksiPorterDelivery.New(e, transaksiPorterUsecaseFactory)

	JSDataFactory := JSData.New(db)
	JSUsecaseFactory := JSUsecase.NewLogic(JSDataFactory)
	JSDelivery.NewHandller(e, JSUsecaseFactory)

	PJSDataFactory := PJSData.New(db)
	PJSUsecaseFactory := PJSUsecase.New(PJSDataFactory)
	PJSDelivery.New(e, PJSUsecaseFactory)

	penjualanClientDataFactory := penjualanClientData.New(db)
	penjualanClientUsecaseFactory := penjualanClientUsecase.New(penjualanClientDataFactory)
	penjualanClientDelivery.New(e, penjualanClientUsecaseFactory)

	pengambilanRosokDataFactory := pengambilanRosok.New(db)
	pengambilanRosokUsecaseFactory := pengambilanRosokUsecase.New(pengambilanRosokDataFactory)
	pengambilanRosokDelivery.New(e, pengambilanRosokUsecaseFactory)
}
