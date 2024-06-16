package services

import (
	"github.com/h3mmy/bloopyboi/bot/arr"
	"github.com/h3mmy/bloopyboi/internal/models"
	"github.com/h3mmy/bloopyboi/pkg/config"
	log "github.com/h3mmy/bloopyboi/pkg/logs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ArrService struct {
	meta           models.BloopyMeta
	logger         *zap.Logger
	clientRegistry *arr.ArrClientRegistry
}

func NewArrService(cfg *config.AppConfig) *ArrService {
	meta := models.NewBloopyMeta()
	lgr := log.NewZapLogger().With(
		zapcore.Field{Type: zapcore.StringType, Key: ServiceLoggerFieldKey, String: "arr_service"},
	)

	registry := arr.NewArrClientRegistry("arr_service")
	for _, arrCfg := range *cfg.Arrs {
		err := registry.AddClient(&arrCfg)
		if err != nil {
			lgr.Error("failed to add client to registry. Skipping config entry")
		}
	}

	return &ArrService{
		meta:           meta,
		clientRegistry: registry,
		logger:         lgr,
	}
}
