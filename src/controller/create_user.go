package controller

import (
	"net/http"
	"sosservice/src/configurations/logger"
	"sosservice/src/configurations/validation"
	"sosservice/src/controller/request"
	"sosservice/src/controller/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateUser(context *gin.Context) {

	logger.Info("Init CreatUser Controller",
		zap.String("key", "journey"),
		zap.String("string", "createUser"),
	)

	var userRequest request.UserRequest

	if err := context.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to bind user info", err,
			zap.String("key", "journey"),
			zap.String("string", "createUser"),
		)
		restErr := validation.ValidateUserError(err)
		context.JSON(restErr.Code, restErr)
		return
	}

	// Add to database

	logger.Info("User created successfully",
		zap.String("key", "journey"),
		zap.String("string", "createUser"),
	)

	response := response.UserResponse{
		ID:    1,
		Email: userRequest.Email,
	}

	context.JSON(http.StatusOK, response)
}
