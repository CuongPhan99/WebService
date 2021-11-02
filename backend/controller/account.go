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
	account := []model.Accounts{}
	id := c.Param("id")

	if err := db.First(&account, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return account, nil
}

func GetAccountById(c echo.Context) error {
	account, _ := GetRepoAccountById(c)
	return c.JSON(http.StatusOK, account)
}

func UpdateCompany(c echo.Context) error {
	db := storage.GetDBInstance()
	account := &model.Accounts{}

	id := c.Param("id")
	logo := c.FormValue("logo")
	company_name := c.FormValue("company_name")
	address_company := c.FormValue("address_company")
	phone := c.FormValue("phone")
	name := c.FormValue("name")
	original_imprint := c.FormValue("original_imprint")
	security := c.FormValue("security")
	notification := c.FormValue("notification")

	db.First(&account, id)
	if logo != "" {
		account.Logo = logo
	}
	if company_name != "" {
		account.CompanyName = company_name
	}
	if address_company != "" {
		account.AddressCompany = address_company
	}
	if phone != "" {
		account.Phone = phone
	}
	if name != "" {
		account.Name = name
	}
	if original_imprint != "" {
		account.OriginalImprint = original_imprint
	}
	if security == "true" {
		account.Security = true
	} else if security == "false" {
		account.Security = false
	}
	if notification == "true" {
		account.Notification = true
	} else if notification == "false" {
		account.Notification = false
	}
	db.Save(&account)

	if account == nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}
	return c.JSON(http.StatusOK, account)
}
