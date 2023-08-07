package main

import (
	"fmt"
	"github.com/Prasetya234/go-crud.git/controller/homeController"
	"github.com/Prasetya234/go-crud.git/controller/studentController"
	"github.com/Prasetya234/go-crud.git/controller/userController"
	"github.com/Prasetya234/go-crud.git/middleware"
	"github.com/Prasetya234/go-crud.git/models"
	"github.com/gin-gonic/gin"
)

var PORT = "3000"

func main() {
	r := gin.Default()
	models.ConnectDb()

	// authentication
	r.POST("/api/register", userController.Register)
	r.POST("/api/login", userController.Login)

	//student
	r.GET("/", middleware.RequireAuth, homeController.Index)
	r.GET("/api/student", middleware.RequireAuth, studentController.FindAll)
	r.GET("/api/student/:id", middleware.RequireAuth, studentController.FindById)
	r.POST("/api/student", middleware.RequireAuth, studentController.Create)
	r.PUT("/api/student/:id", middleware.RequireAuth, studentController.Update)
	r.DELETE("/api/student/:id", middleware.RequireAuth, studentController.Delete)

	//	path don't exist
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"status": "404", "massage": "PAGE_NOT_FOUND", "error": "Page not found"})
	})

	err := r.Run(":" + PORT)
	if err != nil {

		panic("[Error] failed to start Gin server due to: " + err.Error())
		return

	}
	fmt.Println("Server running on port " + PORT)

}
