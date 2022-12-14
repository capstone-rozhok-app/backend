package delivery

import (
	"net/http"
	"rozhok/features/porter"
	"rozhok/middlewares"

	"rozhok/utils/helper"

	"github.com/labstack/echo/v4"
)

type Delivery struct {
	porterUsecase porter.UsecaseInterface
}

func New(e *echo.Echo, usecase porter.UsecaseInterface) {
	handler := &Delivery{
		porterUsecase: usecase,
	}
	e.GET("/porter/dashboard", handler.Dashboard, middlewares.JWTMiddleware(), middlewares.IsPorter)
	e.GET("/porter/:id", handler.GetPorter, middlewares.JWTMiddleware(), middlewares.IsAdmin)
	e.GET("/porters", handler.GetAllPorter, middlewares.JWTMiddleware(), middlewares.IsAdmin)
	e.GET("/porter/:id/pendapatan", handler.GetPendapatan, middlewares.JWTMiddleware(), middlewares.IsAdmin)
	e.POST("/porter", handler.CreatePorter, middlewares.JWTMiddleware(), middlewares.IsAdmin)
	e.PUT("/porter/:id", handler.UpdatePorter, middlewares.JWTMiddleware(), middlewares.IsAdmin)
	e.DELETE("/porter/:id", handler.DeletePorter, middlewares.JWTMiddleware(), middlewares.IsAdmin)
}

func (deliv *Delivery) CreatePorter(c echo.Context) error {

	var dataRequest PorterRequest
	errBind := c.Bind(&dataRequest)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("error binding data"))
	}

	errValidate := c.Validate(&dataRequest)
	if errValidate != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper(errValidate.Error()))
	}

	row, err := deliv.porterUsecase.CreatePorter(toCore(dataRequest))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper(err.Error()))
	}
	if row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("failed to create porter"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponseHelper("success create porter"))
}

func (deliv *Delivery) UpdatePorter(c echo.Context) error {
	var dataRequest PorterRequest
	errBind := c.Bind(&dataRequest)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("error binding data"))
	}

	id := helper.ParamInt(c, "id")
	if id < 0 {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("parameter not valid"))
	}

	coreRequest := toCore(dataRequest)
	row, err := deliv.porterUsecase.UpdatePorter(coreRequest, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper(err.Error()))
	}
	if row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("failed to update porter"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponseHelper("success update porter"))
}

func (deliv *Delivery) DeletePorter(c echo.Context) error {
	id := helper.ParamInt(c, "id")
	if id < 0 {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("parameter not valid"))
	}

	row, err := deliv.porterUsecase.DeletePorter(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper(err.Error()))
	}
	if row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("failed to delete porter"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponseHelper("success delete porter"))
}

func (deliv *Delivery) GetAllPorter(c echo.Context) error {
	rows, err := deliv.porterUsecase.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper(err.Error()))
	}

	responsePorters := []PorterResponse{}
	for _, porter := range rows {
		responsePorters = append(responsePorters, fromCore(porter))
	}

	return c.JSON(http.StatusOK, helper.SuccessDataResponseHelper("success get all porter", responsePorters))
}

func (deliv *Delivery) GetPorter(c echo.Context) error {
	id := helper.ParamInt(c, "id")
	if id < 0 {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("parameter not valid"))
	}

	row, err := deliv.porterUsecase.Get(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessDataResponseHelper("success get porter", fromCore(row)))
}

func (deliv *Delivery) GetPendapatan(c echo.Context) error {
	id := helper.ParamInt(c, "id")
	if id < 0 {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("parameter not valid"))
	}

	var porterCore porter.Core

	porterCore.StartDate = c.QueryParam("start_date")
	porterCore.EndDate = c.QueryParam("end_date")
	porterCore.ID = uint(id)

	row, err := deliv.porterUsecase.GetPendapatan(porterCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessDataResponseHelper("success get porter report", PorterResponseReport{
		TotalLaba:      row.Laba,
		TotalPembelian: row.TotalPembelian,
		TotalPenjualan: row.TotalPenjualan,
	}))
}

func (deliv *Delivery) Dashboard(c echo.Context) error {
	id, _, _ := middlewares.ExtractToken(c)
	if id < 0 {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("not valid porter"))
	}

	var porterCore porter.Core

	porterCore.PeriodicFilter = c.QueryParam("filter")
	porterCore.ID = uint(id)

	row, err := deliv.porterUsecase.GetPendapatan(porterCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessDataResponseHelper("success get porter report", PorterResponseReport{
		TotalLaba:      row.Laba,
		TotalPembelian: row.TotalPembelian,
		TotalPenjualan: row.TotalPenjualan,
	}))
}
