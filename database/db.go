package database

import (
	"log"

	"github.com/Polengolas/gin-api-rest/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectToDatabase() {
	stringConection := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringConection))
	if err != nil {
		log.Panic("Error to conect database")
	}
	DB.AutoMigrate(&models.Student{})
}
