package delivery

import (
	"fmt"
	"net/http"
	"rozhok/config"
	"rozhok/features/produk"
	"rozhok/middlewares"
	"rozhok/utils/helper"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type Delivery struct {
	produkUsecase produk.UsecaseInterface
}

func New(e *echo.Echo, usecase produk.UsecaseInterface) {
	handler := &Delivery{
		produkUsecase: usecase,
	}
	e.POST("product", handler.PostProduk, middlewares.JWTMiddleware(), middlewares.IsAdmin)
	e.PUT("product/:id", handler.UpdateProduk, middlewares.JWTMiddleware(), middlewares.IsAdmin)
	e.GET("products", handler.GetAllProduk)
	e.GET("product/:id", handler.GetProduk)
	e.DELETE("product/:id", handler.DeleteProduk, middlewares.JWTMiddleware(), middlewares.IsAdmin)

}

func (deliv *Delivery) PostProduk(c echo.Context) error {

	var dataRequest ProdukRequest

	errBind := c.Bind(&dataRequest)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("error binding data"))
	}

	errValidate := c.Validate(&dataRequest)
	if errValidate != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper(errValidate.Error()))
	}

	//upload file Image
	imageData, imageInfo, imageErr := c.Request().FormFile("image_url")
	if imageErr == http.ErrMissingFile || imageErr != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("Failed to get image"))
	}

	imageExtension, err_image_extension := helper.CheckFileExtension(imageInfo.Filename, config.ContentImage)
	if err_image_extension != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("Image extension error"))
	}

	// check image size
	err_image_size := helper.CheckFileSize(imageInfo.Size, config.ContentImage)
	if err_image_size != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("Image size error"))
	}

	// memberikan nama file
	imageName := "product" + "_" + time.Now().Format("2006-01-02 15:04:05") + "." + imageExtension

	image, errUploadImg := helper.UploadFileToS3(config.EventImages, imageName, config.ContentImage, imageData)

	if errUploadImg != nil {
		fmt.Println(errUploadImg)
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("failed to upload file"))
	}
	dataRequest.Image_url = image

	row, err := deliv.produkUsecase.CreateProduk(toCore(dataRequest))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("Failed to insert product"))
	}
	if row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("Failed to insert product"))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponseHelper("Inserting product is succes"))
}

func (deliv *Delivery) UpdateProduk(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)

	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("param must be a number"))
	}

	var dataUpdate ProdukRequest
	errBind := c.Bind(&dataUpdate)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("error binding data"))
	}

	row, err := deliv.produkUsecase.UpdateProduk(toCore(dataUpdate), idConv)
	if err != nil || row == 0 {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("Failed to update data"))
	}

	return c.JSON(http.StatusBadRequest, helper.SuccessResponseHelper("Updating data is succes"))

}

func (deliv *Delivery) GetAllProduk(c echo.Context) error {
	result, err := deliv.produkUsecase.GetAllProduk()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("Failed to get data"))
	}
	return c.JSON(http.StatusOK, helper.SuccessDataResponseHelper("Succes get data", fromCoreList(result)))
}

func (deliv *Delivery) GetProduk(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)

	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("param must be a number"))
	}
	result, err := deliv.produkUsecase.GetProduk(idConv)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("failed to get data"))
	}
	return c.JSON(http.StatusOK, helper.SuccessDataResponseHelper("succes get data", fromCore(result)))
}

func (deliv *Delivery) DeleteProduk(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)

	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("param must be a number"))
	}

	row, err := deliv.produkUsecase.DeleteProduk(idConv)
	if err != nil || row == 0 {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("Failed to delete product"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponseHelper("Deleting product is success"))
}
