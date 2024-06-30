package main

import (
	"log"
	"os"
	"sosservice/src/configurations"
	"sosservice/src/configurations/database/postgres"
	"sosservice/src/configurations/logger"
	"sosservice/src/controller"
	"sosservice/src/controller/routes"
	"sosservice/src/model/service"

	"github.com/joho/godotenv"
)

func main() {
	logger.Info("Starting application")

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	config := &postgres.Config{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}
	postgres.InitDatabase(config)

	// models.MigrateUsers(storage.DB)
	// models.MigrateCauses(storage.DB)
	// models.MigrateUserInfos(storage.DB)

	// init dependencies
	service := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)

	router := configurations.SetupRouter()

	routes.InitializeRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
