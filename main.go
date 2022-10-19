package main

import (
	"github.com/rs/zerolog/log"
	"github.com/tellmeac/go-template/internal/adapters"
	"github.com/tellmeac/go-template/internal/app"
	"github.com/tellmeac/go-template/internal/config"
	"github.com/tellmeac/go-template/internal/pkg/server"
	v1 "github.com/tellmeac/go-template/internal/ports/v1"
)

func main() {
	cfg := config.Get()

	client, err := adapters.NewDatabaseClient(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to setup database connection")
	}

	dao := adapters.NewDAO(client)

	a := app.New(dao)

	s := server.New(cfg)
	s.Attach(v1.New(a))

	if err := s.ListenAndServe(); err != nil {
		log.Fatal().Err(err).Msg("Failed to run http server")
	}
}
