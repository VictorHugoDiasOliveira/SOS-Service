package configurations

import (
	"sosservice/src/configurations/logger"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SetupRouter() *gin.Engine {
	logger.Info("Starting Router", zap.String("journey", "SetupRouter"))
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Lenght"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	logger.Info("Router Started Successfully", zap.String("journey", "SetupRouter"))
	return router
}
