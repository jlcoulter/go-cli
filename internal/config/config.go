package config

import (
	"github.com/spf13/viper"
)

func Load(cfgFile string) {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath("$HOME")
		viper.SetConfigName(".go-cli-template")
	}

	viper.AutomaticEnv()
	viper.SetEnvPrefix("APP")

	// Config file is optional — ignore error if not found
	_ = viper.ReadInConfig()
}