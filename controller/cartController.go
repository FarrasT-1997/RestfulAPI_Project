package controller

import (
	"net/http"
	"restfulAPI/lib/database"
	"restfulAPI/models"
	"strconv"

	"github.com/labstack/echo"
)

func MakeCartID(c echo.Context) error {
	transactionId, err := strconv.Atoi(c.Param("transactionId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	productId, err := strconv.Atoi(c.Param("productId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}

	auth, user := Authorized(c)
	transaction, _ := database.GetOneTransaction(transactionId)
	if auth == false || user.FullName != transaction.Users {
		return echo.NewHTTPError(http.StatusUnauthorized, "Cannot access this account")
	}

	if check := database.CheckCart(transactionId, productId); check == false {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "product already added",
		})
	}
	productSelected, err := database.SelectProduct(productId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "invalid id",
		})
	}
	cart := models.ShoppingCart{}
	cart.Price = productSelected.Price
	cart.Quantity = 1
	cart.TransactionID = uint(transactionId)
	cart.ProductID = productSelected.ID
	cart.ProductName = productSelected.Name
	c.Bind(&cart)

	cartAdded, err := database.SaveProduct(cart)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot add product",
		})
	}
	c.Bind(&transaction)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    cartAdded,
	})
}

func ChangeQuantity(c echo.Context) error {
	transactionId, err := strconv.Atoi(c.Param("transactionId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	cartId, err := strconv.Atoi(c.Param("cartId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	quantity, err := strconv.Atoi(c.Param("quantity"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid number",
		})
	}
	cartSelected, err := database.SelectCart(cartId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id cart",
		})
	}

	auth, user := Authorized(c)
	transaction, _ := database.GetOneTransaction(transactionId)
	if auth == false || user.FullName != transaction.Users || cartSelected.TransactionID != transaction.ID {
		return echo.NewHTTPError(http.StatusUnauthorized, "Cannot access this account")
	}
	cartSelected.Quantity = quantity
	cartSelected.Price = cartSelected.Price * quantity
	c.Bind(&cartSelected)
	cartSave, err := database.EditCart(cartSelected)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot edit product",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    cartSave,
	})
}

func DeleteCart(c echo.Context) error {
	transactionId, err := strconv.Atoi(c.Param("transactionId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	cartId, err := strconv.Atoi(c.Param("cartId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	cartSelected, err := database.SelectCart(cartId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id cart",
		})
	}
	auth, user := Authorized(c)
	transaction, _ := database.GetOneTransaction(transactionId)
	if auth == false || user.FullName != transaction.Users || cartSelected.TransactionID != transaction.ID {
		return echo.NewHTTPError(http.StatusUnauthorized, "Cannot access this account")
	}
	deletedCart, err := database.DeleteCart(cartId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot delete cart",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    deletedCart,
	})
}
