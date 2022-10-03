package delivery

import (
	"net/http"
	pengambilanrosok "rozhok/features/pengambilan_rosok"
	"rozhok/middlewares"
	"rozhok/utils/helper"

	"github.com/labstack/echo/v4"
)

type TransaksiPorter struct {
	Usecase pengambilanrosok.PengambilanRosokUsecase
}

func New(e *echo.Echo, usecase pengambilanrosok.PengambilanRosokUsecase) {
	handler := &TransaksiPorter{
		Usecase: usecase,
	}

	e.GET("/pengambilan/porter", handler.Get)
	e.GET("/pengambilan/:penjualan_id/porter", handler.GetAll)
	e.POST("/pengambilan/:penjualan_id/porter", handler.PostTransaksiPenjualan)
}

func (deliv *TransaksiPorter) GetAll(c echo.Context) error {
	idPorter, _, _ := middlewares.ExtractToken(c)

	var pengambilanBarangRosok pengambilanrosok.Core

	pengambilanBarangRosok.PorterID = uint(idPorter)

	transaksiList, err := deliv.Usecase.GetAll(pengambilanBarangRosok)
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
	idPorter, _, _ := middlewares.ExtractToken(c)

	var PengambilanRosokCore pengambilanrosok.Core
	PengambilanRosokCore.ID = uint(id)
	PengambilanRosokCore.PorterID = uint(idPorter)

	pengambilanRosokResult, err := deliv.Usecase.Get(PengambilanRosokCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessDataResponseHelper("success get data", toResponse(pengambilanRosokResult)))
}

func (deliv *TransaksiPorter) PostTransaksiPenjualan(c echo.Context) error {
	id := helper.ParamInt(c, "transaction_id")
	idPorter, _, _ := middlewares.ExtractToken(c)

	var PengambilanRosokCore pengambilanrosok.Core
	PengambilanRosokCore.ID = uint(id)
	PengambilanRosokCore.PorterID = uint(idPorter)

	row, err := deliv.Usecase.CreatePengambilanRosok(PengambilanRosokCore)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper(err.Error()))
	}

	if row < 1 {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("failed to create data"))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponseHelper("success create data"))
}
