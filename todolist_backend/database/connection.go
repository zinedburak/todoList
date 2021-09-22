package database

import (
	"todolist_backend/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=ec2-34-255-134-200.eu-west-1.compute.amazonaws.com user=dusvljbjslskam password=9df38dec15602670e0fc65086493e5d49f7fb80fe8dfc255ddafe317fc5f393e dbname=d1a78jqqnnk9t8 port=5432"
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database ") // panic will stop the aplication
	}
	DB = connection

	connection.AutoMigrate(&models.Todos{})
}
