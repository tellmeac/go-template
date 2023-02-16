package main

import (
	"github.com/spf13/cobra"
	"github.com/tellmeac/go-template/internal/commands"
	"github.com/tellmeac/go-template/internal/commands/runserver"
	"log"
)

func main() {
	cmd := &cobra.Command{
		Use:   "runserver",
		Short: "start server",
		RunE:  commands.RunCommand(runserver.CommandCreator),
	}

	commands.AddDefaultFlagSet(cmd)

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
