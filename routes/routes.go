package routes

import (
	"restfulAPI/constant"
	"restfulAPI/controller"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New(e *echo.Echo) {

	// e.POST("/login", controller.Login)
	// e.POST("/signup", controller.SignUpUser)

	e.GET("/category", controller.GetAllCategory)
	e.GET("/category/:idcategory", controller.GetAllSpecifiedProductCategory)
	e.GET("/product/:idproduct", controller.GetProductDetail)
	// e.PUT("/product/:idproduct", controllers.ChangeProductStock)

	// e.GET("/cart", controller.GetAllProductInChart)
	// e.POST("/cart/:nameproduct", controller.InsertProductIntoChart)
	// e.DELETE("/cart/:nameproduct", controller.DeleteProductInChart)
	// e.PUT("/cart/:nameproduct", controller.ChangeProductStockInCart)

	// e.GET("/transactiondetail", controller.GetTransactionStatus)
	// e.POST("/transactiondetail", controller.MakeTransactionDetail)
	// e.DELETE("/transactiondetail", controller.DeleteTransactionDetail)

	// e.GET("/transaction", controller.StatusTransaction)
	// e.PUT("/transaction/checkout", controller.ChangeTransactionStatus)
	// e.DELETE("/transaction", controller.DeleteTransaction)

	eJwt := e.Group("/jwt")
	eJwt.Use(middleware.JWT([]byte(constant.SECRET_JWT)))
	// eJwt.GET("/users", controller.GetAllUsers)
	// eJwt.PUT("/users/:userid", controller.ChangeProfile)
	// eJwt.GET("/users/:userid", controller.ShowProfile)
}
