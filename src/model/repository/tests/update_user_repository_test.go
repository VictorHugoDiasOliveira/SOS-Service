package tests

import (
	"os"
	"sosservice/src/model"
	"sosservice/src/model/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestUpdateUserRepository(t *testing.T) {
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

	mtestDb.Run("Successfully updated user by ID", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "n", Value: 1},
			{Key: "acknowledged", Value: true},
		})

		databaseMock := mt.Client.Database(database_name)
		repo := repository.NewUserRepository(databaseMock)

		domain := model.NewUserUpdateDomain("victor", 22)
		err := repo.UpdateUser("6685ca7778824239f21f069d", domain)

		assert.Nil(t, err)
	})

	mtestDb.Run("Error trying to connect to database", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databaseMock := mt.Client.Database(database_name)

		repo := repository.NewUserRepository(databaseMock)
		domain := model.NewUserUpdateDomain("victor", 18)
		err := repo.UpdateUser("6685ca7778824239f21f069d", domain)

		assert.NotNil(t, err)
	})
}
