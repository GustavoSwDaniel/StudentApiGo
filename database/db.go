package database

import (
	"api-go-gin/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDatabase() {
	DatabaseString := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(DatabaseString))
	if err != nil {
		log.Panic("Error in database connect")
	}
	DB.AutoMigrate(&models.Student{})
}
