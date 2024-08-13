package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Global variable to hold the database connection
var DB *gorm.DB

// Connect initializes the database connection using GORM
func Connect() {
	var err error
	// Data Source Name (DSN) for MySQL database connection
	dsn := "admin:soeel123@tcp(localhost:3306)/soeeldb?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
}

// GetDB returns the database connection instance
func GetDB() *gorm.DB {
	return DB
}
