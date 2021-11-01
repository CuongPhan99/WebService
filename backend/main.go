package main

import (
	"github.com/CuongPhan99/WebService/backend/controller"
	"github.com/CuongPhan99/WebService/backend/storage"
	"github.com/labstack/echo/v4"
)

func main() {
	// Echo instance
	e := echo.New()
	storage.NewDb()

	// Routes
	e.GET("/accounts", controller.GetAccounts)

	e.GET("/accounts/:id", controller.GetAccountById)

	e.POST("/accounts/:id", controller.UpdateCompany)

	e.GET("/customers", controller.GetCustomers)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
