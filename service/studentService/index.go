package studentService

import (
	"errors"
	"fmt"
	"github.com/Prasetya234/go-crud.git/models"
	"github.com/Prasetya234/go-crud.git/util/responseHelper"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func FindAll(c *gin.Context) {
	var students []models.Student

	models.DB.Find(&students)
	fmt.Println(students)
	responseHelper.Ok(c, students)
}

func FindById(c *gin.Context) {
	var student models.Student
	id := c.Param("id")

	err := models.DB.Find(&student, id).Error
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			responseHelper.NotFound(c, "DATA NOT FOUND")
			return
		default:
			responseHelper.InternalServerError(c, err.Error())
			return
		}
	}
	responseHelper.Ok(c, student)
}

func Create(c *gin.Context) {
	var student models.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		responseHelper.BadRequest(c, err.Error())
		return
	}

	result := models.DB.Create(&student)

	if result.Error != nil {
		responseHelper.BadRequest(c, result.Error.Error())
	}
	responseHelper.Ok(c, student)
}

func Update(c *gin.Context) {
	var student models.Student
	id := c.Param("id")

	if err := c.ShouldBindJSON(&student); err != nil {
		responseHelper.BadRequest(c, err.Error())
		return
	}

	if models.DB.Model(&student).Where("id = ?", id).Updates(&student).RowsAffected == 0 {
		responseHelper.NotFound(c, "ID NOT FOUND")
		return
	}

	responseHelper.Ok(c, student)
}

func Delete(c *gin.Context) {
	var stundent models.Student
	id := c.Param("id")

	if models.DB.Delete(&stundent, id).RowsAffected == 0 {
		responseHelper.NotFound(c, "ID NOT FOUND")
		return
	}
	responseHelper.Ok(c, "SUCCESS")
}
