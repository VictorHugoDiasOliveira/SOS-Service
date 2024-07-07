package tests

import (
	"fmt"
	"os"
	"sosservice/src/model/repository"
	"sosservice/src/model/repository/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestFindUserByEmailRepository(t *testing.T) {
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

	mtestDb.Run("Successfully found user by email", func(mt *mtest.T) {
		userEntity := entity.UserEntity{
			ID:       primitive.NewObjectID(),
			Email:    "victor@gmail.com",
			Password: "senha123",
			Name:     "victor",
			Age:      22,
		}
		mt.AddMockResponses(mtest.CreateCursorResponse(
			1,
			fmt.Sprintf("%s.%s", database_name, collection_name),
			mtest.FirstBatch,
			convertEntityToBson(userEntity),
		))

		databaseMock := mt.Client.Database(database_name)

		repo := repository.NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByEmail(userEntity.Email)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomain.GetID(), userEntity.ID.Hex())
		assert.EqualValues(t, userDomain.GetEmail(), userEntity.Email)
		assert.EqualValues(t, userDomain.GetPassword(), userEntity.Password)
		assert.EqualValues(t, userDomain.GetName(), userEntity.Name)
		assert.EqualValues(t, userDomain.GetAge(), userEntity.Age)
	})

	mtestDb.Run("Error trying to connect to database", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databaseMock := mt.Client.Database(database_name)

		repo := repository.NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByEmail("victor@gmail.com")

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})

	mtestDb.Run("User not found", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCursorResponse(
			0,
			fmt.Sprintf("%s.%s", database_name, collection_name),
			mtest.FirstBatch,
		))

		databaseMock := mt.Client.Database(database_name)

		repo := repository.NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByEmail("victor@gmail.com")

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})
}

func TestFindUserByIdRepository(t *testing.T) {
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

	mtestDb.Run("Successfully found user by ID", func(mt *mtest.T) {
		userEntity := entity.UserEntity{
			ID:       primitive.NewObjectID(),
			Email:    "victor@gmail.com",
			Password: "senha123",
			Name:     "victor",
			Age:      22,
		}
		mt.AddMockResponses(mtest.CreateCursorResponse(
			1,
			fmt.Sprintf("%s.%s", database_name, collection_name),
			mtest.FirstBatch,
			convertEntityToBson(userEntity),
		))

		databaseMock := mt.Client.Database(database_name)

		repo := repository.NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserById(userEntity.ID.Hex())

		assert.Nil(t, err)
		assert.EqualValues(t, userDomain.GetID(), userEntity.ID.Hex())
		assert.EqualValues(t, userDomain.GetEmail(), userEntity.Email)
		assert.EqualValues(t, userDomain.GetPassword(), userEntity.Password)
		assert.EqualValues(t, userDomain.GetName(), userEntity.Name)
		assert.EqualValues(t, userDomain.GetAge(), userEntity.Age)
	})

	mtestDb.Run("Error trying to connect to database", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databaseMock := mt.Client.Database(database_name)

		repo := repository.NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserById("6685ca7778824239f21f069d")

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})

	mtestDb.Run("User not found", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCursorResponse(
			0,
			fmt.Sprintf("%s.%s", database_name, collection_name),
			mtest.FirstBatch,
		))

		databaseMock := mt.Client.Database(database_name)

		repo := repository.NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserById("6685ca7778824239f21f069d")

		assert.NotNil(t, err)
		assert.Equal(t, err.Message, fmt.Sprintf("User not found: %s", "6685ca7778824239f21f069d"))
		assert.Nil(t, userDomain)
	})
}

func TestFindUserByEmailAndPasswordRepository(t *testing.T) {
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

	mtestDb.Run("Successfully found user by email and password", func(mt *mtest.T) {
		userEntity := entity.UserEntity{
			ID:       primitive.NewObjectID(),
			Email:    "victor@gmail.com",
			Password: "senha123",
			Name:     "victor",
			Age:      22,
		}
		mt.AddMockResponses(mtest.CreateCursorResponse(
			1,
			fmt.Sprintf("%s.%s", database_name, collection_name),
			mtest.FirstBatch,
			convertEntityToBson(userEntity),
		))

		databaseMock := mt.Client.Database(database_name)

		repo := repository.NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByEmailAndPassword(userEntity.Email, userEntity.Password)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomain.GetID(), userEntity.ID.Hex())
		assert.EqualValues(t, userDomain.GetEmail(), userEntity.Email)
		assert.EqualValues(t, userDomain.GetPassword(), userEntity.Password)
		assert.EqualValues(t, userDomain.GetName(), userEntity.Name)
		assert.EqualValues(t, userDomain.GetAge(), userEntity.Age)
	})

	mtestDb.Run("Error trying to connect to database", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databaseMock := mt.Client.Database(database_name)

		repo := repository.NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByEmailAndPassword("victor@gmail.com", "senha1234")

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})

	mtestDb.Run("User not found", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCursorResponse(
			0,
			fmt.Sprintf("%s.%s", database_name, collection_name),
			mtest.FirstBatch,
		))

		databaseMock := mt.Client.Database(database_name)

		repo := repository.NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByEmailAndPassword("victor@gmail.com", "senha1234")

		assert.NotNil(t, err)
		assert.Equal(t, err.Message, "Invalid Email or Password")
		assert.Nil(t, userDomain)
	})
}

func convertEntityToBson(userEntity entity.UserEntity) bson.D {
	return bson.D{
		{Key: "_id", Value: userEntity.ID},
		{Key: "email", Value: userEntity.Email},
		{Key: "password", Value: userEntity.Password},
		{Key: "name", Value: userEntity.Name},
		{Key: "age", Value: userEntity.Age},
	}
}
