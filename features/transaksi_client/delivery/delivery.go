package delivery

import (
	"net/http"
	transaksiclient "rozhok/features/transaksi_client"
	"rozhok/middlewares"
	"rozhok/utils/helper"

	"github.com/labstack/echo/v4"
)

type TransaksiClient struct {
	Usecase transaksiclient.TransaksiClientUsecase
}

func New(e *echo.Echo, usecase transaksiclient.TransaksiClientUsecase) {
	handler := TransaksiClient{
		Usecase: usecase,
	}

	e.GET("/transaksi/client", handler.GetAll, middlewares.JWTMiddleware(), middlewares.IsClient)
	e.GET("/transaksi/:id/client/:tipe_transaksi", handler.Get, middlewares.JWTMiddleware(), middlewares.IsClient)
	e.POST("/transaksi/client", handler.Post, middlewares.JWTMiddleware(), middlewares.IsClient)
	e.PUT("/transaksi/:id/client", handler.Put, middlewares.JWTMiddleware(), middlewares.IsClient)
}

func (d *TransaksiClient) GetAll(c echo.Context) error {
	uid, _, _ := middlewares.ExtractToken(c)

	transaksiCore := transaksiclient.Core{}
	transaksiCore.Client.ID = uint(uid)

	transaksiCore.Status = c.QueryParam("status")
	transaksiCore.StartDate = c.QueryParam("start_date")
	transaksiCore.EndDate = c.QueryParam("end_date")
	transaksiCore.TipeTransaksi = c.QueryParam("tipe_transaksi")

	transaksiResults, err := d.Usecase.GetAll(transaksiCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper(err.Error()))
	}

	responseList := []Response{}
	for _, transaksi := range transaksiResults {
		responseList = append(responseList, FromCore(transaksi))
	}

	return c.JSON(http.StatusOK, helper.SuccessDataResponseHelper("success get all data", responseList))
}

func (d *TransaksiClient) Get(c echo.Context) error {
	transaksiID := helper.ParamInt(c, "id")

	transaksiCore := transaksiclient.Core{}
	transaksiCore.ID = uint(transaksiID)
	transaksiCore.TipeTransaksi = c.Param("tipe_transaksi")

	transaksi, err := d.Usecase.Get(transaksiCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessDataResponseHelper("success get data", FromCore(transaksi)))
}

func (d *TransaksiClient) Post(c echo.Context) error {
	uid, _, _ := middlewares.ExtractToken(c)
	transaksiCore := transaksiclient.Core{}
	transaksiCore.Client.ID = uint(uid)

	_, err := d.Usecase.Create(transaksiCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper(err.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponseHelper("success insert data"))
}

func (d *TransaksiClient) Put(c echo.Context) error {
	transaksiID := helper.ParamInt(c, "id")
	transaksiCore := transaksiclient.Core{}
	transaksiCore.ID = uint(transaksiID)

	_, err := d.Usecase.Update(transaksiCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper(err.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponseHelper("success update data"))
}
