package main

import (
	"net/http"

	"github.com/CuongPhan99/WebService/backend/controller"
	"github.com/CuongPhan99/WebService/backend/storage"
	"github.com/labstack/echo/v4"
)

func main() {
	// Echo instance
	e := echo.New()
	storage.NewDb()

	// Routes
	e.GET("/", hello)

	e.GET("/accounts", controller.GetAccounts)

	e.GET("/accounts/:id", controller.GetAccountById)

	e.POST("/accounts/:id", controller.UpdateCompany)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
