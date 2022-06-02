package repository

import (
	"context"
	"database/sql"
	"github.com/ArseniySavin/auth-small-server/config"
	"github.com/ArseniySavin/auth-small-server/internal/pkg/tools"
	catcher "github.com/ArseniySavin/catcher/pkg"
	"github.com/jmoiron/sqlx"
	"go.uber.org/fx"

	_ "github.com/lib/pq"
)

// CreateDbConnection - It is Connection for DB
func CreateDbConnection(lc fx.Lifecycle, cfg *config.Config) *sql.DB {
	DB := cfg.DB

	dsn := tools.GetDsn(DB.Login, DB.Pass, DB.Host, DB.Name, DB.SslMode, DB.Port)

	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		catcher.LogFatal(err)

	}

	lc.Append(fx.Hook{
		OnStop: func(context.Context) error {
			db.Close()
			return nil
		},
	})

	return db.DB
}
