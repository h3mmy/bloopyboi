package config

import (
	"errors"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var currentConfig *AppConfigLoader

type ConfigLoader interface {
	GetConfig() interface{}
	UpdateConfig()
	GetLastUpdated() time.Time
	GetRevision() int
}

// initConfig reads in config file and ENV variables if set.
// It also registers a filewatch to refresh the config when a change is detected
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
	viper.OnConfigChange(func(e fsnotify.Event) {
		logger.Info("Config file changed:", e.Name)
		RefreshAppConfig()
	})
	viper.WatchConfig()
}

// GetConfig returns bloopyboi configuration from the current viper instance
func buildConfig() (*AppConfig, error) {
	var c AppConfig
	err := viper.Unmarshal(&c)
	if err != nil {
		logger.Error("error unmarshalling config", zap.Error(err))
		return nil, err
	}
	err = c.buildFeatureMap()
	if err != nil {
		logger.Error("error building feature map", zap.Error(err))
		return nil, err
	}
	logger.Debug("config loaded", zap.Any("config", c))
	return &c, nil
}

// GetAppConfigLoader returns the current AppConfigLoader instance.
// If it does not exist, it will create a new AppConfigLoader instance.
func GetAppConfigLoader() *AppConfigLoader {
	if currentConfig != nil {
		return currentConfig
	}
	initConfig()
	AppConfig, _ := buildConfig()
	currentConfig = NewAppConfigLoader(AppConfig)
	return currentConfig
}

// GetConfig returns the current bloopyboi configuration from the AppConfigLoader
func GetConfig() *AppConfig {
	return GetAppConfigLoader().GetConfig()
}

func RefreshAppConfig() {
	logger.Debug("refreshing App Config")
	cfg, err := buildConfig()
	if err != nil {
		logger.Error("error updating config", zap.Error(err))
	} else {
		GetAppConfigLoader().UpdateConfig(cfg)
	}
}
