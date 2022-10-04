package delivery

import (
	"net/http"
	"rozhok/features/client"
	"rozhok/middlewares"
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
	e.POST("/register", handler.PostClient)
	e.PUT("client", handler.UpdateClient, middlewares.JWTMiddleware(), middlewares.IsClient)
	e.DELETE("client", handler.DeleteAkun, middlewares.JWTMiddleware(), middlewares.IsClient)
	e.GET("client", handler.GetClient, middlewares.JWTMiddleware())
}

func (deliv *Delivery) PostClient(c echo.Context) error {

	var dataRequest ClientRequest
	dataRequest.Role = "client"

	errBind := c.Bind(&dataRequest)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("error binding data"))
	}

	errValidate := c.Validate(&dataRequest)
	if errValidate != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper(errValidate.Error()))
	}

	row, err := deliv.clientUsecase.CreateClient(toCore(dataRequest))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("Failed to ccreat account"))
	}
	if row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("Failed to ccreat account"))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponseHelper("Creating acccount is succes"))
}

func (deliv *Delivery) UpdateClient(c echo.Context) error {
	idClient, _, _ := middlewares.ExtractToken(c)

	var dataUpdate ClientRequest
	errBind := c.Bind(&dataUpdate)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("error binding data"))
	}

	row, err := deliv.clientUsecase.PutClient(toCore(dataUpdate), idClient)
	if err != nil || row == 0 {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("Failed to update data"))
	}

	return c.JSON(200, helper.SuccessResponseHelper("Updating data is succes"))

}

func (deliv *Delivery) DeleteAkun(c echo.Context) error {
	idUser, _, _ := middlewares.ExtractToken(c)

	row, err := deliv.clientUsecase.DeleteClient(idUser)
	if err != nil || row == 0 {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("Failed to delet account"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponseHelper("Deleting account is succes"))
}

func (deliv *Delivery) GetClient(c echo.Context) error {
	idUser, _, _ := middlewares.ExtractToken(c)

	result, err := deliv.clientUsecase.GetClient(idUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("Failed to get data"))
	}

	return c.JSON(http.StatusOK, helper.SuccessDataResponseHelper("Succes get data", fromCore(result)))
}
