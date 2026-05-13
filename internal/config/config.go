package config

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

func Init(cfgFile string) {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, _ := os.UserHomeDir()
		viper.AddConfigPath(home)
		viper.SetConfigName(".statsync")
		viper.SetConfigType("yaml")
	}

	viper.SetEnvPrefix("STATSYNC")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	viper.SetDefault("valkey.addr", "localhost:6379")

	_ = viper.ReadInConfig()
}

func ValkeyAddr() string {
	return viper.GetString("valkey.addr")
}
