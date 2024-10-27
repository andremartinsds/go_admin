package configs

import (
	"github.com/spf13/viper"
)

func LoadEnvConfig() error {
	viper.SetConfigName(".env")
	// viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AddConfigPath("../..")
	err := viper.ReadInConfig()

	if err != nil {
		return err
	}
	return nil
}
