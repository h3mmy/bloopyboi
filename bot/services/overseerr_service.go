package services

import (
	overseerr_go "github.com/devopsarr/overseerr-go/overseerr"
	"github.com/h3mmy/bloopyboi/internal/models"
	log "github.com/h3mmy/bloopyboi/pkg/logs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// OverseerrService is a service that interacts with the Overseerr API.
type OverseerrService struct {
	bloopyMeta      models.BloopyMeta
	logger          *zap.Logger
	overseerrClient *overseerr_go.APIClient
}

// NewOverseerrService creates a new OverseerrService.
func NewOverseerrService(clientgen *OverseerrClientGenerator) *OverseerrService {
	lgr := log.NewZapLogger().With(
		zapcore.Field{Type: zapcore.StringType, Key: ServiceLoggerFieldKey, String: "overseerr_service"},
	)

	return &OverseerrService{
		bloopyMeta:      models.NewBloopyMeta(),
		logger:          lgr,
		overseerrClient: clientgen.generateClient(),
	}
}
