package userService

import (
	"github.com/Prasetya234/go-crud.git/models"
	"github.com/Prasetya234/go-crud.git/util/responseHelper"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const (
	SECRET = "asdadascaxczxcssdadf"
)

func Register(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		responseHelper.BadRequest(c, err.Error())
		return
	}

	passwordEncode, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	userCreate := models.User{Email: user.Email, Password: string(passwordEncode), FullName: user.FullName, PhoneNumber: user.PhoneNumber}

	result := models.DB.Create(&userCreate)

	if result.Error != nil {
		responseHelper.InternalServerError(c, "EMAIL ALREADY EXIST")
		return
	}

	responseHelper.Ok(c, userCreate)
}

func Login(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}
	if c.Bind(&body) != nil {
		responseHelper.BadRequest(c, "Bad Reqequest")
		return
	}

	var user models.User
	models.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		responseHelper.InternalServerError(c, "EMAIL NOT FOUND")
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		responseHelper.InternalServerError(c, "BAD CREDENTIALS")
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 6 * 1).Unix(),
	})
	tokenResp, err3 := token.SignedString([]byte(SECRET))
	if err3 != nil {
		responseHelper.InternalServerError(c, "INVALID CREATE TOKEN")
		panic(err3)
		return
	}
	responseHelper.Ok(c, gin.H{
		"token": tokenResp,
	})
}
