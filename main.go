package main

import (
	"os"

	"github.com/abdulfrfr/student-crud-api/controllers"
	"github.com/abdulfrfr/student-crud-api/database"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	// Connect to database
	database.ConnectDatabase()

	// Routes
	v1 := router.Group("/api/v1")
	{
		v1.POST("/students", controllers.CreateStudent)
		v1.GET("/students", controllers.GetStudents)
		v1.GET("/students/:id", controllers.GetStudentByID)
		v1.PUT("/students/:id", controllers.UpdateStudent)
		v1.DELETE("/students/:id", controllers.DeleteStudent)
	}

	router.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "UP"})
	})

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}
	// Run the server
	router.Run(":" + port)
}
