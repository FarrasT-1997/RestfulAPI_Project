package routes

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New(e *echo.Echo) {

	e.POST("/login", controllers.Login)
	e.POST("/signup", controllers.SignUpUser)

	e.GET("/category", controllers.GetAllCategory)
	e.GET("/category/:namecategory", controllers.GetAllSpecifiedProductCategory)
	e.GET("/product/:nameproduct", controllers.GetProductDetail)
	// e.PUT("/product/:nameproduct", controllers.ChangeProductStock)

	e.GET("/cart", controllers.GetAllProductInChart)
	e.POST("/cart/:nameproduct", controllers.InsertProductIntoChart)
	e.DELETE("/cart/:nameproduct", controllers.DeleteProductInChart)
	e.PUT("/cart/:nameproduct", controllers.ChangeProductStockInCart)

	e.GET("/transactiondetail", controllers.GetTransactionStatus)
	e.POST("/transactiondetail", controllers.MakeTransactionDetail)
	e.DELETE("/transactiondetail", controllers.DeleteTransactionDetail)

	e.GET("/transaction", controllers.StatusTransaction)
	e.PUT("/transaction/checkout", controllers.ChangeTransactionStatus)
	e.DELETE("/transaction", controllers.DeleteTransaction)

	eJwt := e.Group("/jwt")
	eJwt.Use(middleware.JWT([]byte(constant.SECRET_JWT)))
	eJwt.GET("/users", controllers.GetAllUsers)
	eJwt.PUT("/users/:userid", controllers.ChangeProfile)
}
