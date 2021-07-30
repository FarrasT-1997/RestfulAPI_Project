package main

import (
	"restfulAPI/config"
	"restfulAPI/routes"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	config.Init_DB()

	routes.New(e)

	e.Logger.Fatal(e.Start(":8080"))
}
