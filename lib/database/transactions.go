package database

import (
	"restfulAPI/config"
	"restfulAPI/models"
)

func CreateTransaction(transaction models.Transaction) (interface{}, error) {

	if err := config.DB.Save(&transaction).Error; err != nil {
		return nil, err
	}
	return transaction, nil
}
