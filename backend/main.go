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

	e.GET("/account/:id", controller.GetAccountById)

	e.POST("/account/:id", controller.UpdateCompany)

	e.GET("/customers", controller.GetCustomers)

	e.POST("/customer/add", controller.AddCustomer)

	e.GET("/customer/:id", controller.GetCustomerById)

	e.POST("/customer/:id", controller.UpdateCustomer)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
