package delivery

import (
	"net/http"
	"rozhok/features/cart"
	"rozhok/middlewares"
	"rozhok/utils/helper"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Delivery struct {
	cartUsecase cart.UsecaseInterface
}

func New(e *echo.Echo, usecase cart.UsecaseInterface) {
	handler := &Delivery{
		cartUsecase: usecase,
	}
	e.POST("cart", handler.PostCart, middlewares.JWTMiddleware(), middlewares.IsClient)
	e.GET("carts", handler.GetCarts, middlewares.JWTMiddleware())
	e.PUT("cart/:id", handler.UpdateCart, middlewares.JWTMiddleware(), middlewares.IsClient)
	e.DELETE("cart/:id", handler.DeleteCart, middlewares.JWTMiddleware(), middlewares.IsClient)

}

func (deliv *Delivery) PostCart(c echo.Context) error {
	var dataRequest CartRequest
	userId, _, _ := middlewares.ExtractToken(c)
	dataRequest.UserId = uint(userId)

	errBind := c.Bind(&dataRequest)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("error binding data"))
	}

	row, err := deliv.cartUsecase.CreateCart(toCore(dataRequest))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("Failed to insert address"))
	}
	if row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("Failed to insert address"))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponseHelper("Inserting address is succesfull"))
}

func (deliv *Delivery) GetCarts(c echo.Context) error {
	userId, _, _ := middlewares.ExtractToken(c)
	result, err := deliv.cartUsecase.GetAllCart(userId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("Failed to get data"))
	}
	return c.JSON(http.StatusOK, helper.SuccessDataResponseHelper("Succes get data", fromCoreList(result)))
}

func (deliv *Delivery) UpdateCart(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)

	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("param must be a number"))
	}
	userId, _, _ := middlewares.ExtractToken(c)

	var dataUpdate CartRequest
	errBind := c.Bind(&dataUpdate)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("error binding data"))
	}

	row, err := deliv.cartUsecase.UpdateCart(toCore(dataUpdate), idConv, userId)
	if err != nil || row == 0 {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("Failed to update address"))
	}

	return c.JSON(http.StatusBadRequest, helper.SuccessResponseHelper("Updating data is succesfull"))

}

func (deliv *Delivery) DeleteCart(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)

	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("param must be a number"))
	}
	userId, _, _ := middlewares.ExtractToken(c)

	row, err := deliv.cartUsecase.DeleteCart(idConv, userId)
	if err != nil || row == 0 {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("Failed to delete address"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponseHelper("Deleting data is succesfull"))
}
