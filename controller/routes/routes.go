package routes

import (
	"sosservice/controller"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(rg *gin.RouterGroup) {

	rg.GET("/getUserById/:id", controller.FindUserById)
	rg.POST("/createUser", controller.CreateUser)
	rg.PUT("/updateUser/:id", controller.UpdateUser)
	rg.DELETE("/deleteUser/:id", controller.DeleteUser)
}
