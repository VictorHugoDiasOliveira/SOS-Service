package controller

import (
	"net/http"
	"sosservice/src/configurations/rest_err"
	"sosservice/src/configurations/validation"
	"sosservice/src/controller/model/request"
	"sosservice/src/model"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (uc *userControllerInterface) UpdateUser(context *gin.Context) {
	userId := context.Param("id")
	_, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		errRest := rest_err.NewBadRequestError("Invalid User Id, must be a hex value")
		context.JSON(errRest.Code, errRest)
		return
	}

	var userRequest request.UserUpdateRequest

	if err := context.ShouldBindJSON(&userRequest); err != nil {
		restErr := validation.ValidateUserError(err)
		context.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserUpdateDomain(userRequest.Name, userRequest.Age)
	if err := uc.service.UpdateUserService(userId, domain); err != nil {
		context.JSON(err.Code, err)
		return
	}
	context.Status(http.StatusOK)
}
