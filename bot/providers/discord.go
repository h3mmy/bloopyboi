package providers

import (
	"github.com/h3mmy/bloopyboi/bot/internal/config"
	"github.com/h3mmy/bloopyboi/bot/services"
	"go.uber.org/zap"
)

func NewDiscordServiceWithConfig(cfg *config.DiscordConfig) (*services.DiscordService, error) {
	dsvc := services.NewDiscordService().WithConfig(cfg)
	err := dsvc.RefreshDBConnection()
	logger.Warn("encountered error refreshing db connection. persistence may not be available", zap.Error(err))
	return dsvc, nil
}
