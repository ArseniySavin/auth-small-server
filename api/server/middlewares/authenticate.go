package middlewares

import (
	"context"
	"github.com/ArseniySavin/auth-small-server/internal/controllers"
	"github.com/ArseniySavin/auth-small-server/internal/pkg/auth"
	"github.com/ArseniySavin/auth-small-server/typos"
	catcher "github.com/ArseniySavin/catcher/pkg"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"strings"
)

// AddAuthenticate -
func AddAuthenticate(server *fiber.App, auth auth.IAuthServices, controller controllers.IControllers) {
	server.Use("/v1", basic(auth))
	//server.Use("/v2", token(controller))

}

func basic(auth auth.IAuthServices) fiber.Handler {
	return basicauth.New(basicauth.Config{
		ContextUsername: "clientId",
		ContextPassword: "clientSecret",
		Authorizer: func(clientId, clientSecret string) bool {
			ok, err := auth.Authenticate(clientId, clientSecret)
			if err != nil {
				catcher.LogError(err)
			}

			return ok
		},
	})
}

func token(controller controllers.IControllers) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		token := ctx.Get("Authorization")
		if token == "" {
			catcher.LogError(typos.ErrEmptyToken)
			return typos.ErrEmptyToken
		}

		token = strings.Replace(token, "bearer ", "", 1)

		ctxReq, cancel := context.WithCancel(context.Background())
		defer cancel()

		result, err := controller.GetToken(ctxReq, token)
		if err != nil {
			catcher.LogError(err)
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		if result == nil {
			catcher.LogError(typos.ErrTokenIsDemolished)
			return fiber.NewError(fiber.StatusBadRequest, typos.ErrTokenIsDemolished.Error())
		}

		if active := result["active"].(bool); !active {
			catcher.LogError(typos.ErrInvalidToken)
			return fiber.NewError(fiber.StatusBadRequest, typos.ErrInvalidToken.Error())
		}
		ctx.Next()
		return nil
	}
}
