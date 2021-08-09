package database

import (
	"restfulAPI/config"
	"restfulAPI/models"
)

func GetAllPayment() (interface{}, error) {
	var user []models.PaymentMethodTable
	if err := config.DB.Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func GetAllPaymentTransaction() ([]models.PaymentMethodTable, error) {
	var user []models.PaymentMethodTable
	if err := config.DB.Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func GetOnePayment(id int) (models.PaymentMethodTable, error) {
	var payment models.PaymentMethodTable
	if err := config.DB.Find(&payment, "id=?", id).Error; err != nil {
		return payment, err
	}
	return payment, nil
}
