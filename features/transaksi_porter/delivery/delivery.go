package delivery

import (
	"net/http"
	transaksiporter "rozhok/features/transaksi_porter"
	"rozhok/middlewares"
	"rozhok/utils/helper"

	"github.com/labstack/echo/v4"
)

type TransaksiPorter struct {
	Usecase transaksiporter.TransaksiPorterUsecase
}

func New(e *echo.Echo, usecase transaksiporter.TransaksiPorterUsecase) {
	handler := &TransaksiPorter{
		Usecase: usecase,
	}

	e.GET("/transaksi/porter", handler.GetAll, middlewares.JWTMiddleware(), middlewares.IsPorter)
	e.GET("/transaksi/:transaction_id/porter", handler.Get, middlewares.JWTMiddleware(), middlewares.IsPorter)
	e.PUT("/transaksi/:transaction_id/porter", handler.PutTransaksiPembelian, middlewares.JWTMiddleware(), middlewares.IsPorter)
	e.POST("/transaksi/:transaction_id/porter", handler.PostTransaksiPenjualan, middlewares.JWTMiddleware(), middlewares.IsPorter)
}

func (deliv *TransaksiPorter) GetAll(c echo.Context) error {
	startDate := c.QueryParam("start_date")
	endDate := c.QueryParam("end_date")
	tipeTransaksi := c.QueryParam("type_transaction")
	idPorter, _, _ := middlewares.ExtractToken(c)

	var TransaksiPorterCore transaksiporter.Core

	TransaksiPorterCore.StartDate = startDate
	TransaksiPorterCore.EndDate = endDate
	TransaksiPorterCore.TipeTransaksi = tipeTransaksi
	TransaksiPorterCore.PorterID = uint(idPorter)

	transaksiList, err := deliv.Usecase.GetAll(TransaksiPorterCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper(err.Error()))
	}

	transaksiPorterResponses := []Response{}
	for _, transaksi := range transaksiList {
		transaksiPorterResponses = append(transaksiPorterResponses, toResponse(transaksi))
	}

	return c.JSON(http.StatusOK, helper.SuccessDataResponseHelper("success get data", transaksiPorterResponses))
}

func (deliv *TransaksiPorter) Get(c echo.Context) error {
	id := helper.ParamInt(c, "transaction_id")
	tipeTransaksi := c.QueryParam("type_transaction")
	idPorter, _, _ := middlewares.ExtractToken(c)

	if tipeTransaksi == "" {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("query parameter type_transaction is required"))
	}

	var TransaksiPorterCore transaksiporter.Core
	TransaksiPorterCore.ID = uint(id)
	TransaksiPorterCore.TipeTransaksi = tipeTransaksi
	TransaksiPorterCore.PorterID = uint(idPorter)

	transaksiPorter, err := deliv.Usecase.Get(TransaksiPorterCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessDataResponseHelper("success get data", toResponse(transaksiPorter)))
}

func (deliv *TransaksiPorter) PostTransaksiPenjualan(c echo.Context) error {
	id := helper.ParamInt(c, "transaction_id")
	idPorter, _, _ := middlewares.ExtractToken(c)

	var TransaksiPorterCore transaksiporter.Core
	TransaksiPorterCore.ID = uint(id)
	TransaksiPorterCore.PorterID = uint(idPorter)

	row, err := deliv.Usecase.PostTransaksiPenjualan(TransaksiPorterCore)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper(err.Error()))
	}

	if row < 1 {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("failed to create data"))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponseHelper("success create data"))
}

func (deliv *TransaksiPorter) PutTransaksiPembelian(c echo.Context) error {
	id := helper.ParamInt(c, "transaction_id")
	var TransaksiRequestList ArrayOfReq
	var TransaksiPorterCore transaksiporter.Core

	TransaksiPorterCore.ID = uint(id)

	errBind := c.Bind(&TransaksiRequestList)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper(errBind.Error()))
	}

	BarangRosokList := []transaksiporter.DetailTransaksiPorter{}
	for _, barangRosok := range TransaksiRequestList.BarangRosok {
		BarangRosokList = append(BarangRosokList, toCore(barangRosok))
	}
	TransaksiPorterCore.DetailTransaksiPorter = BarangRosokList

	row, err := deliv.Usecase.PutTransaksiPembelian(TransaksiPorterCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper(err.Error()))
	}

	if row < 1 {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("failed to update data"))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponseHelper("success update data"))
}
