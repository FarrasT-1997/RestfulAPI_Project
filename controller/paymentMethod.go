package controller

import (
	"net/http"
	"restfulAPI/lib/database"
	"restfulAPI/models"

	"github.com/labstack/echo"
)

func AddPaymentMethod(c echo.Context) error {
	addPayment := models.PaymentMethod{}
	c.Bind(&addPayment)
	payment, err := database.InsertPaymentMethod(addPayment)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "error add",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"payment": payment,
	})
}

func GetPaymentMethod(c echo.Context) error {
	listOfPayment, err := database.GetAllPaymentMethod()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to retrive data.",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":         "succes retrive data.",
		"list of payment": listOfPayment,
	})
}
