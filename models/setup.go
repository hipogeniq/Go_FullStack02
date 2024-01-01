package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := fmt.Sprintf(
		"host=localhost user=%s password=%s dbname=%s port=5433 sslmode=disable",
		"postgres",
		"1234",
		"GoTest",
	)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	database.AutoMigrate(&Quest{})

	DB = database
}
