package routes

import (
	"github.com/ArseniySavin/auth-small-server/api/server/handlers"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// Register -
func Register(server *fiber.App, oauth handlers.OAuth) {
	server.Add(http.MethodPost, "/v1/token", oauth.TokenPost)
	server.Add(http.MethodPost, "/v1/introspect", oauth.IntrospectPost)
	server.Add(http.MethodPost, "/v1/revoke", oauth.RevokePost)
	server.Add(http.MethodPost, "/v1/client", oauth.ClientPost)
	server.Add(http.MethodPatch, "/v1/client", oauth.ClientPatch)
	server.Add(http.MethodPut, "/v1/client", oauth.ClientPut)
	server.Add(http.MethodPost, "/v1/revocation", oauth.RevocationPost)
	server.Add(http.MethodGet, "/health", func(c *fiber.Ctx) error {
		return c.JSON("it`s ok!")
	})
}
