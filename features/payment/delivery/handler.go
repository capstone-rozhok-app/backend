package delivery

import (
	"net/http"
	"rozhok/features/payment"
	"rozhok/middlewares"
	"rozhok/utils/helper"

	"github.com/labstack/echo/v4"
)

type Payment struct {
	Usecase payment.PaymentUsecase
}

func New(e *echo.Echo, usecase payment.PaymentUsecase) {
	handler := Payment{
		Usecase: usecase,
	}
	e.POST("/payment", handler.PostPayment, middlewares.JWTMiddleware(), middlewares.IsClient)
}

func (h *Payment) PostPayment(c echo.Context) error {
	var paymentRequest Request
	uid, _, _ := middlewares.ExtractToken(c)
	err := c.Bind(&paymentRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper(err.Error()))
	}

	err = c.Validate(&paymentRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper(err.Error()))
	}

	paymentCore := payment.Core{
		Bank:  paymentRequest.Bank,
		Kurir: paymentRequest.Kurir,
		Client: payment.Client{
			ID: uint(uid),
		},
	}

	invoice, err := h.Usecase.Create(paymentCore)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessDataResponseHelper("success get data", FromCore(invoice)))
}
