package handlers

import (
	"authservice/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var users = []models.User{
	{ID: 1, Email: "victorhugo@gmail.com", Password: "12345"},
	{ID: 2, Email: "hugo@gmail.com", Password: "qwert"},
	{ID: 3, Email: "vitinho@sempudor.com", Password: "brasileirinhas"},
}

func Login(r *gin.Engine) *gin.Engine {
	var userJson models.User

	r.POST("/login", func(c *gin.Context) {
		c.Bind(&userJson)

		for _, user := range users {
			if (user.Email == userJson.Email) && (user.Password == userJson.Password) {
				token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
					"userId": userJson.ID,
					"exp":    time.Now().Add(time.Hour * 72).Unix(),
				})

				tokenString, _ := token.SignedString([]byte("anysecret"))
				c.JSON(200, gin.H{
					"token": tokenString,
				})
				return
			}
		}

		c.JSON(200, gin.H{
			"error": "User not found",
		})
	})

	return r
}
