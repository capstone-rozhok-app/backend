package delivery

import (
	pjs "rozhok/features/pembelian_js"
	"rozhok/middlewares"
	"rozhok/utils/helper"

	"github.com/labstack/echo/v4"
)

type PembelianHandler struct{
	PembelianInterface  pjs.UsecaseInterface
}

func New(e *echo.Echo, data pjs.UsecaseInterface)  {
	handler := &PembelianHandler{
		PembelianInterface: data,
	}
	e.POST("pembelian/junk-station", handler.CreatePembelian, middlewares.JWTMiddleware(), middlewares.IsJunkStation)
}

func (h *PembelianHandler) CreatePembelian(c echo.Context) error{
	idToken, _, _ := middlewares.ExtractToken(c)

	var PjsRequest PembelianRequest
	errBind := c.Bind(&PjsRequest)
	if errBind != nil {
		return c.JSON(400, helper.FailedResponseHelper("failed to bind"))
	}
	if err := c.Validate(PjsRequest); err != nil{
		return c.JSON(400, helper.FailedResponseHelper(err.Error()))
	}
	_, err := h.PembelianInterface.CreatePembelian(ToCore(PjsRequest), idToken)
	if err != nil{
		return c.JSON(400, helper.FailedResponseHelper("failed to create pembelian"))
	}
	return c.JSON(200, helper.SuccessResponseHelper("Success Create pembelian"))
}

func (h *PembelianHandler) GetPembelian(c echo.Context) error {
	idToken, _, _:=  middlewares.ExtractToken(c)

	result, err:= h.PembelianInterface.GetPembelian()
	pembelianResult := PembelianResponse{}

	for _, v := range result {
		if v.ID == idToken{
			pembelianResult.ID = idToken
			pembelianResult.Kategori = v.Kategori
		}
	}

	if err != nil || pembelianResult.ID < 1 {
		return c.JSON(400, helper.FailedResponseHelper("error get data"))
	}
	return c.JSON(200, helper.SuccessDataResponseHelper("Success Get Data", pembelianResult))
}