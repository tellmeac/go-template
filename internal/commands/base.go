package commands

import (
	"context"
	"github.com/spf13/cobra"
	"github.com/tellmeac/go-template/internal/core"
	"github.com/tellmeac/go-template/internal/storage/db"
	"github.com/tellmeac/go-template/pkg/commands"
	"go.uber.org/zap"
)

const (
	ConfigEnv = "EXAMPLE_CONFIG"
)

type Options struct {
	ConfigPath string
	LogLevel   string
}

var options Options

func AddDefaultFlagSet(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVarP(&options.ConfigPath, "config", "c",
		"./config.yaml", "Config file path")
	cmd.PersistentFlags().StringVarP(&options.LogLevel, "level", "l",
		zap.DebugLevel.String(), "Log level")
}

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
