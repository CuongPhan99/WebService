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

	//Account
	e.GET("/accounts", controller.GetAccounts)

	e.GET("/account/:id", controller.GetAccountById)

	e.POST("/account/:id", controller.UpdateCompany)

	//Customer
	e.GET("/customers", controller.GetCustomers)

	e.POST("/customer/add", controller.AddCustomer)

	e.GET("/customer/:id", controller.GetCustomerById)

	e.PUT("/customer/:id", controller.UpdateCustomer)

	e.POST("/customer/hide/:id", controller.HideCustomer)

	e.DELETE("/customer/:id", controller.DeleteCustomer)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
