package flag

import (
	"github.com/feditools/relay/internal/config"
	"github.com/spf13/cobra"
)

// Database adds all flags for running the database command.
func Database(cmd *cobra.Command, values config.Values) {
}

// DatabaseMigrate adds all flags for running the database migrate command.
func DatabaseMigrate(cmd *cobra.Command, values config.Values) {
	Database(cmd, values)
}
