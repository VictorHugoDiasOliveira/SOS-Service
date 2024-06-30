package controller

import (
	"sosservice/src/model/service"

	"github.com/gin-gonic/gin"
)

func NewUserControllerInterface(serviceInterface service.UserDomainService) UserControllerInterface {
	return &userControllerInterface{
		service: serviceInterface,
	}
}

type UserControllerInterface interface {
	FindUserById(context *gin.Context)
	FindUserByEmail(context *gin.Context)
	DeleteUser(context *gin.Context)
	UpdateUser(context *gin.Context)
	CreateUser(context *gin.Context)
}

type userControllerInterface struct {
	service service.UserDomainService
}
