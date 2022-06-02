package handlers

import (
	"context"
	catcher "github.com/ArseniySavin/catcher/pkg"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

// RevocationPost -
func (auth *DefaultOAuth) RevocationPost(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	clientId := c.FormValue("client_id", "")
	active, err := strconv.ParseBool(c.FormValue("active", ""))
	if err != nil {
		catcher.LogError(err)
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err = auth.controller.ChangeClientState(ctx, clientId, active)

	if err != nil {
		catcher.LogError(err)
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON("Ok")
}
