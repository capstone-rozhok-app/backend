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

}
