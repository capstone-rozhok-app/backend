package delivery

import (
	"net/http"
	"rozhok/features/dbadmin"
	"rozhok/utils/helper"

	"github.com/labstack/echo/v4"
)

type Delivery struct {
	authUsecase dbadmin.UsecaseInterface
}

func New(e *echo.Echo, usecase dbadmin.UsecaseInterface) {

	handler := Delivery{
		authUsecase: usecase,
	}
	e.GET("/admin", handler.GetUsers)

}

func (delivery *Delivery) GetUsers(c echo.Context) error {
	result, err := delivery.authUsecase.GetUsers()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("Failed to get data"))
	}
	return c.JSON(http.StatusOK, helper.SuccessDataResponseHelper("Succes get data", fromCore(result)))
}
