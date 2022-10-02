package delivery

import (
	js "rozhok/features/junk_station"
	"rozhok/middlewares"

	"github.com/labstack/echo/v4"
)

type JunkHandler struct {
	JunkInterface js.UsecaseInterface
}

func NewHandller(handler js.UsecaseInterface) *JunkHandler {
	return &JunkHandler{
		JunkInterface: handler,
	}
}

func (h *JunkHandler) CreateJunkStation(c echo.Context) error {
	jsID, _, _ := middlewares.ExtractToken(c)

	var JsRequest JsReq
	errBind := c.Bind(&JsRequest)
	if errBind != nil {
		return c.JSON(400, helper)
	}
}
