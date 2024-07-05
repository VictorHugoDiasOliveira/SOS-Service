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
	logger.Info("Starting User Creation", zap.String("journey", "CreateUser"))
	var userRequest request.UserRequest

	if err := context.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Failed Trying to Bind JSON", err, zap.String("journey", "CreateUser"))
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

	context.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))
	logger.Info("User Created Successfully", zap.String("journey", "CreateUser"))
}
