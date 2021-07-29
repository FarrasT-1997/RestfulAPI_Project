package routes

func TransactionDetail(e *echo.Echo) {
	e.GET("/transactiondetail", GetTransactionStatus)
	e.POST("/transactiondetail", MakeTransactionDetail)
	e.DELETE("/transactiondetail", DeleteTransactionDetail)
}
