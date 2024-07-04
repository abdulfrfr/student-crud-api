package database

import (
	"github.com/abdulfrfr/student-crud-api/models"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("sqlite3", "students.db")
	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&models.Student{})

	DB = database
}
