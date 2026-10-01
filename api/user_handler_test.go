package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"net/http/httptest"

	"github.com/adrianowsh/hotel-reservation/db"
	"github.com/adrianowsh/hotel-reservation/types"
	"github.com/gofiber/fiber/v3"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	db_test_uri  = "mongodb://localhost:27017"
	db_test_name = "hotel_reservation_test"
)

type testDb struct {
	db.UserStore
}

func setup() *testDb {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db_test_uri))
	if err != nil {
		log.Fatal(err)
	}

	return &testDb{
		UserStore: db.NewMongoUserStore(client, db_test_name),
	}
}

func teardown(dbTest *testDb) {
	fmt.Println("...dropping user collection")
	if err := dbTest.UserStore.Drop(context.TODO(), db_test_name); err != nil {
		log.Fatal(err)
	}
}

func TestPostUser(t *testing.T) {
	dbTest := setup()
	defer teardown(dbTest)

	app := fiber.New()
	userHandler := NewUserHandler(dbTest.UserStore)
	app.Post("/", userHandler.HandleCreateUser)

	params := types.CreateUserParam{
		Email:     "test@example.com",
		Password:  "password123",
		FirstName: "Test",
		LastName:  "User",
		Age:       20,
	}

	body, _ := json.Marshal(params)
	req := httptest.NewRequest("POST", "/", bytes.NewReader(body))
	req.Header.Add("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("test request failed: %v", err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}

	var userResp types.UserResponse
	json.NewDecoder(resp.Body).Decode(&userResp)
	user := userResp.Data
	fmt.Printf("Created user: %+v\n", user)

	if user == nil {
		t.Fatalf("expected user data, got nil")
	}
	if user.Email != "test@example.com" {
		t.Fatalf("expected email 'test@example.com', got '%s'", user.Email)
	}
	if user.FirstName != "Test" {
		t.Fatalf("expected first name 'Test', got '%s'", user.FirstName)
	}
	if user.LastName != "User" {
		t.Fatalf("expected last name 'User', got '%s'", user.LastName)
	}
	if user.Age != 20 {
		t.Fatalf("expected age 20, got %d", user.Age)
	}
}
