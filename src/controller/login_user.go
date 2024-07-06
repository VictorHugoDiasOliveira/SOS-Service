package controller

import (
	"net/http"
	"sosservice/src/configurations/validation"
	"sosservice/src/controller/model/request"
	"sosservice/src/model"
	"sosservice/src/view"

	"github.com/gin-gonic/gin"
)

// LoginUser allows a user to log and obtain an authentication token
// @Summary User Login
// @Description Allows a user to log in and receive and authentication token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param userLogin body request.UserLogin true "User login credentials"
// @Success 200 {object} response.UserResponse "Login successfully, authentication token provided"
// @Header 200 {string} Authorization "Authentication token"
// @Failure 403 {object} rest_err.RestErr "Error: invalid login credentials"
// @Router /login [post]
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
