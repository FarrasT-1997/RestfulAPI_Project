package main

import (
	"project/team-alta/api-ecommerce/config"
	"project/team-alta/api-ecommerce/routes"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	config.Init_DB()

	routes.New(e)

	e.Logger.Fatal(e.Start(":8080"))
}
