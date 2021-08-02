package controller

import (
	"net/http"
	"restfulAPI/lib/database"
	"restfulAPI/middlewares"
	"restfulAPI/models"
	"strconv"

	"github.com/labstack/echo"
)

func Authorized(c echo.Context) (bool, models.User) {
	userId, token := middlewares.ExtractTokenUserId(c)
	userList, _ := database.GetOneUser(userId)

	if userList.Token != token {
		return false, userList
	}
	return true, userList
}

func MakeTransaction(c echo.Context) error {
	auth, userList := Authorized(c)
	if auth == false {
		return echo.NewHTTPError(http.StatusUnauthorized, "Cannot access this account")
	}

	payment, err := database.GetAllPaymentTransaction()
	transaction := database.InputTransactioData(userList, payment)
	c.Bind(&transaction)
	transactionId, err := database.CreateTransaction(transaction)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot insert data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "New transaction added successfully",
		"data":    transactionId,
	})
}

func GetAllTransaction(c echo.Context) error {
	auth, userList := Authorized(c)
	if auth == false {
		return echo.NewHTTPError(http.StatusUnauthorized, "Cannot access this account")
	}
	transactionList, err := database.GetTransactionById(userList.FullName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot show transaction list",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    transactionList,
	})
}

func DeleteTransaction(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("transactionId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	auth, userList := Authorized(c)
	transactionList, err := database.GetOneTransaction(id)

	if auth == false || userList.FullName != transactionList.Users {
		return echo.NewHTTPError(http.StatusUnauthorized, "Cannot access this account")
	}

	if transactionList.Checkout == "success" {
		return echo.NewHTTPError(http.StatusBadRequest, "Cannot edit this transaction")
	}

	transaction, err := database.DeleteTransactions(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot delete data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    transaction,
	})
}

func ChangeStatus(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("transactionId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	auth, userList := Authorized(c)
	transaction, err := database.GetOneTransaction(id)

	if auth == false || userList.FullName != transaction.Users {
		return echo.NewHTTPError(http.StatusUnauthorized, "Cannot access this account")
	}

	if transaction.Checkout == "not yet" {
		return echo.NewHTTPError(http.StatusBadRequest, "Cannot edit this transaction")
	}

	transaction.TransactionStatus = "paid"
	c.Bind(&transaction)
	transactionSaved, err := database.EditTransaction(transaction)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot edit data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    transactionSaved,
	})
}

func ChangePaymentMethod(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("transactionId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	paymentId, err := strconv.Atoi(c.Param("paymentId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid payment id",
		})
	}
	payment, err := database.GetOnePayment(paymentId)
	auth, userList := Authorized(c)
	transaction, err := database.GetOneTransaction(id)

	if auth == false || userList.FullName != transaction.Users {
		return echo.NewHTTPError(http.StatusUnauthorized, "Cannot access this account")
	}

	if transaction.Checkout == "success" {
		return echo.NewHTTPError(http.StatusBadRequest, "Cannot edit this transaction")
	}

	transaction.PaymentMethod = payment
	c.Bind(&transaction)
	transactionSaved, err := database.EditTransaction(transaction)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot edit data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    transactionSaved,
	})
}

func TransactionCheckout(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("transactionId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	auth, userList := Authorized(c)
	transactionList, err := database.GetOneTransaction(id)

	if auth == false || userList.FullName != transactionList.Users {
		return echo.NewHTTPError(http.StatusUnauthorized, "Cannot access this account")
	}
	if transactionList.PaymentMethodID == 99 || transactionList.TotalPrice == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Please fill the transaction product and payment method")
	}
	transactionList.Checkout = "success"
	c.Bind(&transactionList)
	checkoutStatus, err := database.EditTransaction(transactionList)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Checkout Failure",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    checkoutStatus,
	})
}
