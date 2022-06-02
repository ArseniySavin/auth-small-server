package handlers

import (
	"github.com/ArseniySavin/auth-small-server/internal/controllers"
	"github.com/gofiber/fiber/v2"
)

// OAuth -
type OAuth interface {
	TokenPost(c *fiber.Ctx) error
	RevokePost(c *fiber.Ctx) error
	IntrospectPost(c *fiber.Ctx) error
	RevocationPost(c *fiber.Ctx) error
	ClientPost(c *fiber.Ctx) error
	ClientPatch(c *fiber.Ctx) error
	ClientPut(c *fiber.Ctx) error
}

// DefaultOAuth -
type DefaultOAuth struct {
	controller controllers.IControllers
}

// NewOAuth -
func NewOAuth(controller controllers.IControllers) OAuth {
	return &DefaultOAuth{
		controller: controller,
	}
}
