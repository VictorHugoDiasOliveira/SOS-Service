package controller

import (
	"net/http"
	"sosservice/src/configurations/validation"
	"sosservice/src/model"
	"sosservice/src/view"

	"sosservice/src/controller/model/request"

	"github.com/gin-gonic/gin"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func (uc *userControllerInterface) CreateUser(context *gin.Context) {
	var userRequest request.UserRequest

	if err := context.ShouldBindJSON(&userRequest); err != nil {
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
}
