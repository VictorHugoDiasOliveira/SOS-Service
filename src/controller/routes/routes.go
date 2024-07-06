package routes

import (
	"sosservice/src/configurations/logger"
	"sosservice/src/controller"
	"sosservice/src/model"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

func InitializeRoutes(rg *gin.RouterGroup, userController controller.UserControllerInterface) {
	logger.Info("Starting Routes", zap.String("journey", "InitializeRoutes"))

	rg.GET("/getUserById/:id", model.ValidateTokenMiddleware, userController.FindUserById)
	rg.GET("/getUserByEmail/:email", model.ValidateTokenMiddleware, userController.FindUserByEmail)
	rg.POST("/createUser", userController.CreateUser)
	rg.PUT("/updateUser/:id", userController.UpdateUser)
	rg.DELETE("/deleteUser/:id", model.ValidateTokenMiddleware, userController.DeleteUser)

	rg.POST("/login", userController.LoginUser)

	rg.GET("swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	logger.Info("Routes Started Successfully", zap.String("journey", "InitializeRoutes"))
}
