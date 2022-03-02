package config

import (
	"errors"

	"github.com/spf13/viper"
	"gitlab.com/h3mmy/bloopyboi/bot/internal/log"
)

var (
	logger = log.New()
)

// Bot Config
type Config struct {
	BotToken string
	BotName  string
	AppId    int64
	Features []FeatureConfig
}

// Feature Specific Config
type FeatureConfig struct {
	Name    string
	Enabled bool
	Data    map[string]string
}

// GetConfig returns bloopyboi configuration
func GetConfig() (*Config, error) {
	var c Config
	err := viper.Unmarshal(&c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (myConfig *Config) GetFeatureConfig(name string) (FeatureConfig, error) {
	for _, feat := range myConfig.Features {
		if feat.Name == name {
			return feat, nil
		}
	}
	logger.Error("Could not find config for feature", name)
	return FeatureConfig{}, errors.New("could not find config for feature")
}
