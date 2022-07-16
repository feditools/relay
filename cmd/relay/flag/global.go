package flag

import (
	"github.com/feditools/relay/internal/config"
	"github.com/spf13/cobra"
)

// Global adds flags that are common to all commands.
func Global(cmd *cobra.Command, values config.Values) {
	cmd.PersistentFlags().String(config.Keys.ConfigPath, values.ConfigPath, usage.ConfigPath)
	cmd.PersistentFlags().String(config.Keys.LogLevel, values.LogLevel, usage.LogLevel)

	// application
	cmd.PersistentFlags().Int(config.Keys.ActorKeySize, values.ActorKeySize, usage.ActorKeySize)
	cmd.PersistentFlags().String(config.Keys.ApplicationName, values.ApplicationName, usage.ApplicationName)
	cmd.PersistentFlags().Int(config.Keys.CachedActivityLimit, values.CachedActivityLimit, usage.CachedActivityLimit)
	cmd.PersistentFlags().Int(config.Keys.CachedActorLimit, values.CachedActorLimit, usage.CachedActorLimit)
	cmd.PersistentFlags().Int(config.Keys.CachedDigestLimit, values.CachedDigestLimit, usage.CachedDigestLimit)

	// database
	cmd.PersistentFlags().String(config.Keys.DbType, values.DbType, usage.DbType)
	cmd.PersistentFlags().String(config.Keys.DbAddress, values.DbAddress, usage.DbAddress)
	cmd.PersistentFlags().Int(config.Keys.DbPort, values.DbPort, usage.DbPort)
	cmd.PersistentFlags().String(config.Keys.DbUser, values.DbUser, usage.DbUser)
	cmd.PersistentFlags().String(config.Keys.DbPassword, values.DbPassword, usage.DbPassword)
	cmd.PersistentFlags().String(config.Keys.DbDatabase, values.DbDatabase, usage.DbDatabase)
	cmd.PersistentFlags().String(config.Keys.DbTLSMode, values.DbTLSMode, usage.DbTLSMode)
	cmd.PersistentFlags().String(config.Keys.DbTLSCACert, values.DbTLSCACert, usage.DbTLSCACert)
	cmd.PersistentFlags().String(config.Keys.DbEncryptionKey, values.DbEncryptionKey, usage.DbEncryptionKey)

	// metrics
	cmd.PersistentFlags().String(config.Keys.MetricsStatsDAddress, values.MetricsStatsDAddress, usage.MetricsStatsDAddress)
	cmd.PersistentFlags().String(config.Keys.MetricsStatsDPrefix, values.MetricsStatsDPrefix, usage.MetricsStatsDPrefix)
}
