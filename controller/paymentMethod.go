package controller

import (
	"net/http"
	"restfulAPI/lib/database"

	"github.com/labstack/echo"
)

func GetAllPaymentMethods(c echo.Context) error {

	category, err := database.GetAllPayment()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"users":   category,
	})
}
