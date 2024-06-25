package controllers

// do not push it to production

import (
	"authservice/config"
	"authservice/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserResponse struct {
	Id       uint   `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GetUsers(r *gin.Engine) *gin.Engine {
	r.GET("/users", func(c *gin.Context) {
		var users []UserResponse

		result := config.DB.Model(models.User{}).Find(&users)

		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Message": "Query failed",
				"Result":  result.Error,
			})
			return
		}
		if result.RowsAffected == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"Message": "No records found",
			})
			return
		}

		c.IndentedJSON(http.StatusOK, users)
	})

	return r
}

func GetUsersById(r *gin.Engine) *gin.Engine {
	r.GET("/users/:id", func(c *gin.Context) {
		var user models.User

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result := config.DB.Where("id = ?", user.ID).First(&user)

		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Message": "Query failed",
				"Result":  result.Error,
			})
			return
		}
		if result.RowsAffected == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"Message": "No records found",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"ID":       user.ID,
			"Email":    user.Email,
			"Password": user.Password,
		})
	})

	return r
}
