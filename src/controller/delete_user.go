package controller

import (
	"net/http"
	"sosservice/src/configurations/rest_err"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DeleteUser deletes a user using id
// @Summary Delete User
// @Description Deletes a user based on the ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "ID of the user to be deleted"
// @Success 200
// @Param Authorization header string true "Insert your access token" default(Bearer <token>)
// @Failure 400 {object} rest_err.RestErr
// @Failure 500 {object} rest_err.RestErr
// @Router /deleteUser/{id} [delete]
func (uc *userControllerInterface) DeleteUser(context *gin.Context) {

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

	context.Status(http.StatusOK)
}
