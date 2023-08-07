package userController

import (
	"github.com/Prasetya234/go-crud.git/service/userService"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	userService.Register(c)
}

func Login(c *gin.Context) {
	userService.Login(c)
}
