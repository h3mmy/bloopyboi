package providers

import (
	"strings"
	"gitlab.com/h3mmy/bloopyboi/bot/internal/config"
)

// Retrieves Bot Token
func GetBotToken() string {
	botConfig, err := config.GetConfig()
	if err != nil {
		logger.Error("Error Loading Config", err)
	}
	return botConfig.BotToken
}

// Retrieves logLevel if set
func GetLogLevel() string {
	botConfig, err := config.GetConfig()
	if err != nil {
		logger.Error("Error Loading Config", err)
	}
	return strings.ToLower(botConfig.LogLevel)
}
