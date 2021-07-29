package database


func SignUp() (interface{}, error) {
	var users []model.User
	if err := config.DB.Save(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func