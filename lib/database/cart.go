package database

import (
	"fmt"
	"restfulAPI/config"
	"restfulAPI/models"
)

func SelectProduct(productId int) (models.Product, error) {
	var product models.Product
	if err := config.DB.Find(&product, "id=?", productId).Error; err != nil {
		return product, err
	}
	return product, nil
}

func SaveProduct(cart models.ShoppingCart) (models.ShoppingCart, error) {
	if err := config.DB.Save(&cart).Error; err != nil {
		return cart, err
	}
	return cart, nil
}

func CheckCart(transactionId, productId int) bool {
	var cart models.ShoppingCart
	err := config.DB.Where("transaction_id=? AND product_id=?", transactionId, productId).First(&cart).Error
	fmt.Print(err)
	if err == nil {
		return false
	}
	return true
}

func SelectCart(id int) (models.ShoppingCart, error) {
	var cart models.ShoppingCart
	if err := config.DB.Find(&cart, "id=?", id).Error; err != nil {
		return cart, err
	}
	return cart, nil
}

func SelectAllCart(id int) ([]models.ShoppingCart, error) {
	var cart []models.ShoppingCart
	if err := config.DB.Find(&cart, "transaction_id=?", id).Error; err != nil {
		return cart, err
	}
	return cart, nil
}

func EditCart(cart models.ShoppingCart) (models.ShoppingCart, error) {
	if err := config.DB.Save(&cart).Error; err != nil {
		return cart, err
	}
	return cart, nil
}

func DeleteCart(id int) (interface{}, error) {
	var deletedCart models.ShoppingCart
	if err := config.DB.Find(&deletedCart, "id=?", id).Error; err != nil {
		return nil, err
	}

	if err := config.DB.Delete(&deletedCart).Error; err != nil {
		return nil, err
	}
	return deletedCart, nil
}
