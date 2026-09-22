package api

import (
	"context"

	"github.com/adrianowsh/hotel-reservation/db"

	"github.com/gofiber/fiber/v3"
)

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) HandleGetUserByID(c fiber.Ctx) error {
	var (
		id  string
		ctx context.Context
	)

	id = c.Params("id")
	ctx = context.Background()

	user, err := h.userStore.GetUserByID(ctx, id)
	if err != nil {
		return err
	}
	return c.JSON(user)
}

// func HandleGetUserByEmail(c fiber.Ctx) error {
// 	email := c.Params("email")
// 	return c.JSON(map[string]string{"message": "Hello, Get User by Email!", "email": email})
// }

// func HandleGetUsers(c fiber.Ctx) error {

// 	u := types.User{
// 		ID:        "1",
// 		Email:     "example@example.com",
// 		FirstName: "John",
// 		LastName:  "Doe",
// 		Age:       30,
// 	}
// 	return c.JSON(u)
// }

// func HandleCreateUser(c fiber.Ctx) error {
// 	return c.JSON(map[string]string{"message": "Hello, Create User!"})
// }

// func HandleUpdateUser(c fiber.Ctx) error {
// 	id := c.Params("id")
// 	return c.JSON(map[string]string{"message": "Hello, Update User!", "id": id})
// }

// func HandleDeleteUser(c fiber.Ctx) error {
// 	id := c.Params("id")
// 	return c.JSON(map[string]string{"message": "Hello, Delete User!", "id": id})
// }
