package controller

import (
	"restfulAPI/modelsTest"

	"github.com/labstack/echo"
)

func CreateGetPaymentController(payment modelsTest.PaymentModel) echo.HandlerFunc {
	return func(c echo.Context) error {
		payment := payment.Get()
		return c.JSON(200, payment)
	}
}
