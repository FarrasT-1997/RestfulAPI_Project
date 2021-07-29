package routes

func ChartRoutes(e *echo.Echo) {
	e.GET("/cart", GetAllProductInChart)
	e.POST("/cart/:nameproduct", InsertProductIntoChart)
	e.DELETE("/cart/:nameproduct", DeleteProductInChart)
	e.PUT("/cart/:nameproduct", ChangeProductStockInCart)
}
