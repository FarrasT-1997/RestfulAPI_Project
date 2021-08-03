#!/bin/bash

cd RestfulAPI_Project
go get github.com/labstack/echo
go get github.com/labstack/echo/middleware
go get github.com/dgrijalva/jwt-go
go get gorm.io/driver/mysql
go get gorm.io/gorm

go run main.go