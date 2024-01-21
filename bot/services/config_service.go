package services

import (
	"fmt"

	"github.com/h3mmy/bloopyboi/bot/internal/config"
	"github.com/h3mmy/bloopyboi/bot/internal/log"
	"github.com/h3mmy/bloopyboi/bot/internal/models"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ConfigService struct {
	bloopyMeta   models.BloopyMeta
	logger       *zap.Logger
	configLoader config.ConfigLoader
}

func NewConfigService(configLoader config.ConfigLoader) *ConfigService {
	lgr := log.NewZapLogger().With(
		zapcore.Field{Type: zapcore.StringType, Key: ServiceLoggerFieldKey, String: "config_service"},
	)

	return &ConfigService{
		bloopyMeta:   models.NewBloopyMeta(),
		logger:       lgr,
		configLoader: configLoader,
	}
}

func (s *ConfigService) GetConfig() (*config.AppConfig, error) {
	cfg := s.configLoader.GetConfig()
	if appCfg, ok := cfg.(*config.AppConfig); ok {
		return appCfg, nil
	}
	return nil, fmt.Errorf("config does not conform to AppConfig")
}
