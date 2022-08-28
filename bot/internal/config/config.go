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
type BotConfig struct {
	BotToken	string			`mapstructure:"botToken"`
	BotName		string			`mapstructure:"botName"`
	AppId		int64			`mapstructure:"appId"`
	Features	[]FeatureConfig
	LogLevel	string			`mapstructure:"logLevel"`
	DBConfig	BloopyDBConfig	`mapstructure:"db"`
}

// Feature Specific Config
type FeatureConfig struct {
	Name		string				`mapstructure:"name"`
	Enabled		bool				`mapstructure:"enabled"`
	Data		map[string]string
}

type BloopyDBConfig struct {
	Name			string		`mapstructure:"name"`
	Type			string		`mapstructure:"type"`
	Host			string		`mapstructure:"host"`
	Port			string		`mapstructure:"port"`
	User			string		`mapstructure:"user"`
	Password		string		`mapstructure:"password"`
}


// GetConfig returns bloopyboi configuration
func GetConfig() (*BotConfig, error) {
	var c BotConfig
	err := viper.Unmarshal(&c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (myConfig *BotConfig) GetFeatureConfig(name string) (FeatureConfig, error) {
	for _, feat := range myConfig.Features {
		if feat.Name == name {
			return feat, nil
		}
	}
	logger.Error("Could not find config for feature", name)
	return FeatureConfig{}, errors.New("could not find config for feature")
}

func (myConfig *BotConfig) GetConfiguredFeatureNames() []string {
	var names []string
	for _, feat := range myConfig.Features {
		names = append(names, feat.Name)
	}
	return names
}
