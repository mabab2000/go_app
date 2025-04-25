package main

import (
	"student-api/database"
	"student-api/handlers"
	"student-api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database connection
	database.InitDB()

	// Automatically migrate the Student model
	database.DB.AutoMigrate(&models.Student{})

	// Set up the Gin router
	r := gin.Default()

	// Define routes for CRUD operations
	r.POST("/students", handlers.CreateStudent)
	r.GET("/students", handlers.GetAllStudents)
	r.GET("/students/:id", handlers.GetStudentByID)
	r.PUT("/students/:id", handlers.UpdateStudent)
	r.DELETE("/students/:id", handlers.DeleteStudent)

	// Start the server on port 3000
	r.Run(":3000")
}
