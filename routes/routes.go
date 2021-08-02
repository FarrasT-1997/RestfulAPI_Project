package routes

import (
	"restfulAPI/constant"
	"restfulAPI/controller"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New(e *echo.Echo) {

	e.POST("/login", controller.Login)       //ok
	e.POST("/signup", controller.SignUpUser) //ok

	e.GET("/category", controller.GetAllCategory)                             //ok
	e.GET("/category/:idcategory", controller.GetAllSpecifiedProductCategory) //ok
	e.GET("/product/:idproduct", controller.GetProductDetail)                 //ok
	e.GET("/paymentmethod", controller.GetPaymentMethod)                      //ok

	eJwt := e.Group("/jwt")
	eJwt.Use(middleware.JWT([]byte(constant.SECRET_JWT)))
	eJwt.PUT("/users/:userid", controller.ChangeProfile) //ok
	eJwt.GET("/users/:userid", controller.ShowProfile)   //ok
	eJwt.PUT("/logout/:userid", controller.Logout)       //ok

	eJwt.POST("/transaction", controller.MakeTransaction)                                      //ok
	eJwt.GET("/transaction", controller.GetAllTransaction)                                     //ok
	eJwt.DELETE("/transaction/:transactionId", controller.DeleteTransaction)                   // pending
	eJwt.PUT("/transaction/payment/:transactionId/:paymentId", controller.ChangePaymentMethod) // ok
	eJwt.PUT("/transaction/paid/:transactionId", controller.ChangeStatus)                      // ok

	eJwt.POST("/transaction/:transactionId/cart/:productId", controller.MakeCartID)           //ok
	eJwt.PUT("/transaction/:transactionId/cart/:cartId/:quantity", controller.ChangeQuantity) //ok
	eJwt.DELETE("/transaction/:transactionId/:cartId", controller.DeleteCart)                 //ok
	eJwt.GET("/transaction/cart/:transactionId", controller.GetCartOfTransaction)             //ok
}
