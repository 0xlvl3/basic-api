package handles

import (
	"fmt"

	"github.com/0xlvl3/basic-api/api/db"
	"github.com/0xlvl3/basic-api/api/types"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) HandlePostUser(c *fiber.Ctx) error {
	var user types.User

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	insertUser, err := h.userStore.PostUser(c.Context(), &user)
	if err != nil {
		return err
	}

	return c.JSON(insertUser)
}

func (h *UserHandler) HandleGetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := h.userStore.GetUser(c.Context(), id)
	if err != nil {
		return err
	}

	fmt.Printf("%+v", user)
	return c.JSON(user)
}
