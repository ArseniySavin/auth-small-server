package scripts

import (
	"database/sql"
	catcher "github.com/ArseniySavin/catcher/pkg"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migration(db *sql.DB) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		catcher.LogFatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://scripts/db",
		"postgres", driver)
	err = m.Up()
	if err == migrate.ErrNoChange {
		return
	}
	if err != nil {
		catcher.LogFatal(err)
	}

	close(m)
}

func close(m *migrate.Migrate) {
	err, dbErr := m.Close()
	switch {
	case err != nil:
		catcher.LogFatal(err)
	case dbErr != nil:
		catcher.LogFatal(dbErr)
	}
}
