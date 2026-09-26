package api

import (
	"errors"

	"github.com/adrianowsh/hotel-reservation/db"
	"github.com/adrianowsh/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/mongo"

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
	id := c.Params("id")
	user, err := h.userStore.GetUserByID(c.Context(), id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.Status(fiber.StatusNotFound).JSON(types.ErrorResponse{Errors: map[string]string{"user": "not found"}})
		}
		return err
	}
	return c.JSON(types.UserResponse{Data: user})
}

func (h *UserHandler) HandleGetUserByEmail(c fiber.Ctx) error {
	email := c.Params("email")
	user, err := h.userStore.GetUserByEmail(c.Context(), email)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.Status(fiber.StatusNotFound).JSON(types.ErrorResponse{Errors: map[string]string{"user": "not found"}})
		}
		return err
	}
	return c.JSON(types.UserResponse{Data: user})
}

func (h *UserHandler) HandleGetUsers(c fiber.Ctx) error {
	users, err := h.userStore.GetUsers(c.Context())
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(types.UsersResponse{Data: users})
}

func (h *UserHandler) HandleCreateUser(c fiber.Ctx) error {
	var params types.CreateUserParam
	if err := c.Bind().Body(&params); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if errors := params.Validate(); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}
	user, err := types.NewUserFromParam(&params)
	if err != nil {
		return err
	}
	createdUser, err := h.userStore.CreateUser(c.Context(), user)
	if err != nil {
		return err
	}
	return c.JSON(types.UserResponse{Data: createdUser})
}

func (h *UserHandler) HandleUpdateUser(c fiber.Ctx) error {
	id := c.Params("id")
	var params types.CreateUserParam
	if err := c.Bind().Body(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	if err := params.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	_, err := h.userStore.GetUserByID(c.Context(), id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.Status(fiber.StatusNotFound).JSON(types.ErrorResponse{Errors: map[string]string{"user": "not found"}})
		}
		return err
	}

	user, err := types.NewUserFromParam(&params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{Errors: map[string]string{"internal": err.Error()}})
	}
	updatedUser, err := h.userStore.UpdateUser(c.Context(), id, user)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(types.UserResponse{Data: updatedUser})
}

func (h *UserHandler) HandleDeleteUser(c fiber.Ctx) error {
	id := c.Params("id")
	if err := h.userStore.DeleteUser(c.Context(), id); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(map[string]string{"message": "User deleted successfully", "id": id})
}
