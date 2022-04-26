package config

import (
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Init starts config collection
func Init(flags *pflag.FlagSet) error {
	viper.SetEnvPrefix("mb")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv()

	err := viper.BindPFlags(flags)
	if err != nil {
		return err
	}
	return nil
}

// ReadConfigFile reads the config file from disk if config path is sent.
func ReadConfigFile() error {
	configPath := viper.GetString(Keys.ConfigPath)
	if configPath != "" {
		viper.SetConfigFile(configPath)

		err := viper.ReadInConfig()
		if err != nil {
			return err
		}
	}
	return nil
}
