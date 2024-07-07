package repository

import (
	"context"
	"os"
	"sosservice/src/configurations/logger"
	"sosservice/src/configurations/rest_err"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (ur *userRepository) DeleteUser(userId string) *rest_err.RestErr {
	logger.Info("Starting Delete User Repository", zap.String("journey", "DeleteUser"))
	collection_name := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.databaseConnection.Collection(collection_name)

	userIdHex, _ := primitive.ObjectIDFromHex(userId)

	filter := bson.D{{Key: "_id", Value: userIdHex}}

	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		logger.Error("Error trying to delete user", err, zap.String("journey", "DeleteUser"))
		return rest_err.NewInternalServerError(err.Error())
	}
	logger.Info("User Deleted Fuccessfully", zap.String("journey", "DeleteUser"))
	return nil
}
