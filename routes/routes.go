package routes

import (
	"restfulAPI/constant"
	"restfulAPI/controller"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New(e *echo.Echo) {

	e.POST("/login", controller.Login)
	e.POST("/signup", controller.SignUpUser)

	e.GET("/category", controller.GetAllCategory)
	e.GET("/category/:idcategory", controller.GetAllSpecifiedProductCategory)
	e.GET("/product/:idproduct", controller.GetProductDetail)
	e.GET("/paymentmethod", controller.GetPaymentMethod)

	eJwt := e.Group("/jwt")
	eJwt.Use(middleware.JWT([]byte(constant.SECRET_JWT)))
	eJwt.PUT("/users/:userid", controller.ChangeProfile)
	eJwt.GET("/users/:userid", controller.ShowProfile)
	eJwt.PUT("/logout/:userid", controller.Logout)

	eJwt.POST("/transaction", controller.MakeTransaction)
	eJwt.GET("/transaction", controller.GetAllTransaction)
	eJwt.DELETE("/transaction/:transactionId", controller.DeleteTransaction)
	eJwt.PUT("/transaction/payment/:transactionId/:paymentId", controller.ChangePaymentMethod)
	eJwt.PUT("/transaction/checkout/:transactionId", controller.TransactionCheckout)
	eJwt.PUT("/transaction/paid/:transactionId", controller.ChangeStatus)

	eJwt.POST("/transaction/:transactionId/cart/:productId", controller.MakeCartID)
	eJwt.PUT("/transaction/:transactionId/cart/:cartId/:quantity", controller.ChangeQuantity)
	eJwt.DELETE("/transaction/:transactionId/:cartId", controller.DeleteCart)
	eJwt.GET("/transaction/cart/:transactionId", controller.GetCartOfTransaction)
}
