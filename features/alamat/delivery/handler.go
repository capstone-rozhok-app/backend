package delivery

import (
	"net/http"
	"rozhok/features/alamat"
	"rozhok/middlewares"
	"rozhok/utils/helper"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Delivery struct {
	addressUsecase alamat.UsecaseInterface
}

func New(e *echo.Echo, usecase alamat.UsecaseInterface) {
	handler := &Delivery{
		addressUsecase: usecase,
	}
	e.POST("alamat", handler.PostAlamat, middlewares.JWTMiddleware(), middlewares.IsClient)
	e.PUT("alamat/:id", handler.UpdateAdress, middlewares.JWTMiddleware(), middlewares.IsClient)
	e.DELETE("alamat/:id", handler.DeleteAddress, middlewares.JWTMiddleware(), middlewares.IsClient)
	e.GET("alamats", handler.GetAllAddress, middlewares.JWTMiddleware())
	e.GET("alamats/:id", handler.GetAddress)
}

func (deliv *Delivery) PostAlamat(c echo.Context) error {

	var dataRequest AlamatRequest
	userId, _, _ := middlewares.ExtractToken(c)
	dataRequest.UserId = uint(userId)

	errBind := c.Bind(&dataRequest)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("error binding data"))
	}

	errValidate := c.Validate(&dataRequest)
	if errValidate != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper(errValidate.Error()))
	}

	row, err := deliv.addressUsecase.CreateAddress(toCore(dataRequest))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("Failed to insert address"))
	}
	if row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("Failed to insert address"))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponseHelper("Inserting address is succesfull"))
}

func (deliv *Delivery) UpdateAdress(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)

	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("param must be a number"))
	}
	userId, _, _ := middlewares.ExtractToken(c)

	var dataUpdate AlamatRequest
	errBind := c.Bind(&dataUpdate)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("error binding data"))
	}

	row, err := deliv.addressUsecase.PutAddress(toCore(dataUpdate), idConv, userId)
	if err != nil || row == 0 {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("Failed to update address"))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponseHelper("Updating data is succesfull"))

}

func (deliv *Delivery) DeleteAddress(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)

	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("param must be a number"))
	}
	userId, _, _ := middlewares.ExtractToken(c)

	row, err := deliv.addressUsecase.DeleteAddress(idConv, userId)
	if err != nil || row == 0 {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("Failed to delete address"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponseHelper("Deleting data is succesfull"))
}

func (deliv *Delivery) GetAllAddress(c echo.Context) error {
	userId, _, _ := middlewares.ExtractToken(c)
	result, err := deliv.addressUsecase.GetAllAddress(userId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("Failed to get data"))
	}
	return c.JSON(http.StatusOK, helper.SuccessDataResponseHelper("Succes get data", fromCoreList(result)))
}

func (deliv *Delivery) GetAddress(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)

	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("param must be a number"))
	}

	result, err := deliv.addressUsecase.GetAddress(idConv)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("Failed to get data"))
	}
	return c.JSON(http.StatusOK, helper.SuccessDataResponseHelper("Succes get data", fromCore(result)))
}
