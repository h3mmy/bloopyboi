package services

import (
	"context"
	"fmt"

	"github.com/h3mmy/bloopyboi/bot/internal/database"
	"github.com/h3mmy/bloopyboi/bot/internal/log"
	"github.com/h3mmy/bloopyboi/bot/internal/models"
	"github.com/h3mmy/bloopyboi/ent"
	"github.com/h3mmy/bloopyboi/ent/discorduser"
	"github.com/h3mmy/bloopyboi/ent/mediarequest"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type MediaService struct {
	bloopyMeta       models.BloopyMeta
	logger           *zap.Logger
	overseerrService *OverseerrService
	bookService      *BookService
	arrService *ArrService
	db               *ent.Client
	dbEnabled        bool
}

func NewMediaService() *MediaService {
	lgr := log.NewZapLogger().With(
		zapcore.Field{Type: zapcore.StringType, Key: ServiceLoggerFieldKey, String: "media_service"},
	)
	return &MediaService{
		bloopyMeta:       models.NewBloopyMeta(),
		logger:           lgr,
		overseerrService: nil,
		bookService:      nil,
		dbEnabled:        false,
		db:               nil,
	}
}

func (s *MediaService) WithBookService(bsvc *BookService) {
	s.bookService = bsvc
}

func (s *MediaService) WithOverseerrService(osvc *OverseerrService) {
	s.overseerrService = osvc
}

func (s *MediaService) WithArrService(svc *ArrService) {
	s.arrService = svc
}

func (s *MediaService) RefreshDBConnection() error {
	if s.dbEnabled {
		s.db.Close()
	}
	dbEnabled := true
	dbClient, err := database.Open()
	if err != nil {
		s.logger.Error("failed to open database", zap.Error(err))
		dbEnabled = false
	}
	s.db = dbClient
	s.dbEnabled = dbEnabled

	return err
}

func (s *MediaService) GetMediaRequestsForUser(ctx context.Context, discordUserId int) ([]*ent.MediaRequest, error) {
	// TODO: Create composite request response??

	if s.dbEnabled {
		requests, err := s.db.MediaRequest.
			Query().
			WithBook().
			Where(
				mediarequest.HasDiscordUsersWith(
					discorduser.DiscordidEQ(fmt.Sprint(discordUserId)),
				),
			).
			All(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to get book requests for discord user %d: %w", discordUserId, err)
		}
		return requests, nil
	}
	s.logger.Warn("database not enabled")
	return nil, nil
}

