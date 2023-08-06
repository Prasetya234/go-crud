package main

import (
	"fmt"
	"github.com/Prasetya234/go-crud.git/controller/homeController"
	"github.com/Prasetya234/go-crud.git/controller/studentController"
	"github.com/Prasetya234/go-crud.git/models"
	"github.com/gin-gonic/gin"
)

var PORT = "3000"

func main() {
	r := gin.Default()
	models.ConnectDb()

	r.GET("/", homeController.Index)
	r.GET("/api/student", studentController.FindAll)
	r.GET("/api/student/:id", studentController.FindById)
	r.POST("/api/student", studentController.Create)
	r.PUT("/api/student/:id", studentController.Update)
	r.DELETE("/api/student/:id", studentController.Delete)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"status": "404", "code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	err := r.Run(":" + PORT)
	if err != nil {

		panic("[Error] failed to start Gin server due to: " + err.Error())
		return

	}
	fmt.Println("Server running on port " + PORT)

}
