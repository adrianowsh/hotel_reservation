package main

import (
	"github.com/adrianowsh/hotel-reservation/api"
	"github.com/gofiber/fiber/v3"
)

const Port = ":5000"

func main() {
	app := fiber.New()

	app.Get("/users", api.HandleGetUsers)
	app.Get("/user/:id", api.HandleGetUser)
	app.Get("/user/email/:email", api.HandleGetUserByEmail)
	app.Post("/user", api.HandleCreateUser)
	app.Put("/user/:id", api.HandleUpdateUser)
	app.Delete("/user/:id", api.HandleDeleteUser)

	app.Listen(Port)
}
