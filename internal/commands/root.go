package commands

import (
	"context"
	"errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tellmeac/go-template/internal/core"
	"github.com/tellmeac/go-template/pkg/commands"
	"github.com/tellmeac/go-template/pkg/config"
)

// Options stores common values for a command.
type Options struct {
	ConfigPath string
}

var options Options

// AddDefaultFlagSet provides default flag set to command.
func AddDefaultFlagSet(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVarP(&options.ConfigPath, "config", "c",
		"./config.yaml", "Configuration file path")
}

// Creator represents a builder function for any application command.
type Creator func(
	ctx context.Context,
	cmd *cobra.Command,
	configLoader *viper.Viper,
) (*core.Config, commands.Command, error)

// RunCommand returns ready to use cobra run function.
func RunCommand(creator Creator) func(cmd *cobra.Command, args []string) error {
	confLoader := config.PrepareLoader(
		config.WithConfigPath(options.ConfigPath),
	)

	return func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		_, command, err := creator(ctx, cmd, confLoader)
		if err != nil {
			return err
		}
		if command == nil {
			return errors.New("creator returned nil command")
		}

		err = command.Init(ctx)
		if err != nil {
			return err
		}

		return command.Run()
	}
}
