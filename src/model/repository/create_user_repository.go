package repository

import (
	"context"
	"os"
	"sosservice/src/configurations/logger"
	"sosservice/src/configurations/rest_err"
	"sosservice/src/model"
	"sosservice/src/model/repository/entity/converter"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (ur *userRepository) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Creating User Repository", zap.String("journey", "CreateUser"))
	collection_name := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.databaseConnection.Collection(collection_name)

	value := converter.ConvertDomainToEntity(userDomain)

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		logger.Error("Failed to insert user into database", err, zap.String("journey", "CreateUser"))
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	value.ID = result.InsertedID.(primitive.ObjectID)

	logger.Info("User Repository Created Successfully", zap.String("journey", "CreateUser"))
	return converter.ConvertEntityToDomain(*value), nil
}
