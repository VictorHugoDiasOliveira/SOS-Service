package main

import (
	"authservice/config"
	"authservice/controllers"
	"authservice/middlewares"
	"authservice/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := config.SetupRouter()

	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.User{})

	authorized := r.Group("/")
	authorized.Use(middlewares.AuthMiddleware())
	{
		authorized.GET("/protected", func(c *gin.Context) {
			userID := c.MustGet("userID").(uint)
			c.JSON(200, gin.H{
				"message": "Hello User",
				"userID":  userID,
			})
		})
	}

	// r = controllers.GetUsers(r)
	// r = controllers.GetUserById(r)
	r = controllers.Register(r)
	r = controllers.Login(r)

	r.Run(":8080")
}
