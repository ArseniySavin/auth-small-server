package handlers

import (
	"context"
	"github.com/ArseniySavin/auth-small-server/typos"
	catcher "github.com/ArseniySavin/catcher/pkg"
	"github.com/gofiber/fiber/v2"
)

// ClientPost -
func (auth *DefaultOAuth) ClientPost(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	clientRequest := new(typos.ClientRequest)

	if err := c.BodyParser(clientRequest); err != nil {
		catcher.LogError(err)
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err := auth.controller.MakeClientGrantScopeClaim(ctx, *clientRequest)

	if err != nil {
		catcher.LogError(err)
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON("Ok")
}

// ClientPatch -
func (auth *DefaultOAuth) ClientPatch(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	clientRequest := new(typos.ClientRequest)

	if err := c.BodyParser(clientRequest); err != nil {
		catcher.LogError(err)
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err := auth.controller.UpdateClientGrantScopeClaim(ctx, *clientRequest)

	if err != nil {
		catcher.LogError(err)
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON("Ok")
}

// ClientPut -
func (auth *DefaultOAuth) ClientPut(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	clientRequest := new(typos.ClientRequest)

	if err := c.BodyParser(clientRequest); err != nil {
		catcher.LogError(err)
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err := auth.controller.MakeClientGrant(ctx, *clientRequest)

	if err != nil {
		catcher.LogError(err)
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON("Ok")
}
