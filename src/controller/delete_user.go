package controller

import (
	"net/http"
	"sosservice/src/configurations/logger"
	"sosservice/src/configurations/rest_err"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) DeleteUser(context *gin.Context) {
	logger.Info("Init DeleteUser Controller",
		zap.String("journey", "DeleteUser"),
	)

	userId := context.Param("id")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errRest := rest_err.NewBadRequestError("Invalid User Id, must be a hex value")
		context.JSON(errRest.Code, errRest)
		return
	}

	err := uc.service.DeleteUserService(userId)
	if err != nil {
		context.JSON(err.Code, err)
		return
	}

	logger.Info("User Deleted successfully",
		zap.String("journey", "DeleteUser"),
	)

	context.Status(http.StatusOK)
}
