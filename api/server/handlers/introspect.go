package handlers

import (
	"context"
	catcher "github.com/ArseniySavin/catcher/pkg"
	"github.com/gofiber/fiber/v2"
)

// IntrospectPost -
func (auth *DefaultOAuth) IntrospectPost(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	token := c.FormValue("token", "")

	result, err := auth.controller.GetToken(ctx, token)
	if err != nil {
		catcher.LogError(err)
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(result)
}
