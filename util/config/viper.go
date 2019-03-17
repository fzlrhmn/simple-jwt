package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var (
	// Instance is a viper instance that store application configs
	Instance *viper.Viper
)

func init() {
	if configFileFromEnv := os.Getenv("UMS_CONFIG_FILE"); configFileFromEnv != "" {
		viper.SetConfigFile(fmt.Sprintf("%s", configFileFromEnv))
	} else {
		env := os.Getenv("USR_ENV")
		if env == "" {
			env = "development"
		}

		viper.SetConfigFile(GetConfigFile(fmt.Sprintf("%s.toml", env)))
	}

	viper.BindEnv("app.env", "USR_ENV")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	Instance = viper.GetViper()
}
