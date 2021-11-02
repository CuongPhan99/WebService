package controller

import (
	"net/http"

	model "github.com/CuongPhan99/WebService/backend/models"
	"github.com/CuongPhan99/WebService/backend/storage"
	"github.com/labstack/echo/v4"
)

//Handler
func GetRepoCustomers() ([]model.Customers, error) {
	db := storage.GetDBInstance()
	customers := []model.Customers{}

	if err := db.Find(&customers).Error; err != nil {
		return nil, err
	}
	return customers, nil
}

func GetCustomers(c echo.Context) error {
	customers, _ := GetRepoCustomers()
	return c.JSON(http.StatusOK, customers)
}

func AddCustomer(c echo.Context) error {
	db := storage.GetDBInstance()

	last_name := c.FormValue("last_name")
	first_name := c.FormValue("first_name")
	email := c.FormValue("email")
	company_name := c.FormValue("company_name")
	department := c.FormValue("department")

	customer := &model.Customers{FirstName: first_name, LastName: last_name, Email: email, CompanyName: company_name, Department: department}
	db.Create(&customer)
	if customer == nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}
	return c.JSON(http.StatusOK, &customer)
}
