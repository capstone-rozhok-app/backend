package delivery

import (
	mitra "rozhok/features/kemitraan"
	"rozhok/middlewares"
	"rozhok/utils/helper"

	"github.com/labstack/echo/v4"
)

type MitraHandler struct{
	MitraInterface mitra.UsecaseInterface
}

func New(e *echo.Echo, data mitra.UsecaseInterface)  {
	handler := &MitraHandler{
		MitraInterface: data,
	}
	e.PUT("kemitraan/:id", handler.PutKemitraan, middlewares.JWTMiddleware(), middlewares.IsAdmin)
}

func (h *MitraHandler) PutKemitraan(c echo.Context) error {
	var mitraUpdate MitraRequest

	errBind := c.Bind(&mitraUpdate)
	if errBind != nil{
		return c.JSON(400, helper.FailedResponseHelper("error bind data"))
	}

	if errValidation := c.Validate(mitraUpdate); errValidation != nil{
		return c.JSON(400, helper.FailedResponseHelper(errBind.Error()))
	}

	row, err := h.MitraInterface.PutKemitraan(helper.ParamInt(c, "id"), ToCore(mitraUpdate))
	if err != nil || row == 0{
		return c.JSON(400, helper.FailedResponseHelper("failed to update kemitraan"))
	}
	return c.JSON(200, helper.SuccessResponseHelper("Succses Update data"))
}