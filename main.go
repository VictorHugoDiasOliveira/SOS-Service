package main

import (
	"authservice/models"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var users = []models.Users{
	{ID: 1, Email: "victorhugo@gmail.com", Password: "12345"},
	{ID: 2, Email: "larissinha@gmail.com", Password: "qwert"},
	{ID: 3, Email: "mamadinha@sempudor.com", Password: "brasileirinhas"},
}

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost"},
		AllowMethods:     []string{"POST", "GET"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Lenght"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.POST("/register", Register)
	router.POST("/login", Login)

	router.GET("/users", GetUsers)

	router.Run("localhost:8080")
}

func Register(c *gin.Context) {
	var userJson models.Users

	if err := c.BindJSON(&userJson); err != nil {
		return
	}

	users = append(users, userJson)
	c.IndentedJSON(http.StatusCreated, userJson)
}

func Login(c *gin.Context) {
	var userJson models.Users

	if err := c.BindJSON(&userJson); err != nil {
		return
	}

	for _, user := range users {
		if (user.Email == userJson.Email) && (user.Password == userJson.Password) {
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"userId": userJson.ID,
				"exp":    time.Now().Add(time.Hour * 72).Unix(),
			})

			tokenString, _ := token.SignedString([]byte("secret"))
			c.JSON(200, gin.H{
				"token": tokenString,
			})
			return
		}
	}

	c.JSON(200, gin.H{
		"error": "User not found",
	})

	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	// 	"userId": userJson.ID,
	// 	"exp":    time.Now().Add(time.Hour * 72).Unix(),
	// })

	// tokenString, _ := token.SignedString([]byte("secret"))

	// c.JSON(200, gin.H{
	// 	"token": tokenString,
	// })
}

func GetUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}
