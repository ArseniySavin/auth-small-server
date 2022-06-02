package handlers

import (
	"context"
	catcher "github.com/ArseniySavin/catcher/pkg"
	"github.com/gofiber/fiber/v2"
)

// TokenHandler -
func (auth *DefaultOAuth) TokenPost(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	clientId := c.Locals("clientId").(string)
	scopes := c.FormValue("scope", "")

	result, err := auth.controller.MakeToken(ctx, clientId, scopes)
	if err != nil {
		catcher.LogError(err)
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(result)
}
