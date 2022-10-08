package delivery

import (
	"net/http"
	"rozhok/config"
	js "rozhok/features/junk_station"
	"rozhok/middlewares"
	"rozhok/utils/helper"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type JunkHandler struct {
	JunkInterface js.UsecaseInterface
}

func NewHandller(e *echo.Echo, data js.UsecaseInterface) {
	handler := &JunkHandler{
		JunkInterface: data,
	}
	e.GET("junk-station/dashboard", handler.Dashboard, middlewares.JWTMiddleware(), middlewares.IsJunkStation)
	e.GET("junk-station", handler.GetJunkStationAll, middlewares.JWTMiddleware())
	e.POST("junk-station", handler.CreateJunkStation)
	e.GET("junk-station/profile", handler.GetJunkStationByToken, middlewares.JWTMiddleware(), middlewares.IsJunkStation)
	e.GET("junk-station/:id", handler.GetJunkStationById, middlewares.JWTMiddleware())
	e.PUT("junk-station/:id", handler.PutJunkStation, middlewares.JWTMiddleware(), middlewares.IsJunkStation)
	e.PUT("kemitraan/:id", handler.PutKemitraan, middlewares.JWTMiddleware(), middlewares.IsAdmin)
}

func (h *JunkHandler) Dashboard(c echo.Context) error {
	uid, _, _ := middlewares.ExtractToken(c)
	core := js.Core{
		JunkStationID:  uid,
		FilterPeriodic: c.QueryParam("filter"),
	}

	grandTotal, err := h.JunkInterface.Dashboard(core)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessDataResponseHelper("success get dashboard", map[string]interface{}{
		"total_pembelian": grandTotal,
	}))
}

func (h *JunkHandler) CreateJunkStation(c echo.Context) error {
	var JsRequest JsReq
	errBind := c.Bind(&JsRequest)
	if errBind != nil {
		return c.JSON(400, helper.FailedResponseHelper(errBind.Error()))
	}

	imageData, ImageInfo, ImageErr := c.Request().FormFile("foto")
	if ImageErr == http.ErrMissingFile || ImageErr != nil {
		return c.JSON(400, helper.FailedResponseHelper("failed get image"))
	}

	imageExtension, errImageExtension := helper.CheckFileExtension(ImageInfo.Filename, config.ContentImage)
	if errImageExtension != nil {
		return c.JSON(400, helper.FailedResponseHelper("your extension is illegal format"))
	}

	errImageSize := helper.CheckFileSize(ImageInfo.Size, config.ContentImage)
	if errImageSize != nil {
		return c.JSON(400, helper.FailedResponseHelper("size not up to standard"))
	}

	imageName := "junkstation" + "_" + time.Now().Format("2006-01-02 15:04:05") + "." + imageExtension
	image, errUploadImg := helper.UploadFileToS3(config.EventImages, imageName, config.ContentImage, imageData)
	if errUploadImg != nil {
		return c.JSON(400, helper.FailedResponseHelper("image can't be upload"))
	}

	core := ToCoreReq(JsRequest)
	core.Image_url = image

	row, err := h.JunkInterface.CreateJunkStation(core)
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
	JunkFilter.StatusKemitraan = c.QueryParam("status_kemitraan")

	JunkFilter.Provinsi = Provinsi
	JunkFilter.Kota = Kota
	JunkFilter.Kecamatan = Kecamatam
	res, err := h.JunkInterface.GetJunkStationAll(JunkFilter)
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("failed to get data"))
	}
	return c.JSON(200, helper.SuccessDataResponseHelper(("succses get data"), CoreList(res)))
}

func (h *JunkHandler) GetJunkStationByToken(c echo.Context) error {
	idConv, _, _ := middlewares.ExtractToken(c)
	result, err := h.JunkInterface.GetJunkStationById(idConv)
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("error Get data"))
	}
	return c.JSON(200, helper.SuccessDataResponseHelper(("Succses get data"), FromCoreToResponse(result)))
}

func (h *JunkHandler) GetJunkStationById(c echo.Context) error {
	id := helper.ParamInt(c, "id")
	result, err := h.JunkInterface.GetJunkStationById(id)
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
	row, err := h.JunkInterface.PutKemitraan(helper.ParamInt(c, "id"))
	if err != nil || row == 0 {
		return c.JSON(400, helper.FailedResponseHelper("failed to update kemitraan"))
	}
	return c.JSON(200, helper.SuccessResponseHelper("Succses Update data"))
}
