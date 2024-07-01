package main

import (
	"sosservice/src/controller"
	"sosservice/src/model/repository"
	"sosservice/src/model/service"

	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(database *mongo.Database) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(service)
}
