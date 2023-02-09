package runserver

import (
	"context"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tellmeac/go-template/internal/commands"
	"github.com/tellmeac/go-template/internal/core"
	"github.com/tellmeac/go-template/internal/server"
	xcommands "github.com/tellmeac/go-template/pkg/commands"
)

func CommandCreator(_ context.Context, _ *cobra.Command, configLoader *viper.Viper,
) (*core.Config, xcommands.Command, error) {
	conf, err := core.ParseConfig(configLoader)
	if err != nil {
		return nil, nil, err
	}

	command := newRunServerCommand(conf)
	return conf, &command, nil
}

func newRunServerCommand(conf *core.Config) Command {
	cmd := Command{
		BaseCommand: commands.BaseCommand{},
	}
	cmd.BaseCommand.Config = conf

	return cmd
}

type Command struct {
	commands.BaseCommand
}

func (c *Command) Run() error {
	app := server.New(c.NewRepository())
	// TODO: go func with ctx.Done handle to stop app

	return app.Start(c.Context())
}
