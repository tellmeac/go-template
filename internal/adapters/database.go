package adapters

import (
	"database/sql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/tellmeac/go-template/internal/adapters/ent"
	"github.com/tellmeac/go-template/internal/config"
	// Required to connect to postgres database
	_ "github.com/jackc/pgx/v5/stdlib"
)

// NewDatabaseClient returns new ent client.
func NewDatabaseClient(cfg config.Config) (*ent.Client, error) {
	db, err := sql.Open("pgx", cfg.DatabaseAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	drv := entsql.OpenDB(dialect.Postgres, db)

	client := ent.NewClient(ent.Driver(drv))

	if cfg.Debug {
		client = client.Debug()
	}

	return client, nil
}
