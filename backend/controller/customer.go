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

	if err := db.Order("id desc").Where("active_all = ?", true).Find(&customers).Error; err != nil {
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

	customer := &model.Customers{FirstName: first_name, LastName: last_name, Email: email, CompanyName: company_name, Department: department, Active: true, ActiveAll: true}
	db.Create(&customer)
	if customer == nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}
	return c.JSON(http.StatusOK, &customer)
}

func GetRepoCustomerById(c echo.Context) ([]model.Customers, error) {
	db := storage.GetDBInstance()
	customer := []model.Customers{}
	id := c.Param("id")

	if err := db.First(&customer, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return customer, nil
}

func GetCustomerById(c echo.Context) error {
	customer, _ := GetRepoCustomerById(c)
	return c.JSON(http.StatusOK, customer)
}

func UpdateCustomer(c echo.Context) error {
	db := storage.GetDBInstance()
	customer := &model.Customers{}

	id := c.Param("id")
	first_name := c.FormValue("first_name")
	last_name := c.FormValue("last_name")
	email := c.FormValue("email")
	company_name := c.FormValue("company_name")
	department := c.FormValue("department")

	db.Find(&customer, id)
	customer.FirstName = first_name
	customer.LastName = last_name
	customer.Email = email
	customer.CompanyName = company_name
	customer.Department = department
	db.Save(&customer)

	if customer == nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}
	return c.JSON(http.StatusOK, customer)
}

func HideCustomer(c echo.Context) error {
	db := storage.GetDBInstance()
	customer := &model.Customers{}
	id := c.Param("id")

	db.Model(&customer).Where("id = ?", id).Update("active", false)
	if customer == nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}
	return c.JSON(http.StatusOK, customer)
}

func UnhideCustomer(c echo.Context) error {
	db := storage.GetDBInstance()
	customer := &model.Customers{}

	db.Model(&customer).Update("active_all", true)
	if customer == nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}
	return c.JSON(http.StatusOK, customer)
}

func DeleteCustomer(c echo.Context) error {
	db := storage.GetDBInstance()
	customer := &model.Customers{}
	id := c.Param("id")
	db.Delete(&customer, id)
	if customer == nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}
	return c.JSON(http.StatusOK, customer)
}
