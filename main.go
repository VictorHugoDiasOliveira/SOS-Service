package main

import (
	"log"
	"sosservice/src/configurations"
	"sosservice/src/configurations/logger"
	"sosservice/src/controller"
	"sosservice/src/controller/routes"
	"sosservice/src/model/service"

	"github.com/joho/godotenv"
)

func main() {
	logger.Info("Starting Application")

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	// init dependencies
	service := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)

	router := configurations.SetupRouter()

	routes.InitializeRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
