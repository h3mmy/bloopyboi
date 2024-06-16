package providers

import (
	"github.com/h3mmy/bloopyboi/bot/services"
	"github.com/h3mmy/bloopyboi/pkg/config"
	"go.uber.org/zap"
)

func NewDiscordServiceWithConfig(cfg *config.DiscordConfig) (*services.DiscordService, error) {
	dsvc := services.NewDiscordService().WithConfig(cfg)
	err := dsvc.RefreshDBConnection()
	if err != nil {
		logger.Warn("encountered error refreshing db connection. persistence may not be available", zap.Error(err))
	}
	return dsvc, nil
}
