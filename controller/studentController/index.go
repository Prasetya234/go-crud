package studentController

import (
	"github.com/Prasetya234/go-crud.git/service/studentService"
	"github.com/gin-gonic/gin"
)

func FindAll(c *gin.Context) {
	studentService.FindAll(c)
}

func FindById(c *gin.Context) {
	studentService.FindById(c)
}

func Create(c *gin.Context) {
	studentService.Create(c)
}

func Update(c *gin.Context) {
	studentService.Update(c)
}

func Delete(c *gin.Context) {
	studentService.Delete(c)
}
