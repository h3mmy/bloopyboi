package providers

import (
	"context"

	"github.com/h3mmy/bloopyboi/bot/services"
	"github.com/h3mmy/bloopyboi/pkg/config"
	"go.uber.org/zap"
)

func NewDiscordServiceWithConfig(cfg *config.DiscordConfig) (*services.DiscordService, error) {
	imageAnalyzer, err := NewGoogleVisionAnalyzer(context.Background())
	if err != nil {
		logger.Error("failed to create image analyzer", zap.Error(err))
	}
	imageAnalysisSvc := services.NewImageAnalyzerService(imageAnalyzer)

	dsvc := services.NewDiscordService().
		WithConfig(cfg).
		WithImageAnalyzer(*imageAnalysisSvc)
	err = dsvc.RefreshDBConnection()
	if err != nil {
		logger.Warn("encountered error refreshing db connection. persistence may not be available", zap.Error(err))
	}
	return dsvc, nil
}
