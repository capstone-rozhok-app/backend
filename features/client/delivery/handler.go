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
	e.POST("/register/client", handler.PostClient)
	e.PUT("client", handler.UpdateClient, middlewares.JWTMiddleware())
	e.DELETE("client", handler.DeleteAkun, middlewares.JWTMiddleware())
	e.POST("/login/client", handler.Auth)
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
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("Gagal membuat akun"))
	}
	if row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("Gagal membuat akun"))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponseHelper("Berhasil membuat akun"))
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
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("Gagal memperbarui data"))
	}

	return c.JSON(http.StatusBadRequest, helper.SuccessResponseHelper("Berhasil memperbarui data"))

}

func (deliv *Delivery) DeleteAkun(c echo.Context) error {
	idUser, _, _ := middlewares.ExtractToken(c)

	row, err := deliv.clientUsecase.DeleteClient(idUser)
	if err != nil || row == 0 {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("Gagal menghapus akun"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponseHelper("Berhasil menghapus akun"))
}

func (deliv *Delivery) Auth(c echo.Context) error {

	var req AuthRequest
	errBind := c.Bind(&req)
	if errBind != nil {
		return c.JSON(400, helper.FailedResponseHelper("wrong request"))
	}

	str, role, username := deliv.clientUsecase.LoginAuthorized(req.Email, req.Password)
	if str == "email dan password tidak boleh kosong" || str == "email tidak ditemukan" || str == "password salah" {
		return c.JSON(400, helper.FailedResponseHelper(str))
	} else if str == "failed to created token" {
		return c.JSON(500, helper.FailedResponseHelper(str))
	} else {
		return c.JSON(200, helper.SuccessDataResponseHelper("Berhasil masuk", fromLoginCore(str, role, username)))
	}

}
