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

func NewHandller(e *echo.Echo, data js.UsecaseInterface){
	handler := &JunkHandler{
		JunkInterface: data,
	}
	e.GET("junk-station", handler.GetJunkStationAll, middlewares.JWTMiddleware())
	e.POST("junk-station", handler.CreateJunkStation, middlewares.JWTMiddleware(), middlewares.IsJunkStation)
	e.GET("junk-station/:id", handler.GetJunkStationById, middlewares.JWTMiddleware(), middlewares.IsJunkStation)
	e.PUT("junk-station/:id", handler.PutJunkStation,middlewares.JWTMiddleware(),middlewares.IsJunkStation)
}

func (h *JunkHandler) CreateJunkStation(c echo.Context) error {
	jsID, _, _ := middlewares.ExtractToken(c)

	var JsRequest JsReq
	errBind := c.Bind(&JsRequest)
	if errBind != nil {
		return c.JSON(400, helper.FailedResponseHelper("failed to bind"))
	}
	if err := c.Validate(JsRequest); err != nil{
		return c.JSON(400, helper.FailedResponseHelper(err.Error()))
	}

	row, err := h.JunkInterface.CreateJunkStation(ToCoreReq(JsRequest), jsID)
	if err != nil || row != 1 {
		return c.JSON(400, helper.FailedResponseHelper("failed to create"))
	}
	return c.JSON(200, helper.SuccessDataResponseHelper("Succses create Junk Station", jsID))
}

func (h *JunkHandler) GetJunkStationAll(c echo.Context)error {
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

func (h *JunkHandler) GetJunkStationById(c echo.Context)error {
	idToken, _ ,_ := middlewares.ExtractToken(c)

	idParam := c.Param("id")
	idConv, errConv := strconv.Atoi(idParam)
	if errConv != nil {
		return c.JSON(400, helper.FailedResponseHelper("error get by param"))
	}
	result, err := h.JunkInterface.GetJunkStationById(idConv, idToken)
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("error Get data"))
	}
	return c.JSON(200, helper.SuccessDataResponseHelper(("Succses get data"), FromCoreToResponse(result)))
}

func (h *JunkHandler) PutJunkStation(c echo.Context)error  {
	idToken, _ ,_ := middlewares.ExtractToken(c)

	idParam := c.Param("id")
	idConv, errConv := strconv.Atoi(idParam)
	if errConv != nil || idConv == 0 {
		return c.JSON(400, helper.FailedResponseHelper("error update by param"))
	}


	var JunkRequest JsReq
	errBind := c.Bind(&JunkRequest)
	if errBind != nil {
		return c.JSON(400, helper.FailedResponseHelper("error bind JS"))
	}

	if err := c.Validate(JunkRequest); err != nil {
		return c.JSON(400, helper.FailedResponseHelper(err.Error()))
	}
	
	row, err := h.JunkInterface.PutJunkStation(idConv, ToCoreReq(JunkRequest))
	if err != nil || row == 0 {
		return c.JSON(400, helper.FailedResponseHelper("error Update JS"))
	}
	return c.JSON(200, helper.SuccessDataResponseHelper("Succses update JS", idToken))
}
