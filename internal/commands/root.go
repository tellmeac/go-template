package commands

import (
	"context"
	"errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tellmeac/go-template/internal/core"
	"github.com/tellmeac/go-template/pkg/commands"
	"github.com/tellmeac/go-template/pkg/config"
	"go.uber.org/zap"
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

type Creator func(
	ctx context.Context,
	cmd *cobra.Command,
	config *viper.Viper,
) (*core.Config, commands.Command, error)

func RunCommand(creator Creator) func(cmd *cobra.Command, args []string) error {
	ctx := context.Background()
	f := makeInitFunc(ctx, creator)

	return func(cmd *cobra.Command, args []string) error {
		var err error

		_, command, err := f(cmd)
		if err != nil {
			return err
		}
		if command == nil {
			return errors.New("CommandCreator returned nil command")
		}

		// TODO: setup logging
		//logLevel, err := log.ParseLevel(options.LogLevel)
		//if err != nil {
		//	return err
		//}

		// TODO: Wrapper that uses commands.Command interface to perform it like commands.Main(ctx, logLevel, ...)
		err = command.Init(ctx)
		if err != nil {
			return err
		}

		return command.Run()
	}
}

func makeInitFunc(ctx context.Context, creator Creator) func(*cobra.Command) (*core.Config, commands.Command, error) {
	return func(cmd *cobra.Command) (*core.Config, commands.Command, error) {
		configLoader, err := config.PrepareLoader(
			config.WithConfigPath(options.ConfigPath),
		)
		if err != nil {
			return nil, nil, err
		}

		conf, command, err := creator(ctx, cmd, configLoader)
		if err != nil {
			return nil, nil, err
		}

		return conf, command, nil
	}
}
