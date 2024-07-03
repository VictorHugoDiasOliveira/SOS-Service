package controller

import (
	"net/http"
	"sosservice/src/configurations/logger"
	"sosservice/src/configurations/validation"
	"sosservice/src/model"
	"sosservice/src/view"

	"sosservice/src/controller/model/request"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func (uc *userControllerInterface) CreateUser(context *gin.Context) {

	logger.Info("Init CreatUser Controller",
		zap.String("journey", "createUser"),
	)

	var userRequest request.UserRequest

	if err := context.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to bind user info", err,
			zap.String("journey", "createUser"),
		)
		restErr := validation.ValidateUserError(err)
		context.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserDomain(userRequest.Email, userRequest.Password, userRequest.Name, userRequest.Age)

	domainResult, err := uc.service.CreateUserService(domain)
	if err != nil {
		context.JSON(err.Code, err)
		return
	}

	logger.Info("User created successfully",
		zap.String("journey", "createUser"),
	)

	context.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))
}
