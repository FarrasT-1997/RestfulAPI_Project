package main

import (
	"project/team-alta/api-ecommerce/config"
	"project/team-alta/api-ecommerce/routes"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	config.Init_DB()

	routes.CategoryRoutes(e)
	routes.ChartRoutes(e)
	routes.TransactionDetail(e)
	routes.Transaction(e)
	routes.UserRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
