package delivery

import (
	"net/http"
	transaksijunkstation "rozhok/features/transaksi_junk_station"
	"rozhok/middlewares"
	"rozhok/utils/helper"

	"github.com/labstack/echo/v4"
)

type TransaksiJunkStation struct {
	Usecase transaksijunkstation.TransaksiJunkStationUsecase
}

func New(e *echo.Echo, usecase transaksijunkstation.TransaksiJunkStationUsecase) {
	handler := &TransaksiJunkStation{
		Usecase: usecase,
	}

	e.POST("/transaksi/junk-station", handler.Create, middlewares.JWTMiddleware(), middlewares.IsJunkStation)
	e.GET("/transaksi/junk-station", handler.GetAll, middlewares.JWTMiddleware(), middlewares.IsJunkStation)
	e.GET("/transaksi/:id/junk-station", handler.Get, middlewares.JWTMiddleware(), middlewares.IsJunkStation)

}

func (d *TransaksiJunkStation) GetAll(c echo.Context) error {
	id, _, _ := middlewares.ExtractToken(c)

	var Core transaksijunkstation.Core
	Core.UserID = uint(id)
	Core.StartDate = c.QueryParam("start_date")
	Core.EndDate = c.QueryParam("end_date")

	transaksiList, err := d.Usecase.GetAll(Core)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper(err.Error()))
	}

	transakiResponse := []Response{}
	for _, transaksi := range transaksiList {
		transakiResponse = append(transakiResponse, FromCore(transaksi))
	}

	return c.JSON(http.StatusOK, helper.SuccessDataResponseHelper("success get all data", transakiResponse))
}

func (d *TransaksiJunkStation) Get(c echo.Context) error {
	idTransaksi := helper.ParamInt(c, "id")
	var Core transaksijunkstation.Core

	Core.ID = uint(idTransaksi)

	transaksi, err := d.Usecase.Get(Core)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessDataResponseHelper("success get data", FromCore(transaksi)))
}

func (d *TransaksiJunkStation) Create(c echo.Context) error {
	id, _, _ := middlewares.ExtractToken(c)
	var Core transaksijunkstation.Core

	Core.UserID = uint(id)

	_, err := d.Usecase.Create(Core)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponseHelper("success insert data"))
}
