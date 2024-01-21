package services

import (
	overseerr_go "github.com/h3mmy/overseerr_go"
	"github.com/h3mmy/bloopyboi/bot/internal/log"
	"github.com/h3mmy/bloopyboi/bot/internal/models"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type OverseerrService struct {
	bloopyMeta      models.BloopyMeta
	logger          *zap.Logger
	overseerrClient *overseerr_go.APIClient
}

func NewOverseerrService(clientgen *OverseerrClientGenerator) {
	lgr := log.NewZapLogger().With(
		zapcore.Field{Type: zapcore.StringType, Key: ServiceLoggerFieldKey, String: "overseerr_service"},
	)

	return &OverseerrService{
		bloopyMeta:      models.NewBloopyMeta(),
		logger:          lgr,
		overseerrClient: clientgen.generateClient(),
	}
}
