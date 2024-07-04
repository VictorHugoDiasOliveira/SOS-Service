package controller

import (
	"net/http"
	"sosservice/src/configurations/validation"
	"sosservice/src/controller/model/request"
	"sosservice/src/model"
	"sosservice/src/view"

	"github.com/gin-gonic/gin"
)

func (uc *userControllerInterface) LoginUser(context *gin.Context) {
	var userRequest request.UserLogin

	if err := context.ShouldBindJSON(&userRequest); err != nil {
		restErr := validation.ValidateUserError(err)
		context.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserLoginDomain(userRequest.Email, userRequest.Password)

	domainResult, token, err := uc.service.LoginUserService(domain)
	if err != nil {
		context.JSON(err.Code, err)
		return
	}

	context.Header("Authorization", token)

	context.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))
}
