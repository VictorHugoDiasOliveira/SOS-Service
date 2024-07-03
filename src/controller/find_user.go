package controller

import (
	"net/http"
	"net/mail"
	"sosservice/src/configurations/rest_err"
	"sosservice/src/view"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
