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
