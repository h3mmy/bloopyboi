package config

import (
	"errors"
	"time"

	"github.com/spf13/viper"
)

var currentConfig *AppConfigLoader

type ConfigLoader interface {
	GetConfig() interface{}
	UpdateConfig()
	GetLastUpdated() time.Time
	GetRevision() int
}

func initConfig() {
	viper.SetConfigName("config")           // name of config file (without extension)
	viper.SetConfigType("yaml")             // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/config")          // path to look for the config file in
	viper.AddConfigPath("$HOME/.bloopyboi") // call multiple times to add many search paths
	viper.AddConfigPath(".")                // optionally look for config in the working directory
	viper.AutomaticEnv()
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(errors.New("Fatal error config file: " + err.Error()))
	}
}

// GetConfig returns bloopyboi configuration
func buildConfig() (*AppConfig, error) {
	initConfig()
	var c AppConfig
	err := viper.Unmarshal(&c)
	if err != nil {
		return nil, err
	}
	err = c.buildFeatureMap()
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func GetAppConfigLoader() *AppConfigLoader {
	if currentConfig != nil {
		return currentConfig
	}
	AppConfig, _ := buildConfig()
	currentConfig = NewAppConfigLoader(AppConfig)
	return currentConfig
}

func GetConfig() *AppConfig {
	return GetAppConfigLoader().GetConfig()
}
