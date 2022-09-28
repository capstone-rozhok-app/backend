package delivery

import (
	"net/http"
	"rozhok/features/client"
	"rozhok/utils/helper"

	"github.com/labstack/echo/v4"
)

type Delivery struct {
	clientUsecase client.UsecaseInterface
}

func New(e *echo.Echo, usecase client.UsecaseInterface) {
	handler := &Delivery{
		clientUsecase: usecase,
	}
	e.POST("/register/client", handler.PostClient)
}

func (deliv *Delivery) PostClient(c echo.Context) error {

	var dataRequest ClientRequest
	dataRequest.Role = "client"
	dataRequest.AlamatId = 1
	errBind := c.Bind(&dataRequest)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("error binding data"))
	}

	row, err := deliv.clientUsecase.CreateClient(toCore(dataRequest))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("error insert data"))
	}
	if row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("error insert data"))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponseHelper("success insert data"))
}
