package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/h3mmy/bloopyboi/internal/models"
	log "github.com/h3mmy/bloopyboi/pkg/logs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// BlissfestService is a service that provides information about the Blissfest event.
type BlissfestService struct {
	bloopymeta models.BloopyMeta
	config     models.BlissfestConfig
	logger     *zap.Logger
}

// 2023: "https://www.blissfestfestival.org/wp-content/uploads/2023/04/Bliss23_LineUpIG-2-2048x2048.jpg"
// 2024: "https://www.blissfestfestival.org/wp-content/uploads/2024/04/Bliss24_IGAnnouncement3-2048x2048.jpg"
// 2025: "https://www.blissfestfestival.org/wp-content/uploads/2024/04/Bliss24_IGAnnouncement3-2048x2048.jpg"
var lineupImageURI = "https://blissfest.org/cdn/shop/files/Blissfest25_Lineup_HomePage5a.jpg?v=1741384591&width=3000"

// 2024 blissfest showclix "event_id": 9297272, "parent_event_id": 8615552,
// 2024 blissfest showclix venue_id = 64139

var blissfestShowclixEventID = 9297272

// 2024 blissfest logo art "https://www.blissfestfestival.org/wp-content/uploads/2024/01/Bliss_Logo_2024F3.png"

// var apiPrefix = "/wp-json/wp/v2"

// var blissfestFestivalLogoFilename = "blissfest-musical-festival-logo.png"

// WP: "https://www.blissfestfestival.org/wp-content/uploads/2022/06/blissfest-musical-festival-logo.png"
var blissfestLogoURI = "https://blissfest.org/cdn/shop/files/Bliss_Logo_2024sm.jpg?v=1735155150&width=1080"

// NewBlissfestService creates a new BlissfestService.
func NewBlissfestService(config models.BlissfestConfig) *BlissfestService {
	lgr := log.NewZapLogger().With(
		zapcore.Field{Type: zapcore.StringType, Key: ServiceLoggerFieldKey, String: "blissfest_service"},
	)

	return &BlissfestService{
		bloopymeta: models.NewBloopyMeta(),
		config:     config,
		logger:     lgr,
	}
}

// GetTimeUntilStart gets the time until the start of the event.
func (bs *BlissfestService) GetTimeUntilStart(fromDate *time.Time) time.Duration {
	comparingFrom := bs.config.Start
	if fromDate == nil {
		return time.Until(comparingFrom)
	}
	return comparingFrom.Sub(*fromDate)
}

// GetStartTime returns the start time of the event.
func (bs *BlissfestService) GetStartTime() *time.Time {
	return &bs.config.Start
}

// GetEndTime returns the end time of the event.
func (bs *BlissfestService) GetEndTime() *time.Time {
	return &bs.config.End
}

// GetTimeUntilEnd gets the time until the end of the event.
// TODO: Use humanize.Time once the pull request is merged.
// pending https://github.com/dustin/go-humanize/pull/92
// func (bs *BlissfestService) GetHumanTimeUntilStart(fromDate *time.Time) string {

//     comparingFrom := bs.config.Start
// 	if fromDate == nil {
// 		return humanize.Time(comparingFrom)
// 	}
// 	return comparingFrom.Sub(*fromDate)
// }

// GetTimeUntilEnd gets the time until the end of the event.
func (bs *BlissfestService) GetTimeUntilEnd(fromDate *time.Time) time.Duration {
	comparingFrom := bs.config.End
	if fromDate == nil {
		return time.Until(comparingFrom)
	}
	return comparingFrom.Sub(*fromDate)
}

// IsInProgress returns true if the event is in progress.
func (bs *BlissfestService) IsInProgress() bool {
	start := bs.config.Start
	end := bs.config.End
	return (start.Before(time.Now()) && end.After(time.Now()))
}

// GetLineupImageURI returns the URI for the lineup image.
func (bs *BlissfestService) GetLineupImageURI() string {
	return lineupImageURI
}

// GetShowclixTicketData returns the ticket data from Showclix.
func (bs *BlissfestService) GetShowclixTicketData() (*[]models.PriceLevel, error) {
	resp, err := http.Get(fmt.Sprintf("%s%s/%d/all_levels", models.ShowclixAPIURL, models.ShowclixAPIEventPrefix, blissfestShowclixEventID))
	if err != nil {
		bs.logger.Error("error getting showclix ticket data", zap.Error(err))
		return nil, err
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			bs.logger.Error("failed to close http response body", zap.Error(err))
		}
	}()
	result, err := io.ReadAll(resp.Body)
	if err != nil {
		bs.logger.Error("error reading showclix ticket data", zap.Error(err))
		return nil, err
	}
	var priceLevels map[int]models.PriceLevel
	err = json.Unmarshal(result, &priceLevels)
	if err != nil {
		bs.logger.Error("error unmarshalling showclix ticket data", zap.Error(err), zap.ByteString("response", result))
		return nil, err
	}
	priceLevelSlice := []models.PriceLevel{}
	for _, priceLevel := range priceLevels {
		priceLevelSlice = append(priceLevelSlice, priceLevel)
	}
	return &priceLevelSlice, nil
}

// GetAdultWeekendPriceLevel returns the price level for an adult weekend ticket.
func (bs *BlissfestService) GetAdultWeekendPriceLevel() (*models.PriceLevel, error) {
	priceLevelName := "Adult Weekend (18+)"
	priceLevels, err := bs.GetShowclixTicketData()
	if err != nil {
		bs.logger.Error("error getting price levels", zap.Error(err))
		return nil, err
	}
	for _, priceLevel := range *priceLevels {
		if priceLevel.Level == priceLevelName {
			return &priceLevel, nil
		}
	}
	bs.logger.Warn("no price level found", zap.String("priceLevelName", priceLevelName))
	return nil, nil
}

// GetBlissfestLogoURI returns the URI for the Blissfest logo.
func (bs *BlissfestService) GetBlissfestLogoURI() string {
	return blissfestLogoURI
}
