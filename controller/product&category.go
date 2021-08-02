package controller

import (
	"net/http"
	"restfulAPI/lib/database"
	"strconv"

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

func GetAllSpecifiedProductCategory(c echo.Context) error {
	categoryId, err := strconv.Atoi(c.Param("idcategory"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	categoryList, err := database.GetAllProductInSpecifiedCategory(categoryId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot find product list",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    categoryList,
	})
}

func GetProductDetail(c echo.Context) error {
	productId, err := strconv.Atoi(c.Param("idproduct"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid product id",
		})
	}
	product, err := database.GetSpecificProduct(productId)
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
