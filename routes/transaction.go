package routes

func Transaction(e *echo.Echo) {
	e.GET("/transaction", StatusTransaction)
	e.PUT("/transaction/checkout", ChangeTransactionStatus)
	e.DELETE("/transaction", DeleteTransaction)
}
