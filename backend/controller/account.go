package controller

import (
	"net/http"

	model "github.com/CuongPhan99/WebService/backend/models"
	"github.com/CuongPhan99/WebService/backend/storage"
	"github.com/labstack/echo/v4"
)

// Handler
func GetRepoAccounts() ([]model.Accounts, error) {
	db := storage.GetDBInstance()
	accounts := []model.Accounts{}

	if err := db.Find(&accounts).Error; err != nil {
		return nil, err
	}
	return accounts, nil
}

func GetAccounts(c echo.Context) error {
	accounts, _ := GetRepoAccounts()
	return c.JSON(http.StatusOK, accounts)
}

func GetRepoAccountById(c echo.Context) ([]model.Accounts, error) {
	db := storage.GetDBInstance()
	accounts := []model.Accounts{}
	id := c.Param("id")

	if err := db.First(&accounts, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return accounts, nil
}

func GetAccountById(c echo.Context) error {
	accounts, _ := GetRepoAccountById(c)
	return c.JSON(http.StatusOK, accounts)
}

func UpdateCompany(c echo.Context) error {
	db := storage.GetDBInstance()
	accounts := &model.Accounts{}

	id := c.Param("id")
	logo := c.FormValue("logo")
	company_name := c.FormValue("company_name")
	address_company := c.FormValue("address_company")
	phone := c.FormValue("phone")
	name := c.FormValue("name")
	originer_imprint := c.FormValue("originer_imprint")
	security := c.FormValue("security")
	notification := c.FormValue("notification")

	db.First(&accounts, id)
	if logo != "" {
		accounts.Logo = logo
	}
	if company_name != "" {
		accounts.CompanyName = company_name
	}
	if address_company != "" {
		accounts.AddressCompany = address_company
	}
	if phone != "" {
		accounts.Phone = phone
	}
	if name != "" {
		accounts.Name = name
	}
	if originer_imprint != "" {
		accounts.OriginerImprint = originer_imprint
	}
	if security == "true" {
		accounts.Security = true
	} else if security == "false" {
		accounts.Security = false
	}
	if notification == "true" {
		accounts.Notification = true
	} else if notification == "false" {
		accounts.Notification = false
	}
	db.Save(&accounts)

	if accounts == nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}
	return c.JSON(http.StatusOK, accounts)
}
