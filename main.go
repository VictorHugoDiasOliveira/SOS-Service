package main

import (
	"context"
	"log"
	"sosservice/src/configurations"
	"sosservice/src/configurations/database/mongodb"
	"sosservice/src/configurations/logger"
	"sosservice/src/controller/routes"
)

func main() {
	logger.Info("Starting Application")
	// if err := godotenv.Load(".env"); err != nil {
	// 	log.Fatal(err)
	// }

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf("error trying to connect to database, error=%s \n", err.Error())
	}

	userController := initDependencies(database)

	router := configurations.SetupRouter()

	routes.InitializeRoutes(&router.RouterGroup, userController)

	logger.Info("Running Routes")
	if err := router.Run(":8080"); err != nil {
		logger.Error("Error Running Routes", err)
		log.Fatal(err)
	}
}
