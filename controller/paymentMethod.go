package controller

import (
	"net/http"
	"restfulAPI/lib/database"

	"github.com/labstack/echo"
)

func GetPaymentMethod(c echo.Context) error {
	listOfPayment, err := database.GetAllPayment()
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
