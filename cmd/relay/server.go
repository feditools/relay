package main

import (
	"github.com/feditools/relay/cmd/relay/action/server"
	"github.com/feditools/relay/cmd/relay/flag"
	"github.com/feditools/relay/internal/config"
	"github.com/spf13/cobra"
)

// serverCommands returns the 'server' subcommand
func serverCommands() *cobra.Command {
	serverCmd := &cobra.Command{
		Use:   "server",
		Short: "runs a relay server",
	}

	serverStartCmd := &cobra.Command{
		Use:   "start",
		Short: "start the feditools server",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return preRun(cmd)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(cmd.Context(), server.Start)
		},
	}

	flag.Server(serverStartCmd, config.Defaults)

	serverCmd.AddCommand(serverStartCmd)

	return serverCmd
}
