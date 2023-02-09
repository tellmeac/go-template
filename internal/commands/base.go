package commands

import (
	"context"
	"github.com/tellmeac/go-template/internal/core"
	"github.com/tellmeac/go-template/internal/storage/db"
	"github.com/tellmeac/go-template/pkg/commands"
)

// BaseCommand содержит в себе общую логику инициализации команд.
// TODO: вынести db.Storage, добавить клиенты метрик, трейсинга и чего-нибудь еще.
type BaseCommand struct {
	commands.BaseCommand

	Config  *core.Config
	Storage *db.Storage // Should be in StorageCommand
}

func (c *BaseCommand) Init(ctx context.Context) error {
	var err error

	c.Storage, err = db.NewStorage(ctx, c.Config.Storage)
	if err != nil {
		return err
	}

	return nil
}

func (c *BaseCommand) NewRepository() *Repository {
	return &Repository{
		Config:  c.Config,
		Storage: c.Storage,
	}
}
