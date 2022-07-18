package flag

import (
	"github.com/feditools/relay/internal/config"
	"github.com/spf13/cobra"
)

// Server adds all flags for running the server.
func Server(cmd *cobra.Command, values config.Values) {
	// server
	cmd.PersistentFlags().String(config.Keys.ServerExternalHostname, values.ServerExternalHostname, usage.ServerExternalHostname)
	cmd.PersistentFlags().String(config.Keys.ServerHTTPBind, values.ServerHTTPBind, usage.ServerHTTPBind)
	cmd.PersistentFlags().Bool(config.Keys.ServerMinifyHTML, values.ServerMinifyHTML, usage.ServerMinifyHTML)
	cmd.PersistentFlags().StringArray(config.Keys.ServerRoles, values.ServerRoles, usage.ServerRoles)

	// runner
	cmd.PersistentFlags().Int(config.Keys.RunnerConcurrency, values.RunnerConcurrency, usage.RunnerConcurrency)
}
