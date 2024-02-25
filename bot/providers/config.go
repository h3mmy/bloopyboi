package providers

import (
	"fmt"
	"strings"

	"github.com/h3mmy/bloopyboi/bot/internal/config"
	"go.uber.org/zap/zapcore"
)


func GetDiscordConfig() *config.DiscordConfig {
	AppConfig := config.GetConfig()
	return AppConfig.DiscordConfig
}

// Retrieves logLevel if set
func GetLogLevel() string {
	AppConfig := config.GetConfig()
	return strings.ToLower(AppConfig.LogLevel)
}

// Returns
func GetFeatures() map[string]config.FeatureConfig {
	AppConfig := config.GetConfig()
	logger.Debug(fmt.Sprintf("Got FeatureMap %v", AppConfig.FeatureMap), zapcore.Field{Key: "package", Type: zapcore.StringType, String: "providers"})
	return AppConfig.FeatureMap
}

// Checks FeatureConfigs for key
func IsFeaturedConfigured(key string) bool {
	AppConfig := config.GetConfig()
	fCfg, ok := AppConfig.FeatureMap[key]
	if !ok {
		return false
	}
	return fCfg.Enabled
}
