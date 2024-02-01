package services

import (
	overseerr_go "github.com/devopsarr/overseerr-go/overseerr"
	"github.com/h3mmy/bloopyboi/bot/internal/log"
	"github.com/h3mmy/bloopyboi/bot/internal/models"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type MediaService struct {
	bloopyMeta      models.BloopyMeta
	logger          *zap.Logger
	overseerrClient *overseerr_go.APIClient
	bookService     *BookService
}

func NewMediaService(overseerrGen *OverseerrClientGenerator, bookService *BookService) *MediaService {
	lgr := log.NewZapLogger().With(
		zapcore.Field{Type: zapcore.StringType, Key: ServiceLoggerFieldKey, String: "media_service"},
	)
	return &MediaService{
		bloopyMeta:      models.NewBloopyMeta(),
		logger:          lgr,
		overseerrClient: overseerrGen.generateClient(),
		bookService:     bookService,
	}
}
