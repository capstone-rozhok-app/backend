package delivery

import (
	"net/http"
	"rozhok/features/dbadmin"
	"rozhok/middlewares"
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
	e.GET("/admin", handler.GetUsers, middlewares.JWTMiddleware(), middlewares.IsAdmin)
	e.GET("/transaksi/admin", handler.GetTransaksi, middlewares.JWTMiddleware(), middlewares.IsAdmin)
	e.GET("/transaksi/:id/admin", handler.GetTransaksiDetail, middlewares.JWTMiddleware(), middlewares.IsAdmin)
	e.PUT("/transaksi/:id/admin", handler.PutTransaksi, middlewares.JWTMiddleware(), middlewares.IsAdmin)
}

func (delivery *Delivery) GetUsers(c echo.Context) error {
	result, err := delivery.authUsecase.GetUsers()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("Failed to get data"))
	}
	return c.JSON(http.StatusOK, helper.SuccessDataResponseHelper("Succes get data", fromCore(result)))
}

func (delivery *Delivery) GetTransaksiDetail(c echo.Context) error {
	idParam := helper.ParamInt(c, "id")
	TransaksiCore := dbadmin.TransaksiCore{
		ID: uint(idParam),
	}

	transaksiResult, err := delivery.authUsecase.GetTransaksiDetail(TransaksiCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessDataResponseHelper("success get data", FromCoreTransaksi(transaksiResult)))
}

func (delivery *Delivery) GetTransaksi(c echo.Context) error {
	transaksiCore := dbadmin.TransaksiCore{}

	transaksiCore.Status = c.QueryParam("status")
	transaksiCore.StartDate = c.QueryParam("start_date")
	transaksiCore.EndDate = c.QueryParam("end_date")

	transaksiResults, err := delivery.authUsecase.GetTransaksi(transaksiCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper(err.Error()))
	}

	responseList := []TransaksiResponse{}
	for _, transaksi := range transaksiResults {
		responseList = append(responseList, FromCoreTransaksi(transaksi))
	}

	return c.JSON(http.StatusOK, helper.SuccessDataResponseHelper("success get all data", responseList))
}

func (delivery *Delivery) PutTransaksi(c echo.Context) error {
	idParam := helper.ParamInt(c, "id")
	TransaksiCore := dbadmin.TransaksiCore{
		ID: uint(idParam),
	}

	err := delivery.authUsecase.UpdateTransaksi(TransaksiCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponseHelper("success update status"))
}
