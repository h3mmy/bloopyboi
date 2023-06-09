package providers

import (
	"fmt"
	"strings"

	"gitlab.com/h3mmy/bloopyboi/bot/internal/config"
	"go.uber.org/zap/zapcore"
)

// Retrieves Bot Token
func GetBotToken() string {
	botConfig, err := config.GetConfig()
	if err != nil {
		logger.Sugar().Error("Error Loading Config", err)
	}
	return botConfig.BotToken
}

func GetBotName() string {
	botConfig, err := config.GetConfig()
	if err != nil {
		logger.Sugar().Error("Error Loading Config", err)
	}
	return botConfig.BotToken
}

// Retrieves logLevel if set
func GetLogLevel() string {
	botConfig, err := config.GetConfig()
	if err != nil {
		logger.Sugar().Error("Error Loading Config", err)
	}
	return strings.ToLower(botConfig.LogLevel)
}

// Returns
func GetFeatures() map[string]config.FeatureConfig {
	botConfig, err := config.GetConfig()
	if err != nil {
		logger.Sugar().Error("Error Loading Config", err)
		return nil
	}
	logger.Debug(fmt.Sprintf("Got FeatureMap %v", botConfig.FeatureMap), zapcore.Field{Key: "package", Type: zapcore.StringType, String: "providers"})
	return botConfig.FeatureMap
}

// Checks FeatureConfigs for key
func IsFeaturedConfigured(key string) bool {
	botConfig, err := config.GetConfig()
	if err != nil {
		logger.Sugar().Error("Error Loading Config", err)
		return false
	}
	fCfg, ok := botConfig.FeatureMap[key]
	if !ok {
		return false
	}
	return fCfg.Enabled
}
