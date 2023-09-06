package api

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http/httptest"
	"testing"

	"github.com/ck3g/SomeDaysOfGo/hotel-reservation/db"
	"github.com/ck3g/SomeDaysOfGo/hotel-reservation/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	testdburi  = "mongodb://localhost:27017"
	testdbname = "hotel-reservation-test"
)

type testdb struct {
	db.UserStore
}

func setup(t *testing.T) *testdb {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(testdburi))
	if err != nil {
		log.Fatal(err)
	}

	return &testdb{
		UserStore: db.NewMongoUserStore(client, testdbname),
	}
}

func (tdb *testdb) teardown(t *testing.T) {
	if err := tdb.Drop(context.Background()); err != nil {
		t.Fatal(err)
	}
}

func Test_HandlePostUser(t *testing.T) {
	tdb := setup(t)
	defer tdb.teardown(t)

	app := fiber.New()
	userHandler := NewUserHandler(tdb.UserStore)
	app.Post("/", userHandler.HandlePostUser)

	params := types.CreateUserParams{
		Email:     "some@foo.com",
		FirstName: "John",
		LastName:  "Doe",
		Password:  "password",
	}

	b, _ := json.Marshal(params)

	req := httptest.NewRequest("POST", "/", bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")

	resp, _ := app.Test(req)

	var user types.User
	json.NewDecoder(resp.Body).Decode(&user)

	if len(user.ID) == 0 {
		t.Errorf("Expected user ID to be set, got %s", user.ID)
	}

	if len(user.EncryptedPassword) > 0 {
		t.Errorf("Expected encrypted password not to be included, got %s", user.EncryptedPassword)
	}

	if user.FirstName != params.FirstName {
		t.Errorf("Expected firstName to be %s, got %s", params.FirstName, user.FirstName)
	}

	if user.LastName != params.LastName {
		t.Errorf("Expected lastName to be %s, got %s", params.LastName, user.LastName)
	}

	if user.Email != params.Email {
		t.Errorf("Expected email to be %s, got %s", params.Email, user.Email)
	}
}
