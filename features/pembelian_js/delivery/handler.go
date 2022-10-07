package delivery

import (
	pjs "rozhok/features/pembelian_js"
	"rozhok/middlewares"
	"rozhok/utils/helper"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PembelianHandler struct {
	PembelianInterface pjs.UsecaseInterface
}

func New(e *echo.Echo, data pjs.UsecaseInterface) {
	handler := &PembelianHandler{
		PembelianInterface: data,
	}
	e.POST("pembelian/junk-station", handler.CreatePembelian, middlewares.JWTMiddleware(), middlewares.IsJunkStationVerified)
	e.GET("pembelian/junk-station", handler.GetPembelian, middlewares.JWTMiddleware())
	e.PUT("pembelian/:id/junk-station", handler.PutPembelian, middlewares.JWTMiddleware(), middlewares.IsJunkStationVerified)
	e.DELETE("pembelian/:id/junk-station", handler.DeletePembelian, middlewares.JWTMiddleware(), middlewares.IsJunkStationVerified)
}

func (h *PembelianHandler) CreatePembelian(c echo.Context) error {
	idToken, _, _ := middlewares.ExtractToken(c)
	var PjsRequest PembelianRequest
	errBind := c.Bind(&PjsRequest)
	if errBind != nil {
		return c.JSON(400, helper.FailedResponseHelper(errBind.Error()))
	}

	if err := c.Validate(PjsRequest); err != nil {
		return c.JSON(400, helper.FailedResponseHelper(err.Error()))
	}

	var Core = ToCore(PjsRequest)
	Core.JunkStationID = idToken
	_, err := h.PembelianInterface.CreatePembelian(Core)
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper(err.Error()))
	}
	return c.JSON(200, helper.SuccessResponseHelper("Success Create pembelian"))
}

func (h *PembelianHandler) GetPembelian(c echo.Context) error {
	idToken, _, _ := middlewares.ExtractToken(c)

	result, err := h.PembelianInterface.GetPembelian(idToken)
	pembelianResult := []PembelianResponse{}
	for _, v := range result {
		pembelianResult = append(pembelianResult, ToResponse(v))
	}
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("error get data"))
	}
	return c.JSON(200, helper.SuccessDataResponseHelper("Success Get Data", pembelianResult))
}

func (h *PembelianHandler) PutPembelian(c echo.Context) error {
	idParam := c.Param("id")
	idConv, errConv := strconv.Atoi(idParam)
	if errConv != nil || idConv == 0 {
		return c.JSON(400, helper.FailedResponseHelper("error update by param"))
	}

	var pjsUpdate PembelianRequest
	errBind := c.Bind(&pjsUpdate)
	if errBind != nil {
		return c.JSON(400, helper.FailedResponseHelper("status bad request for update"))
	}

	if errValidation := c.Validate(pjsUpdate); errValidation != nil {
		return c.JSON(400, helper.FailedResponseHelper(errValidation.Error()))
	}

	row, err := h.PembelianInterface.PutPembelian(idConv, ToCore(pjsUpdate))
	if err != nil || row == 0 {
		return c.JSON(400, helper.FailedResponseHelper("Failed to update pembelian"))
	}
	return c.JSON(200, helper.SuccessResponseHelper("Success update data pembelian"))
}

func (h *PembelianHandler) DeletePembelian(c echo.Context) error {
	var pjsDelete PembelianRequest

	row, err := h.PembelianInterface.DeletePembelian(helper.ParamInt(c, "id"), ToCore(pjsDelete))
	if err != nil || row < 1 {
		return c.JSON(400, helper.FailedResponseHelper("failed delete pembelian"))
	}
	return c.JSON(200, helper.SuccessResponseHelper("Success delete pembelian"))
}
