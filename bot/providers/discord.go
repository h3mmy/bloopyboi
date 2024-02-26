package providers

import (
	"github.com/h3mmy/bloopyboi/bot/internal/config"
	"github.com/h3mmy/bloopyboi/bot/services"
)

func NewDiscordServiceWithConfig(cfg *config.DiscordConfig) (*services.DiscordService, error) {
	dsvc := services.NewDiscordService().WithConfig(cfg)
	err := dsvc.RefreshDBConnection()
	return dsvc, err
}
