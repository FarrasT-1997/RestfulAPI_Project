package routes

func UserRoutes(e *echo.Echo) {
	e.GET("/jwt/users", GetAllUsers)
	e.POST("/login", Login)
	e.POST("/signup", SignUpUser)
	e.PUT("/jwt/users/:userid", ChangeProfile)
}
