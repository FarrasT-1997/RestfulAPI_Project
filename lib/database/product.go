package database

import (
	"restfulAPI/config"
	"restfulAPI/models"
)

func GetAllCategories() (interface{}, error) {
	var category []models.Category
	result := config.DB.Find(&category)
	if err := result.Error; err != nil {
		return nil, err
	}
	return result, nil
}

func GetAllProductInSpecifiedCategory(namecategory int, category interface{}) (interface{}, error) {
	if err := config.DB.Find(&category, "namecategory=?", namecategory).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func GetSpecificProduct(nameproduct int) (interface{}, error) {
	var product models.Product
	if err := config.DB.Find(&product, "Name=?", nameproduct).Error; err != nil {
		return nil, err
	}
	return product, nil
}
