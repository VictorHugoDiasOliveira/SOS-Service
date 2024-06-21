package handlers

import (
	"authservice/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) *gin.Engine {
	r.POST("/register", func(c *gin.Context) {
		var user models.User

		if err := c.Bind(&user); err != nil {
			return
		}

		users = append(users, user)
		c.IndentedJSON(http.StatusCreated, user)
	})

	return r
}
