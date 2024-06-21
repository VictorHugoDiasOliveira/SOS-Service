package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(r *gin.Engine) *gin.Engine {
	r.GET("/users", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, users)
	})

	return r
}
