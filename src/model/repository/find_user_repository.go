package repository

import (
	"context"
	"fmt"
	"os"
	"sosservice/src/configurations/logger"
	"sosservice/src/configurations/rest_err"
	"sosservice/src/model"
	"sosservice/src/model/repository/entity"
	"sosservice/src/model/repository/entity/converter"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (ur *userRepository) FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Finding User By Email", zap.String("journey", "FindUserByEmail"))
	collection_name := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "email", Value: email}}

	err := collection.FindOne(context.Background(), filter).Decode(userEntity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User not found: %s", email)
			return nil, rest_err.NewNotFoundError(errorMessage)
		}
		logger.Error("Error trying to find user by email", err, zap.String("journey", "FindUserByEmail"))
		errorMessage := "Error trying to find user by email"
		return nil, rest_err.NewInternalServerError(errorMessage)
	}
	logger.Info("User Found By Email", zap.String("journey", "FindUserByEmail"))
	return converter.ConvertEntityToDomain(*userEntity), nil
}

func (ur *userRepository) FindUserById(id string) (model.UserDomainInterface, *rest_err.RestErr) {
	collection_name := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	objectId, _ := primitive.ObjectIDFromHex(id)

	filter := bson.D{{Key: "_id", Value: objectId}}

	err := collection.FindOne(context.Background(), filter).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User not found: %s", id)
			return nil, rest_err.NewNotFoundError(errorMessage)
		}

		errorMessage := "Error trying to find user by id"
		return nil, rest_err.NewInternalServerError(errorMessage)
	}
	return converter.ConvertEntityToDomain(*userEntity), nil
}

func (ur *userRepository) FindUserByEmailAndPassword(email string, password string) (model.UserDomainInterface, *rest_err.RestErr) {
	collection_name := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	filter := bson.D{
		{Key: "email", Value: email},
		{Key: "password", Value: password},
	}

	err := collection.FindOne(context.Background(), filter).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := "Invalid Email or Password"
			return nil, rest_err.NewForbiddenError(errorMessage)
		}

		errorMessage := "Error trying to find user"
		return nil, rest_err.NewInternalServerError(errorMessage)
	}
	return converter.ConvertEntityToDomain(*userEntity), nil
}
