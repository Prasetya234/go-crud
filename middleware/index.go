package middleware

import (
	"fmt"
	"github.com/Prasetya234/go-crud.git/models"
	"github.com/Prasetya234/go-crud.git/service/userService"
	"github.com/Prasetya234/go-crud.git/util/responseHelper"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func RequireAuth(c *gin.Context) {
	tokenStr := c.GetHeader("Authorization")
	if tokenStr == "" {
		responseHelper.Unauthorized(c)
		return
	}
	token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unxpected sigening method: %v: ", token.Header["alg"])
		}
		return []byte(userService.SECRET), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			responseHelper.Unauthorized(c)
			return
		}
		var user models.User
		models.DB.First(&user, claims["sub"])
		if user.ID == 0 {
			responseHelper.Unauthorized(c)
			return
		}
		c.Next()
	} else {
		responseHelper.Unauthorized(c)
		return
	}
}
