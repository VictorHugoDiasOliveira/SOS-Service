package main

import (
	"sosservice/src/configurations/logger"
	"sosservice/src/controller"
	"sosservice/src/model/repository"
	"sosservice/src/model/service"

	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func initDependencies(database *mongo.Database) controller.UserControllerInterface {
	logger.Info("Starting Dependencies", zap.String("journey", "initDependencies"))
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	logger.Info("Dependencies Started Successfully", zap.String("journey", "initDependencies"))
	return controller.NewUserControllerInterface(service)
}
