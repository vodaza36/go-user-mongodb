package mongo_test

import (
	"log"
	"testing"

	"github.com/vodaza36/go-user-mongodb/pck"

	"github.com/vodaza36/go-user-mongodb/pck/mock"
	"github.com/vodaza36/go-user-mongodb/pck/mongo"
)

const (
	mongoUrl           = "localhost:27017"
	dbName             = "test_db"
	userCollectionName = "user"
)

func Test_UserService(t *testing.T) {
	t.Run("CreateUser", createUser_should_insert_user_into_mongo)
}

func createUser_should_insert_user_into_mongo(t *testing.T) {
	//Arrange
	session, err := mongo.NewSession(mongoUrl)
	if err != nil {
		log.Fatalf("Unable to connect to mongo: %s", err)
	}
	defer func() {
		session.DropDatabase(dbName)
		session.Close()
	}()
	mockHash := mock.Hash{}
	userService := mongo.NewUserService(session.Copy(), dbName, userCollectionName, &mockHash)

	testUsername := "integration_test_user"
	testPassword := "integration_test_password"
	user := root.User{
		Username: testUsername,
		Password: testPassword}

	//Act
	err = userService.CreateUser(&user)

	//Assert
	if err != nil {
		t.Errorf("Unable to create user: %s", err)
	}
	var results []root.User
	session.GetCollection(dbName, userCollectionName).Find(nil).All(&results)

	count := len(results)
	if count != 1 {
		t.Error("Incorrect number of results. Expected `1`, got: `%i`", count)
	}
	if results[0].Username != user.Username {
		t.Errorf("Incorrect Username. Expected `%s`, Got: `%s`", testUsername, results[0].Username)
	}
}
