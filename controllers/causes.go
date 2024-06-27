package controllers

import (
	"authservice/config"
	"authservice/middlewares"
	"authservice/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CauseInput struct {
	Name string `json:"name" binding:"required"`
}

func RegisterCause(rg *gin.RouterGroup) {
	rg.POST("/causes", middlewares.AuthorizationMiddleware(), func(c *gin.Context) {

		var input CauseInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		cause := models.Cause{Name: input.Name}
		if err := config.DB.Create(&cause).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Cause already exists"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
	})
}

func GetCauses(rg *gin.RouterGroup) {
	rg.GET("/causes", func(c *gin.Context) {

		var causes []CauseInput
		result := config.DB.Model(models.Cause{}).Find(&causes)

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

		c.IndentedJSON(http.StatusOK, causes)
	})
}
