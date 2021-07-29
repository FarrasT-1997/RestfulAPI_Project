package routes

func CategoryRoutes(e *echo.Echo) {
	e.GET("/category", GetAllCategory)
	e.GET("/category/:namecategory", GetAllSpecifiedProductCategory)
	e.GET("/product/:nameproduct", GetProductDetail)
	// e.PUT("/product/:nameproduct", ChangeProductStock)
}
