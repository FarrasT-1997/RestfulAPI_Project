package database

import (
	"restfulAPI/config"
	"restfulAPI/models"
)

func GetAllCategories() (interface{}, error) {
	var category []models.Category

	if err := config.DB.Find(&category).Error; err != nil {
		return nil, err
	}
	var product []models.Product

	var productDescription []models.ProductDescription
	if err := config.DB.Find(&productDescription).Error; err != nil {
		return nil, err
	}

	for i := 1; i <= len(category); i++ {
		var categoryList models.Category

		if err := config.DB.Where("id=?", i).First(&categoryList).Error; err != nil {
			return nil, err
		}

		if err := config.DB.Find(&product, "category_id=?", categoryList.ID).Error; err != nil {
			return nil, err
		}

		for _, products := range product {
			category[i-1].Products = append(category[i-1].Products, products)
			category[i-1].Products[len(category[i-1].Products)-1].ProductDescription.ID = productDescription[products.ID-1].ID
			category[i-1].Products[len(category[i-1].Products)-1].ProductDescription.ProductDescriptions = productDescription[products.ID-1].ProductDescriptions
			category[i-1].Products[len(category[i-1].Products)-1].ProductDescription.ProductID = productDescription[products.ID-1].ProductID
		}

	}
	return category, nil
}

func GetAllProductInSpecifiedCategory(idcategory int) (interface{}, error) {

	var category models.Category
	if err := config.DB.Where("id=?", idcategory).First(&category).Error; err != nil {
		return nil, err
	}
	var product []models.Product
	if err := config.DB.Find(&product, "category_id=?", category.ID).Error; err != nil {
		return nil, err
	}
	for i := 1; i <= len(product); i++ {
		var productList models.Product
		var productDescription models.ProductDescription

		if err := config.DB.Where("id=?", i).First(&productList).Error; err != nil {
			return nil, err
		}

		if err := config.DB.Find(&productDescription, "product_id=?", product[i-1].ID).Error; err != nil {
			return nil, err
		}
		product[i-1].ProductDescription.ID = productDescription.ID
		product[i-1].ProductDescription.ProductDescriptions = productDescription.ProductDescriptions
		product[i-1].ProductDescription.ProductID = productDescription.ProductID

	}
	return product, nil
}

func GetSpecificProduct(productId int) (interface{}, error) {
	var product []models.Product

	if err := config.DB.Where("id=?", productId).First(&product).Error; err != nil {
		return nil, err
	}
	var description models.ProductDescription
	if err := config.DB.Find(&description, "product_id=?", product[0].ID).Error; err != nil {
		return nil, err
	}
	product[0].ProductDescription.ID = description.ID
	product[0].ProductDescription.ProductDescriptions = description.ProductDescriptions
	product[0].ProductDescription.ProductID = description.ProductID

	return product, nil
}
