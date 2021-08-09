package database

import (
	"restfulAPI/config"
	"restfulAPI/middlewares"
	"restfulAPI/models"
)

func SignUp(user models.User) (interface{}, error) {
	if err := config.DB.Save(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

////////////////////////////////////////////////////////

func CreateUser(user models.User) (interface{}, error) {
	if err := config.DB.Save(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func LoginUsers(email, password string) (interface{}, error) {
	var err error
	var user models.User
	if err = config.DB.Where("email=? AND password=?", email, password).First(&user).Error; err != nil {
		return nil, err
	}

	user.Token, err = middlewares.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	}
	if err := config.DB.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func GetDetailUser(id int) (interface{}, error) {
	var user models.User
	if err := config.DB.Find(&user, id).Error; err != nil {
		return user, err
	}
	return user, nil
}

func EditUser(user models.User) (interface{}, error) {
	if err := config.DB.Save(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func GetOneUser(id int) (models.User, error) {
	var user models.User
	if err := config.DB.Find(&user, "id=?", id).Error; err != nil {
		return user, err
	}
	return user, nil
}
