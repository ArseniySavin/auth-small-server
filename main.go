package main

import (
	"github.com/ArseniySavin/auth-small-server/api/server"
	"github.com/ArseniySavin/auth-small-server/api/server/handlers"
	"github.com/ArseniySavin/auth-small-server/api/server/middlewares"
	"github.com/ArseniySavin/auth-small-server/api/server/routes"
	"github.com/ArseniySavin/auth-small-server/config"
	"github.com/ArseniySavin/auth-small-server/internal/controllers"
	"github.com/ArseniySavin/auth-small-server/internal/pkg/auth"
	"github.com/ArseniySavin/auth-small-server/internal/pkg/cache"
	"github.com/ArseniySavin/auth-small-server/internal/pkg/client"
	"github.com/ArseniySavin/auth-small-server/internal/pkg/jwt"
	"github.com/ArseniySavin/auth-small-server/internal/pkg/repository"
	"github.com/ArseniySavin/auth-small-server/internal/pkg/token"
	"github.com/ArseniySavin/auth-small-server/scripts"
	catcher "github.com/ArseniySavin/catcher/pkg"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

func main() {
	viper.AutomaticEnv()
	config.LoadDotEnv()

	trace := viper.GetBool("TRACING")
	mode := catcher.Regular

	if trace {
		mode = catcher.Tracing
	}
	catcher.Mode = mode

	defer cache.Close()

	app := fx.New(
		fx.Provide(
			config.NewConfig,
			server.NewServer,
			handlers.NewOAuth,
			controllers.NewControllers,
			repository.CreateDbConnection,
			repository.NewRepository,
			auth.NewAuthServices,
			token.NewTokenServices,
			client.NewClientServices,
			jwt.NewJwtServices,
		),

		fx.Invoke(scripts.Migration),
		fx.Invoke(middlewares.AddAuthenticate),
		fx.Invoke(middlewares.AddAuthorizate),
		fx.Invoke(routes.Register),
		fx.Invoke(config.OpenCache),
		fx.Invoke(server.Run),
	)

	app.Run()
	if err := app.Err(); err != nil {
		catcher.LogFatal(err)
	}
}
