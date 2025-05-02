package db

import (
	"database/sql"
	"embed"
	"fmt"

	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func Migrate(db *sql.DB) (err error) {
	goose.SetBaseFS(embedMigrations)

	if err = goose.SetDialect(string(goose.DialectMySQL)); err != nil {
		return fmt.Errorf("set migrations dialect as postgres: %w", err)
	}

	if err = goose.Up(db, "migrations"); err != nil {
		return fmt.Errorf("up migrations: %w", err)
	}

	return nil
}
