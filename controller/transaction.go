package controller

import (
	"fmt"
	"net/http"
	"restfulAPI/lib/database"
	"restfulAPI/middlewares"
	"restfulAPI/models"

	"github.com/labstack/echo"
)

func MakeTransaction(c echo.Context) error {
	transaction := models.Transaction{}
	userId := middlewares.ExtractTokenUserId(c)
	userList, err := database.GetOneUser(userId)
	fmt.Println(userList)
	userName := userList.FullName
	userAddress := userList.Address

	payment, err := database.GetAllPaymentTransaction()
	transaction.Users = userName
	transaction.Address = userAddress
	transaction.PaymentMethodID = 99
	transaction.PaymentMethod = payment[len(payment)-1]
	transaction.TransactionStatus = "waiting"
	transaction.TotalQuantity = 0
	transaction.TotalPrice = 0
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