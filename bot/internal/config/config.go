package config

import (
	"errors"

	"github.com/spf13/viper"
	"gitlab.com/h3mmy/bloopyboi/bot/internal/log"
	"go.uber.org/zap/zapcore"
)

var (
	logger = log.NewZapLogger().With(
		zapcore.Field{
			Key:    "group",
			Type:   zapcore.StringType,
			String: "config",
		}).Sugar()
)

// Bot Config
type BotConfig struct {
	BotToken   string `mapstructure:"botToken"`
	BotName    string `mapstructure:"botName"`
	AppId      int64  `mapstructure:"appId"`
	features   []FeatureConfig
	LogLevel   string         `mapstructure:"logLevel"`
	DBConfig   BloopyDBConfig `mapstructure:"db"`
	FeatureMap map[string]FeatureConfig
}

// Feature Specific Config
type FeatureConfig struct {
	Name    string `mapstructure:"name"`
	Enabled bool   `mapstructure:"enabled"`
	Data    map[string]string
}

type BloopyDBConfig struct {
	Name     string `mapstructure:"name"`
	Type     string `mapstructure:"type"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

// GetConfig returns bloopyboi configuration
func GetConfig() (*BotConfig, error) {
	var c BotConfig
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

// Gets FeatureConfig for Key if exists
func (myConfig *BotConfig) GetFeatureConfig(name string) (FeatureConfig, error) {
	feat, ok := myConfig.FeatureMap[name]
	if ok {
		return feat, nil
	}
	logger.Error("Could not find config for feature", name)
	return FeatureConfig{}, errors.New("could not find config for feature")
}

// Builds FeatureMap. Faster to reference than array
func (myConfig *BotConfig) buildFeatureMap() error {
	featMap := make(map[string]FeatureConfig)
	for _, feat := range myConfig.features {
		featMap[feat.Name] = feat
	}
	myConfig.FeatureMap = featMap
	return nil
}

// deprecated. Refactor to use FeatureMap Keys
func (myConfig *BotConfig) GetConfiguredFeatureNames() []string {
	var names []string
	for _, feat := range myConfig.features {
		names = append(names, feat.Name)
	}
	return names
}
