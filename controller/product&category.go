package controller

import (
	"net/http"
	"restfulAPI/lib/database"
	"strconv"

	// "restfulAPI/middlewares"
	"github.com/labstack/echo"
)

func GetAllCategory(c echo.Context) error {
	category, err := database.GetAllCategories()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"users":   category,
	})
}

// func GetAllSpecifiedProductCategory(c echo.Context) error {
// 	categoryName, err := strconv.Atoi(c.Param("namecategory"))
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]interface{}{
// 			"message": "invalid id",
// 		})
// 	}
// }

func GetProductDetail(c echo.Context) error {
	nameProduct, err := strconv.Atoi(c.Param(":nameproduct"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid name",
		})
	}
	product, err := database.GetSpecificProduct(nameProduct)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot find product",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    product,
	})
}
