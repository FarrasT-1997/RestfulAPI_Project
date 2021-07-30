package database

import (
	"restfulAPI/config"
	"restfulAPI/models"
)

func SignUp(user models.User) (interface{}, error) {
	if err := config.DB.Save(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
