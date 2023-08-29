package api

import (
	"github.com/ck3g/SomeDaysOfGo/hotel-reservation/types"
	"github.com/gofiber/fiber/v2"
)

func HandleGetUsers(c *fiber.Ctx) error {
	u := types.User{
		ID:        "1",
		FirstName: "John",
		LastName:  "Doe",
	}
	return c.JSON(u)
}

func HandleGetUser(c *fiber.Ctx) error {
	u := types.User{
		ID:        "1",
		FirstName: "John",
		LastName:  "Doe",
	}
	return c.JSON(u)
}
