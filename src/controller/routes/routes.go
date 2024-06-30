package routes

import (
	"sosservice/src/controller"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(rg *gin.RouterGroup, userController controller.UserControllerInterface) {

	rg.GET("/getUserById/:id", userController.FindUserById)
	rg.POST("/createUser", userController.CreateUser)
	rg.PUT("/updateUser/:id", userController.UpdateUser)
	rg.DELETE("/deleteUser/:id", userController.DeleteUser)
}
