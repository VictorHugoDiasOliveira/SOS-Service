package controller

import (
	"net/http"
	"sosservice/configurations/logger"
	"sosservice/configurations/validation"
	"sosservice/controller/model/request"
	"sosservice/controller/model/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap/zapcore"
)

func CreateUser(context *gin.Context) {

	logger.Info("Init CreatUser Controller", zapcore.Field{
		Key:    "journey",
		String: "createUser",
	})

	var userRequest request.UserRequest

	if err := context.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to bind user info", err, zapcore.Field{
			Key:    "journey",
			String: "createUser",
		})
		restErr := validation.ValidateUserError(err)
		context.JSON(restErr.Code, restErr)
		return
	}

	// Add to database

	logger.Info("User created successfully", zapcore.Field{
		Key:    "journey",
		String: "createUser",
	})

	response := response.UserResponse{
		ID:    1,
		Email: userRequest.Email,
	}

	context.JSON(http.StatusOK, response)
}
