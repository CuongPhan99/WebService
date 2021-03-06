package storage

import (
	"log"

	config "github.com/CuongPhan99/WebService/backend/config"
	model "github.com/CuongPhan99/WebService/backend/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func NewDb(params ...string) *gorm.DB {
	var err error
	conString := config.GetPostgresConnectionString()
	DB, err = gorm.Open(config.GetDBType(), conString)

	if err != nil {
		log.Panic(err)
	} else {
		DB.AutoMigrate(&model.Accounts{}, &model.Customers{})
	}

	return DB
}

func GetDBInstance() *gorm.DB {
	return DB
}
