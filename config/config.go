package config

import (
	"github.com/fawzy17/rest-api-gin-gorm/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/rest-api"))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&models.Product{})

	DB = database
}
