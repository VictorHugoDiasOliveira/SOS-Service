package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUsers(r *gin.Engine) *gin.Engine {
	r.GET("/users", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, users)
	})

	return r
}

func GetUserById(r *gin.Engine) *gin.Engine {
	r.GET("/users/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))

		for _, user := range users {
			if user.ID == id {
				c.JSON(http.StatusOK, user)
				return
			}
		}
		c.JSON(200, gin.H{
			"user": "not found",
		})
	})

	return r
}
