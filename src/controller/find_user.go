package controller

import (
	"net/http"
	"net/mail"
	"sosservice/src/configurations/logger"
	"sosservice/src/configurations/rest_err"
	"sosservice/src/view"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc userControllerInterface) FindUserById(context *gin.Context) {
	logger.Info("Starting FindUserById Controller", zap.String("journey", "FindUserById"))

	userId := context.Param("id")

	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errorMessage := rest_err.NewBadRequestError("ID is not a valid ID")
		logger.Error("Error trying to validate ID", err, zap.String("journey", "FindUserByIdService"))
		context.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByIdService(userId)
	if err != nil {
		logger.Error("Error trying to call service FindUserByIdService", err, zap.String("journey", "FindUserByIdService"))
		context.JSON(err.Code, err)
		return
	}
	logger.Info("Finished FindUserById Controller", zap.String("journey", "FindUserById"))
	context.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}

func (uc userControllerInterface) FindUserByEmail(context *gin.Context) {
	logger.Info("Starting FindUserByEmail Controller",
		zap.String("journey", "FindUserByEmail"),
	)

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
	logger.Info("Finished FindUserByEmail Controller", zap.String("journey", "FindUserByEmail"))
	context.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}
