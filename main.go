package main

import (
	"context"
	"flag"
	"log"

	"github.com/adrianowsh/hotel-reservation/api"
	"github.com/adrianowsh/hotel-reservation/db"
	"github.com/gofiber/fiber/v3"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const Port = ":5000"
const db_uri = "mongodb://localhost:27017"
const dbname = "hotel_reservation"
const usercoll = "users"

var config = fiber.Config{
	ErrorHandler: func(c fiber.Ctx, err error) error {
		return c.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {
	listenAddr := flag.String("listenAddr", Port, "Port to listen on")
	flag.Parse()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db_uri))
	if err != nil {
		log.Fatal(err)
	}

	// handlers
	userHandler := api.NewUserHandler(db.NewMongoUserStore(client))

	app := fiber.New(config)
	apiGroup := app.Group("/api/v1")

	//apiGroup.Get("/users", userHandler.HandleGetUsers)
	apiGroup.Get("/user/:id", userHandler.HandleGetUserByID)
	// apiGroup.Get("/user/email/:email", userHandler.HandleGetUserByEmail)
	// apiGroup.Post("/user", userHandler.HandleCreateUser)
	// apiGroup.Put("/user/:id", userHandler.HandleUpdateUser)
	//apiGroup.Delete("/user/:id", userHandler.HandleDeleteUser)

	app.Listen(*listenAddr)
}
