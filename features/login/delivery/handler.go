package delivery

import (
	"net/http"
	"rozhok/features/login"
	"rozhok/utils/helper"

	"github.com/labstack/echo/v4"
)

type Delivery struct {
	authUsecase login.UsecaseInterface
}

func New(e *echo.Echo, usecase login.UsecaseInterface) {

	handler := Delivery{
		authUsecase: usecase,
	}

	e.POST("/login", handler.Auth)
	e.GET("/admin", handler.GetUsers)

}

func (delivery *Delivery) Auth(c echo.Context) error {

	var req Request
	errBind := c.Bind(&req)
	if errBind != nil {
		return c.JSON(400, helper.FailedResponseHelper("wrong request"))
	}

	str, role, username, status := delivery.authUsecase.LoginAuthorized(req.Email, req.Password)
	if str == "please input email and password" || str == "email not found" || str == "wrong password" {
		return c.JSON(400, helper.FailedResponseHelper(str))
	} else if str == "failed to created token" {
		return c.JSON(500, helper.FailedResponseHelper(str))
	} else {
		return c.JSON(200, helper.SuccessDataResponseHelper("Login Success", FromCore(str, role, username, status)))
	}

}

func (delivery *Delivery) GetUsers(c echo.Context) error {
	result, err := delivery.authUsecase.GetUsers()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("Failed to get data"))
	}
	return c.JSON(http.StatusOK, helper.SuccessDataResponseHelper("Succes get data", fromCore(result)))
}
