package runserver

import (
	"context"
	"github.com/heetch/confita"
	"github.com/spf13/cobra"
	"github.com/tellmeac/go-template/internal/commands"
	"github.com/tellmeac/go-template/internal/core"
	"github.com/tellmeac/go-template/internal/server"
)

func CommandCreator(ctx context.Context, _ *cobra.Command, loader *confita.Loader,
) (*core.Config, *Command, error) {
	conf, err := core.ParseConfig(ctx, loader)
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
	app := server.New()
}
