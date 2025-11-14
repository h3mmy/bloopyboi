package services

import (
	"context"
	"fmt"

	"github.com/h3mmy/bloopyboi/ent"
	"github.com/h3mmy/bloopyboi/ent/discorduser"
	"github.com/h3mmy/bloopyboi/ent/mediarequest"
	"github.com/h3mmy/bloopyboi/internal/models"
	"github.com/h3mmy/bloopyboi/pkg/database"
	log "github.com/h3mmy/bloopyboi/pkg/logs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// MediaService is a service that interacts with media-related services.
type MediaService struct {
	bloopyMeta       models.BloopyMeta
	logger           *zap.Logger
	overseerrService *OverseerrService
	bookService      *BookService
	arrService       *ArrService
	db               *ent.Client
	dbEnabled        bool
}

// NewMediaService creates a new MediaService.
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

// WithBookService sets the BookService for the MediaService.
func (s *MediaService) WithBookService(bsvc *BookService) {
	s.bookService = bsvc
}

// WithOverseerrService sets the OverseerrService for the MediaService.
func (s *MediaService) WithOverseerrService(osvc *OverseerrService) {
	s.overseerrService = osvc
}

// WithArrService sets the ArrService for the MediaService.
func (s *MediaService) WithArrService(svc *ArrService) {
	s.arrService = svc
}

// RefreshDBConnection refreshes the database connection.
func (s *MediaService) RefreshDBConnection() error {
	if s.dbEnabled {
		if err := s.db.Close(); err != nil {
			s.logger.Error("failed to close database connection", zap.Error(err))
		}
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

// GetMediaRequestsForUser gets the media requests for a user.
// TODO: Create composite request response??
func (s *MediaService) GetMediaRequestsForUser(ctx context.Context, discordUserId int) ([]*ent.MediaRequest, error) {
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
