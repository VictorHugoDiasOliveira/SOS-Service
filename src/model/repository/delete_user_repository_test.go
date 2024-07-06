package repository

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestDeleteUserRepository(t *testing.T) {
	database_name := "user_database_test"
	collection_name := "user_collection_test"

	err := os.Setenv("MONGODB_USER_COLLECTION", collection_name)
	if err != nil {
		t.FailNow()
		return
	}
	defer os.Clearenv()

	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	// defer mtestDb.Close()

	mtestDb.Run("valid id return success", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCursorResponse(
			1,
			fmt.Sprintf("%s.%s", database_name, collection_name),
			mtest.FirstBatch,
		))

		databaseMock := mt.Client.Database(database_name)

		repo := NewUserRepository(databaseMock)
		err := repo.DeleteUser(primitive.NewObjectID().Hex())

		assert.Nil(t, err)
	})

	mtestDb.Run("mongodb return error", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databaseMock := mt.Client.Database(database_name)

		repo := NewUserRepository(databaseMock)
		err := repo.DeleteUser(primitive.NewObjectID().Hex())

		assert.NotNil(t, err)
	})

	mtestDb.Run("user not found", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCursorResponse(
			0,
			fmt.Sprintf("%s.%s", database_name, collection_name),
			mtest.FirstBatch,
		))

		databaseMock := mt.Client.Database(database_name)

		repo := NewUserRepository(databaseMock)
		err := repo.DeleteUser(primitive.NewObjectID().Hex())

		assert.Nil(t, err)
	})
}
