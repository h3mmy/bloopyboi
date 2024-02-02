package services

import (
	"time"

	"github.com/h3mmy/bloopyboi/bot/internal/log"
	"github.com/h3mmy/bloopyboi/bot/internal/models"
	pkgmodels "github.com/h3mmy/bloopyboi/pkg/models"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var lineupImageURI = "https://www.blissfestfestival.org/wp-content/uploads/2023/04/Bliss23_LineUpIG-2-2048x2048.jpg"
var lineupSrcUrl = "https://www.blissfestfestival.org/wp-content/uploads/2023/04/Bliss23_poster_full-8F-3-scaled-e1683260548650.jpg"

var apiPrefix = "/wp-json/wp/v2"

type BlissfestService struct {
	bloopymeta models.BloopyMeta
	config     pkgmodels.BlissfestConfig
	logger     *zap.Logger
}

func NewBlissfestService(config pkgmodels.BlissfestConfig) *BlissfestService {
	lgr := log.NewZapLogger().With(
		zapcore.Field{Type: zapcore.StringType, Key: ServiceLoggerFieldKey, String: "blissfest_service"},
	)

	return &BlissfestService{
		bloopymeta: models.NewBloopyMeta(),
		config:     config,
		logger:     lgr,
	}
}

// Gets time until start of event
func (bs *BlissfestService) GetTimeUntilStart(fromDate *time.Time) time.Duration {
	comparingFrom := bs.config.Start
	if fromDate == nil {
		return time.Until(comparingFrom)
	}
	return comparingFrom.Sub(*fromDate)
}

func (bs *BlissfestService) GetStartTime() *time.Time {
	return &bs.config.Start
}

func (bs *BlissfestService) GetEndTime() *time.Time {
	return &bs.config.End
}

// Gets time until start of event
// pending https://github.com/dustin/go-humanize/pull/92
// func (bs *BlissfestService) GetHumanTimeUntilStart(fromDate *time.Time) string {

//     comparingFrom := bs.config.Start
// 	if fromDate == nil {
// 		return humanize.Time(comparingFrom)
// 	}
// 	return comparingFrom.Sub(*fromDate)
// }

// Gets time until end of event
func (bs *BlissfestService) GetTimeUntilEnd(fromDate *time.Time) time.Duration {
	comparingFrom := bs.config.End
	if fromDate == nil {
		return time.Until(comparingFrom)
	}
	return comparingFrom.Sub(*fromDate)
}

// Returns true if event inProgress
func (bs *BlissfestService) IsInProgress() bool {
	start := bs.config.Start
	end := bs.config.End
	return (start.Before(time.Now()) && end.After(time.Now()))
}

func (bs *BlissfestService) GetLineupImageURI() string {
	return lineupImageURI
}
