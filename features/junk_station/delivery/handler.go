package delivery

import (
	js "rozhok/features/junk_station"
	"rozhok/middlewares"
	"rozhok/utils/helper"
	"strconv"

	"github.com/labstack/echo/v4"
)

type JunkHandler struct {
	JunkInterface js.UsecaseInterface
}

func NewHandller(e *echo.Echo, data js.UsecaseInterface) {
	handler := &JunkHandler{
		JunkInterface: data,
	}
	e.GET("junk-station", handler.GetJunkStationAll, middlewares.JWTMiddleware())
	e.POST("junk-station", handler.CreateJunkStation)
	e.GET("junk-station/:id", handler.GetJunkStationById, middlewares.JWTMiddleware(), middlewares.IsJunkStation)
	e.PUT("junk-station/:id", handler.PutJunkStation, middlewares.JWTMiddleware(), middlewares.IsJunkStation)
	e.PUT("kemitraan/:id", handler.PutKemitraan, middlewares.JWTMiddleware(), middlewares.IsAdmin)
}

func (h *JunkHandler) CreateJunkStation(c echo.Context) error {
	var JsRequest JsReq
	errBind := c.Bind(&JsRequest)
	if errBind != nil {
		return c.JSON(400, helper.FailedResponseHelper(errBind.Error()))
	}

	errValidate := c.Validate(&JsRequest)
	if errValidate != nil {
		return c.JSON(400, helper.FailedResponseHelper("failed to bind"))
	}

	row, err := h.JunkInterface.CreateJunkStation(ToCoreReq(JsRequest))
	if err != nil || row != 1 {
		return c.JSON(400, helper.FailedResponseHelper(err.Error()))
	}
	return c.JSON(200, helper.SuccessResponseHelper("Succses create Junk Station"))
}

func (h *JunkHandler) GetJunkStationAll(c echo.Context) error {
	Provinsi := c.QueryParam("provinsi")
	Kota := c.QueryParam("kota")
	Kecamatam := c.QueryParam("kecamatan")

	var JunkFilter js.Core

	JunkFilter.Provinsi = Provinsi
	JunkFilter.Kota = Kota
	JunkFilter.Kecamatan = Kecamatam
	res, err := h.JunkInterface.GetJunkStationAll(JunkFilter)
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("failed to get data"))
	}
	return c.JSON(200, helper.SuccessDataResponseHelper(("succses get data"), CoreList(res)))
}

func (h *JunkHandler) GetJunkStationById(c echo.Context) error {
	idParam := c.Param("id")
	idConv, errConv := strconv.Atoi(idParam)
	if errConv != nil {
		return c.JSON(400, helper.FailedResponseHelper("error get by param"))
	}
	result, err := h.JunkInterface.GetJunkStationById(idConv)
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("error Get data"))
	}
	return c.JSON(200, helper.SuccessDataResponseHelper(("Succses get data"), FromCoreToResponse(result)))
}

func (h *JunkHandler) PutJunkStation(c echo.Context) error {
	idUser, _, _ := middlewares.ExtractToken(c)

	idParam := c.Param("id")
	idConv, errConv := strconv.Atoi(idParam)
	if errConv != nil || idConv == 0 {
		return c.JSON(400, helper.FailedResponseHelper("error update by param"))
	}

	if idConv != idUser {
		return c.JSON(403, helper.FailedResponseHelper("forbidden update"))
	}

	var JunkRequest JsReq
	errBind := c.Bind(&JunkRequest)
	if errBind != nil {
		return c.JSON(400, helper.FailedResponseHelper("error bind JS"))
	}
	row, err := h.JunkInterface.PutJunkStation(idConv, ToCoreReq(JunkRequest))
	if err != nil || row == 0 {
		return c.JSON(400, helper.FailedResponseHelper("error Update JS"))
	}
	return c.JSON(200, helper.SuccessResponseHelper("Succses update JS"))
}

func (h *JunkHandler) PutKemitraan(c echo.Context) error {
	var mitraUpdate JsReq

	errBind := c.Bind(&mitraUpdate)
	if errBind != nil{
		return c.JSON(400, helper.FailedResponseHelper("error bind data"))
	}

	row, err := h.JunkInterface.PutKemitraan(helper.ParamInt(c, "id"), ToCoreMitra(mitraUpdate))
	if err != nil || row == 0{
		return c.JSON(400, helper.FailedResponseHelper("failed to update kemitraan"))
	}
	return c.JSON(200, helper.SuccessResponseHelper("Succses Update data"))
}
