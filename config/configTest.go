package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func CreateDBTest() (*gorm.DB, error) {
	config := map[string]string{
		"HTTP_USERNAME": "root",
		"HTTP_PASSWORD": "gconus5226",
		"HTTP_PORT":     "3306",
		"HTTP_HOST":     "127.0.0.1",
		"HTTP_NAME":     "RestfulProject",
	}
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config["HTTP_USERNAME"],
		config["HTTP_PASSWORD"],
		config["HTTP_HOST"],
		config["HTTP_PORT"],
		config["HTTP_NAME"])

	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{})
	if err != nil {
		return db, err
	}
	return db, err
}
