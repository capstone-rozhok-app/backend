package delivery

import (
	"net/http"
	"rozhok/features/kategori"
	"rozhok/middlewares"
	"rozhok/utils/helper"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Delivery struct {
	kategoriUsecase kategori.UsecaseInterface
}

func New(e *echo.Echo, usecase kategori.UsecaseInterface) {
	handler := &Delivery{
		kategoriUsecase: usecase,
	}
	e.POST("category", handler.PostKategori, middlewares.JWTMiddleware(), middlewares.IsAdmin)
	e.PUT("category/:id", handler.UpdateKategori, middlewares.JWTMiddleware(), middlewares.IsAdmin)
	e.GET("categories", handler.GetAllKategori)
	e.DELETE("category/:id", handler.DeleteKategori, middlewares.JWTMiddleware(), middlewares.IsAdmin)
}

func (deliv *Delivery) PostKategori(c echo.Context) error {

	var dataRequest Request

	// errValidate := c.Validate(&dataRequest)
	// if errValidate != nil {
	// 	return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper(errValidate.Error()))
	// }
	errBind := c.Bind(&dataRequest)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("error binding data"))
	}

	row, err := deliv.kategoriUsecase.CreateKategori(toCore(dataRequest))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("Gagal membuat kategori"))
	}
	if row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("Gagal membuat kategori"))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponseHelper("Berhasil membuat kategori"))
}

func (deliv *Delivery) UpdateKategori(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)

	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("param must be a number"))
	}

	var dataUpdate Request
	errBind := c.Bind(&dataUpdate)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("error binding data"))
	}

	row, err := deliv.kategoriUsecase.UpdateKategori(toCore(dataUpdate), idConv)
	if err != nil || row == 0 {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("Gagal memperbarui data"))
	}

	return c.JSON(http.StatusBadRequest, helper.SuccessResponseHelper("Berhasil memperbarui data"))

}

func (deliv *Delivery) GetAllKategori(c echo.Context) error {
	result, err := deliv.kategoriUsecase.GetAllKategori()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("Failed to get data"))
	}
	return c.JSON(http.StatusOK, helper.SuccessDataResponseHelper("Succes get data", fromCoreList(result)))
}

func (deliv *Delivery) DeleteKategori(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)

	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("param must be a number"))
	}

	row, err := deliv.kategoriUsecase.DeleteKategori(idConv)
	if err != nil || row == 0 {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("Failed to get category"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponseHelper("Berhasil menghapus kategori"))
}
