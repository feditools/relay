package flag

import (
	"github.com/feditools/relay/internal/config"
	"github.com/spf13/cobra"
)

// Redis adds flags that are common to redis.
func Redis(cmd *cobra.Command, values config.Values) {
	cmd.PersistentFlags().String(config.Keys.RedisAddress, values.RedisAddress, usage.RedisAddress)
	cmd.PersistentFlags().Int(config.Keys.RedisDB, values.RedisDB, usage.RedisDB)
	cmd.PersistentFlags().String(config.Keys.RedisPassword, values.RedisPassword, usage.RedisPassword)
}
