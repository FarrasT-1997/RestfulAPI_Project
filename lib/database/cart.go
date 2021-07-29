package database

func GetAllCartProduct(user_id int) (interface{}, error) {
	var cart models.Cart
	if err := config.DB.First(&cart, user_id); err != nil {
		return nil, err
	}
	return cart, nil
}

func InsertProductIntoCart(user_id int, product_id int) (interface{}, error) {
	var cart models.Cart
	if err := config.DB.Where(&cart, user_id).Save(&cart).Error; err != nil {
		return nil, err
	}
	return cart, nil
}

func DeleteProductInCart()
