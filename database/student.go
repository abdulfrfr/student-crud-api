package database

import (
	"fmt"
	"os"

	"github.com/abdulfrfr/student-crud-api/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func ConnectDatabase() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbName, dbPassword)

	database, err := gorm.Open("postgres", dsn)
	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&models.Student{})

	DB = database
}
