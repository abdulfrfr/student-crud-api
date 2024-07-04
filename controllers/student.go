package controllers

import (
    "github.com/gin-gonic/gin"
    "github.com/your-username/student-crud-api/database"
    "github.com/your-username/student-crud-api/models"
    "net/http"
)

func CreateStudent(c *gin.Context) {
    var input models.Student
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    database.DB.Create(&input)
    c.JSON(http.StatusOK, input)
}

func GetStudents(c *gin.Context) {
    var students []models.Student
    database.DB.Find(&students)
    c.JSON(http.StatusOK, students)
}

func GetStudentByID(c *gin.Context) {
    var student models.Student
    if err := database.DB.Where("id = ?", c.Param("id")).First(&student).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    c.JSON(http.StatusOK, student)
}

func UpdateStudent(c *gin.Context) {
    var student models.Student
    if err := database.DB.Where("id = ?", c.Param("id")).First(&student).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    var input models.Student
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    database.DB.Model(&student).Updates(input)
    c.JSON(http.StatusOK, student)
}

func DeleteStudent(c *gin.Context) {
    var student models.Student
    if err := database.DB.Where("id = ?", c.Param("id")).First(&student).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    database.DB.Delete(&student)
    c.JSON(http.StatusOK, gin.H{"data": true})
}
