package delivery

import (
	js "rozhok/features/junk_station"
	"rozhok/middlewares"
	"rozhok/utils/helper"

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

// func (h *JunkHandler) CreateJunkStation(c echo.Context) error {
// 	junkToken, errToken = middlewares.ExtractToken(c)
// 	if junkToken == 0 || errToken != nil {
// 		return c.JSON(400, helper.FailedResponseHelper("Token invalid !"))
// 	}
// }
