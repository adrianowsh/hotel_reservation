package api

import (
	"github.com/adrianowsh/hotel-reservation/types"

	"github.com/gofiber/fiber/v3"
)

func HandleGetUserByEmail(c fiber.Ctx) error {
	email := c.Params("email")
	return c.JSON(map[string]string{"message": "Hello, Get User by Email!", "email": email})
}

func HandleGetUsers(c fiber.Ctx) error {

	u := types.User{
		ID:        "1",
		Email:     "example@example.com",
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
	}
	return c.JSON(u)
}

func HandleGetUser(c fiber.Ctx) error {
	id := c.Params("id")
	return c.JSON(map[string]string{"message": "Hello, Get User by ID!", "id": id})
}
func HandleCreateUser(c fiber.Ctx) error {
	return c.JSON(map[string]string{"message": "Hello, Create User!"})
}

func HandleUpdateUser(c fiber.Ctx) error {
	id := c.Params("id")
	return c.JSON(map[string]string{"message": "Hello, Update User!", "id": id})
}

func HandleDeleteUser(c fiber.Ctx) error {
	id := c.Params("id")
	return c.JSON(map[string]string{"message": "Hello, Delete User!", "id": id})
}
