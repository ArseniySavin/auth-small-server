package middlewares

import (
	"github.com/ArseniySavin/auth-small-server/internal/pkg/auth"
	catcher "github.com/ArseniySavin/catcher/pkg"
	"path"

	"github.com/gofiber/fiber/v2"
)

// AddAuthorizate -
func AddAuthorizate(server *fiber.App, auth auth.IAuthServices) {
	server.Use("/v1", func(c *fiber.Ctx) error {
		clientId := c.Locals("clientId").(string)
		endpoint := path.Base(c.Path())

		if err := auth.AuthorizateGrants(clientId, endpoint); err != nil {
			catcher.LogError(err)
			return fiber.NewError(fiber.StatusForbidden, err.Error())
		}

		return c.Next()
	})

	server.Use("/v1/token", func(c *fiber.Ctx) error {
		scopes := c.FormValue("scope", "")
		clientId := c.Locals("clientId").(string)

		if err := auth.Authorizate(clientId, scopes); err != nil {
			catcher.LogError(err)
			return fiber.NewError(fiber.StatusForbidden, err.Error())
		}

		return c.Next()
	})
}
