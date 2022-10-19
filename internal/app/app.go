package app

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/tellmeac/go-template/internal/pkg/errors"
	"strings"
)

type ModifyRequest struct {
	Target   string `json:"target"`
	Template string `json:"template"`
}

type TemplatesStore interface {
	Get(ctx context.Context, target string) (string, error)
	Upsert(ctx context.Context, target string, desired string) error
}

const DefaultGreetingTemplate = "Hello, $1"

type App struct {
	templates TemplatesStore
}

func New(store TemplatesStore) *App {
	return &App{templates: store}
}

func (app App) GetGreeting(ctx context.Context, name string) string {
	template, err := app.templates.Get(ctx, name)
	switch {
	case err != nil && !errors.Is(err, errors.ErrNotFound):
		log.Error().Err(err).Msg("Failed to fetch for user specific template")
	}

	if template == "" {
		template = DefaultGreetingTemplate
	}

	return strings.ReplaceAll(template, "$1", name)
}

func (app App) ModifyGreeting(ctx context.Context, req ModifyRequest) error {
	// TODO: make template type with parse/validation and compile methods
	if !strings.Contains(req.Template, "$1") {
		return errors.ErrInvalidTemplate
	}

	err := app.templates.Upsert(ctx, req.Target, req.Template)
	if err != nil {
		return fmt.Errorf("failed to modify tempalte: %w", err)
	}
	return nil
}
