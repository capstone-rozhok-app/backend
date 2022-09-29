package factory

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	clientData "rozhok/features/client/data"
	clientDelivery "rozhok/features/client/delivery"
	clientUsecase "rozhok/features/client/usecase"
	// loginData "rozhok/features/login/data"
	// loginDelivery "rozhok/features/login/delivery"
	// loginUsecase "rozhok/features/login/usecase"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	clientDataFactory := clientData.New(db)
	clientUsecaseFactory := clientUsecase.New(clientDataFactory)
	clientDelivery.New(e, clientUsecaseFactory)

	// loginDataFactory := loginData.New(db)
	// loginUsecaseFactory := loginUsecase.New(loginDataFactory)
	// loginDelivery.New(e, loginUsecaseFactory)

}
