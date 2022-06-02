package handlers

import (
	"context"
	catcher "github.com/ArseniySavin/catcher/pkg"
	"github.com/gofiber/fiber/v2"
)

// RevokePost -
func (auth *DefaultOAuth) RevokePost(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	token := c.FormValue("token", "")

	err := auth.controller.DeleteToken(ctx, token)
	if err != nil {
		catcher.LogError(err)
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return nil
}
