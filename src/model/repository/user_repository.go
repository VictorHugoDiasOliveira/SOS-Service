package repository

import (
	"sosservice/src/configurations/rest_err"
	"sosservice/src/model"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	MONGODB_USER_COLLECTION = "MONGODB_USER_COLLECTION"
)

func NewUserRepository(database *mongo.Database) UserRepository {
	return &userRepository{
		databaseConnection: database,
	}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepository interface {
	CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateUser(id string, userDomain model.UserDomainInterface) *rest_err.RestErr
	FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserById(id string) (model.UserDomainInterface, *rest_err.RestErr)
	DeleteUser(id string) *rest_err.RestErr

	FindUserByEmailAndPassword(email string, password string) (model.UserDomainInterface, *rest_err.RestErr)
}
