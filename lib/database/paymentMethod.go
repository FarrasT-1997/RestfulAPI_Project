package database

import (
	"restfulAPI/config"
	"restfulAPI/models"
)

func GetAllPaymentMethod() (interface{}, error) {
	var paymentmethod []models.PaymentMethod
	result := config.DB.Find(&paymentmethod)
	if err := result.Error; err != nil {
		return nil, err
	}
	return result, nil
}

func InsertPaymentMethod(paymentmethod models.PaymentMethod) (interface{}, error) {
	if err := config.DB.Save(&paymentmethod).Error; err != nil {
		return nil, err
	}
	return paymentmethod, nil
}
