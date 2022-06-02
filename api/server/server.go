package server

import (
	"context"
	"github.com/ArseniySavin/auth-small-server/config"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

// NewServer -
func NewServer() *fiber.App {
	return fiber.New(fiber.Config{
		Prefork: false,
	})
}

// Run -
func Run(lc fx.Lifecycle, server *fiber.App, cfg *config.Config) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				log.Println("Starting service...")
				if err := server.Listen(cfg.Port); err != nil {
					log.Fatalf("Run | %s", err.Error())
				}
			}()

			return nil
		},
		OnStop: func(context.Context) error {
			log.Println("Stoping service...")
			return server.Shutdown()
		},
	})
}
