package controller

import (
	"net/http"
	"net/mail"
	"sosservice/src/configurations/rest_err"
	"sosservice/src/view"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// FindUserById retrieves a user information from id
// @Summary Find User by ID
// @Description Retrieves a user details based on user ID provided as parameter
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "ID of the user to be retrieved"
// @Param Authorization header string true "Insert your access token" default(Bearer <token>)
// @Success 200 {object} response.UserResponse "User information retrieved successfully"
// @Failure 400 {object} rest_err.RestErr "Error: Invalid user id"
// @Failure 404 {object} rest_err.RestErr "User not found"
// @Router /getUserById/{id} [get]
func (uc *userControllerInterface) FindUserById(context *gin.Context) {
	userId := context.Param("id")

	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errorMessage := rest_err.NewBadRequestError("ID is not a valid ID")
		context.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByIdService(userId)
	if err != nil {
		context.JSON(err.Code, err)
		return
	}
	context.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}

// FindUserByEmail retrieves a user information from email
// @Summary Find User by Email
// @Description Retrieves a user details based on user Email provided as parameter
// @Tags Users
// @Accept json
// @Produce json
// @Param email path string true "Email of the user to be retrieved"
// @Param Authorization header string true "Insert your access token" default(Bearer <token>)
// @Success 200 {object} response.UserResponse "User information retrieved successfully"
// @Failure 400 {object} rest_err.RestErr "Error: Invalid user email"
// @Failure 404 {object} rest_err.RestErr "User not found"
// @Router /getUserByEmail/{email} [get]
func (uc userControllerInterface) FindUserByEmail(context *gin.Context) {
	userEmail := context.Param("email")

	if _, err := mail.ParseAddress(userEmail); err != nil {
		errorMessage := rest_err.NewBadRequestError("email is not a valid email")

		context.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByEmailService(userEmail)
	if err != nil {
		context.JSON(err.Code, err)
		return
	}
	context.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}
