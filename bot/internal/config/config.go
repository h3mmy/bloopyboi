package config

import (
	"github.com/spf13/viper"
	"gitlab.com/h3mmy/bloopyboi/bot/internal/log"
)

var (
	logger = log.New()
)

// Bot Config
type Config struct {
	BotToken string
	BotName string
	AppId int64
	Features []FeatureConfig
}

// Feature Specific Config
type FeatureConfig struct {
	Name string
	Enabled bool
	Data map[string]string
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
