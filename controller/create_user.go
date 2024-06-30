package controller

import (
	"sosservice/configurations/validation"
	"sosservice/controller/model/request"

	"github.com/gin-gonic/gin"
)

func CreateUser(context *gin.Context) {

	var userRequest request.UserRequest

	if err := context.ShouldBindJSON(&userRequest); err != nil {
		restErr := validation.ValidateUserError(err)
		context.JSON(restErr.Code, restErr)
		return
	}

	context.IndentedJSON(200, userRequest)
}
