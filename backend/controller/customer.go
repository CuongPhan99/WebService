package controller

import (
	"net/http"

	model "github.com/CuongPhan99/WebService/backend/models"
	"github.com/CuongPhan99/WebService/backend/storage"
	"github.com/labstack/echo/v4"
)

//Handler
func GetCustomers(c echo.Context) error {
	db := storage.GetDBInstance()
	customers := []model.Customers{}

	hide_all := c.QueryParam("hide_all")
	if hide_all == "false" {
		db.Order("id desc").Find(&customers)
	} else if hide_all == "true" {
		db.Order("id desc").Where("active = ?", true).Find(&customers)
	}
	if customers == nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}
	return c.JSON(http.StatusOK, &customers)
}

func AddCustomer(c echo.Context) error {
	db := storage.GetDBInstance()

	last_name := c.FormValue("last_name")
	first_name := c.FormValue("first_name")
	email := c.FormValue("email")
	company_name := c.FormValue("company_name")
	department := c.FormValue("department")

	customer := &model.Customers{FirstName: first_name, LastName: last_name, Email: email, CompanyName: company_name, Department: department, Active: true}
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

	db.First(&customer, id)
	if !customer.Active {
		customer.Active = true
	} else if customer.Active {
		customer.Active = false
	}
	db.Save(&customer)
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
