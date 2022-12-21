//go:build ignore

package main

import (
	"context"
	"entgo.io/ent/dialect/sql/schema"
	"flag"
	"github.com/rs/zerolog/log"
	"github.com/tellmeac/go-template/internal/infrastructure/ent/migrate"
	"os"

	atlas "ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect"

	_ "github.com/lib/pq"
)

func main() {
	ctx := context.Background()

	dir, err := atlas.NewLocalDir("migrations")
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to init atlas migration directory: %v", err)
	}

	// Migrate diff options.
	opts := []schema.MigrateOption{
		schema.WithDir(dir),
		schema.WithMigrationMode(schema.ModeReplay),
		schema.WithDialect(dialect.Postgres),
		schema.WithFormatter(atlas.DefaultFormatter),
	}

	if len(os.Args) != 2 {
		log.Fatal().Msg("Migration name is required. Use: 'go run -mod=mod migrations.go <name>'")
	}

	addr := flag.String("url", "postgres://postgres:postgres@localhost:5432/Test?sslmode=disable", "")
	flag.Parse()

	// Generate migrations using Atlas
	err = migrate.NamedDiff(ctx, *addr, os.Args[1], opts...)
	if err != nil {
		log.Fatal().Msgf("Failed generating migration file: %v", err)
	}
}
