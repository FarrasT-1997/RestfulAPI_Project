package database

import (
	"fmt"
	"restfulAPI/config"
	"restfulAPI/models"
)

func CreateTransaction(transaction models.Transaction) (interface{}, error) {

	if err := config.DB.Save(&transaction).Error; err != nil {
		return nil, err
	}
	return transaction, nil
}

func GetTransactionById(name string) (interface{}, error) {
	var transaction []models.Transaction
	if err := config.DB.Find(&transaction, "users=?", name).Error; err != nil {
		return transaction, err
	}
	fmt.Println(transaction)
	return transaction, nil
}
