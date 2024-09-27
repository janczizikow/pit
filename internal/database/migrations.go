package database

import (
	"errors"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// Run migrations runs migrations for the database.
// It uses the `MIGRATIONS_PATH` environment variable to specify the path to the migrations.
// If the variable is not set, it defaults to `file:///migrations`.
// The `ErrNoChange` error is ignored, as it means that there are no migrations to run.
func RunMigrations(dataSourceName string) error {
	path := os.Getenv("MIGRATIONS_PATH")
	if path == "" {
		path = "file:///migrations"
	}
	m, err := migrate.New(path, dataSourceName)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}
	return nil
}
