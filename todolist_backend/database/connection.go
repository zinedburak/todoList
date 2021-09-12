package database

import (
	"todolist_backend/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=db user=postgres password=postgres dbname=postgres port=5432"
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database ") // panic will stop the aplication
	}
	DB = connection

	connection.AutoMigrate(&models.Todos{})
}
