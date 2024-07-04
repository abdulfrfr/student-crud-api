package database

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
    "github.com/your-username/student-crud-api/models"
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
