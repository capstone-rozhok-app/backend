package delivery

import (
	"net/http"
	penjualanclient "rozhok/features/penjualan_client"
	"rozhok/middlewares"
	"rozhok/utils/helper"

	"github.com/labstack/echo/v4"
)

type PenjualanRosokClient struct {
	Usecase penjualanclient.PenjualanClientUsecase
}

func New(e *echo.Echo, usecase penjualanclient.PenjualanClientUsecase) {
	handler := &PenjualanRosokClient{
		Usecase: usecase,
	}

	e.GET("/penjualan/client", handler.GetAll, middlewares.JWTMiddleware(), middlewares.IsClient)
	e.POST("/penjualan/client", handler.Store, middlewares.JWTMiddleware(), middlewares.IsClient)
	e.PUT("/penjualan/:penjualan_id/client", handler.Update, middlewares.JWTMiddleware(), middlewares.IsClient)
	e.DELETE("/penjualan/:penjualan_id/client", handler.Delete, middlewares.JWTMiddleware(), middlewares.IsClient)
}

func (d *PenjualanRosokClient) GetAll(c echo.Context) error {
	clientId, _, _ := middlewares.ExtractToken(c)

	penjualanClientCore := penjualanclient.Core{}
	penjualanClientCore.ClientID = uint(clientId)

	penjualanClientList, err := d.Usecase.GetAll(penjualanClientCore)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper(err.Error()))
	}

	penjualanClientResponse := []Response{}
	for _, rosok := range penjualanClientList {
		penjualanClientResponse = append(penjualanClientResponse, ToResponse(rosok))
	}

	return c.JSON(http.StatusOK, helper.SuccessDataResponseHelper("success get all data", penjualanClientResponse))
}

func (d *PenjualanRosokClient) Update(c echo.Context) error {
	idPenjualan := helper.ParamInt(c, "penjualan_id")
	penjualanClientRequest := Request{}

	penjualanClientCore := penjualanclient.Core{}
	penjualanClientCore.ID = uint(idPenjualan)

	err := c.Bind(&penjualanClientRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper(err.Error()))
	}

	err = c.Validate(&penjualanClientRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper(err.Error()))
	}

	penjualanClientCore.KategoriID = penjualanClientRequest.IdKategori

	_, err = d.Usecase.Update(penjualanClientCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponseHelper("success update data"))
}

func (d *PenjualanRosokClient) Store(c echo.Context) error {
	idClient, _, _ := middlewares.ExtractToken(c)
	penjualanClientRequest := Request{}

	penjualanClientCore := penjualanclient.Core{}
	penjualanClientCore.ClientID = uint(idClient)

	err := c.Bind(&penjualanClientRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper(err.Error()))
	}

	err = c.Validate(&penjualanClientRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper(err.Error()))
	}

	penjualanClientCore.KategoriID = penjualanClientRequest.IdKategori

	_, err = d.Usecase.Store(penjualanClientCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponseHelper("success store data"))
}

func (d *PenjualanRosokClient) Delete(c echo.Context) error {
	idPenjualan := helper.ParamInt(c, "penjualan_id")

	penjualanClientCore := penjualanclient.Core{}
	penjualanClientCore.ID = uint(idPenjualan)

	_, err := d.Usecase.Delete(penjualanClientCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponseHelper("success delete data"))
}
