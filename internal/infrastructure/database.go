package infrastructure

import (
	"database/sql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/tellmeac/go-template/internal/config"
	"github.com/tellmeac/go-template/internal/infrastructure/ent"

	// Required to connect to postgres database
	_ "github.com/jackc/pgx/v5/stdlib"
)

// NewDatabaseClient returns new ent client.
func NewDatabaseClient(cfg config.Config) (*ent.Client, error) {
	db, err := sql.Open("pgx", cfg.DatabaseAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	drv := ent.Driver(entsql.OpenDB(dialect.Postgres, db))

	return ent.NewClient(drv), nil
}
