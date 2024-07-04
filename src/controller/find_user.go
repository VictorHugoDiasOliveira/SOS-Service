package controller

import (
	"net/http"
	"net/mail"
	"sosservice/src/configurations/logger"
	"sosservice/src/configurations/rest_err"
	"sosservice/src/model"
	"sosservice/src/view"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) FindUserById(context *gin.Context) {

	user, err := model.ValidateToken(context.Request.Header.Get("Authorization"))
	if err != nil {
		context.JSON(err.Code, err)
		return
	}

	logger.Info("Mal Feito, Feito", zap.String("Email", user.GetEmail()))

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
