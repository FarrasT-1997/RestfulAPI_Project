package database

func GetAllCategory(c. echo.Context) (interface, error) {
	category := models.Category
	result := config.DB.Find(&category)
	if err := result; err != nil {
		return nil, err
	}
	return result, nil
}

func GetAllProductInSpecifiedCategory(namecategory int, category interface{}) (interface{}, error) {
	if err := config.DB.Find(&category, "namecategory=?", namecategory); err != nil {
		return nil, err
	}
	return category, nil
}

func GetSpecificProduct(nameproduct int, product interface{}) (interface{}, error) {
	if err := config.DB.Find(&product) ; err != nil {
		return nil, error
	}
	return product, nil
}