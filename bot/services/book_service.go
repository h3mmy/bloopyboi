package services

import (
	"github.com/h3mmy/bloopyboi/bot/internal/log"
	"github.com/h3mmy/bloopyboi/bot/internal/models"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	books "google.golang.org/api/books/v1"
)

type BookService struct {
	bloopyMeta *models.BloopyMeta
	logger     *zap.Logger
	svc        *books.Service
}

func NewBookService() *BookService {
	lgr := log.NewZapLogger().With(
		zapcore.Field{Type: zapcore.StringType, Key: ServiceLoggerFieldKey, String: "book_service"},
	)
	bookSvc := books.New()

	return &BookService{
		svc:        bookSvc,
		logger:     lgr,
		bloopyMeta: models.NewBloopyMeta(),
	}
}
