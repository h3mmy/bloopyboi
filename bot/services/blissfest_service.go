package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/h3mmy/bloopyboi/bot/internal/log"
	"github.com/h3mmy/bloopyboi/bot/internal/models"
	pkgmodels "github.com/h3mmy/bloopyboi/internal/models"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// 2023: "https://www.blissfestfestival.org/wp-content/uploads/2023/04/Bliss23_LineUpIG-2-2048x2048.jpg"
// 2024: "https://www.blissfestfestival.org/wp-content/uploads/2024/04/Bliss24_IGAnnouncement3-2048x2048.jpg"
var lineupImageURI = "https://www.blissfestfestival.org/wp-content/uploads/2024/04/Bliss24_IGAnnouncement3-2048x2048.jpg"

// 2024 blissfest showclix "event_id": 9297272, "parent_event_id": 8615552,
// 2024 blissfest showclix venue_id = 64139

var blissfestShowclixEventID = 9297272

// 2024 blissfest logo art "https://www.blissfestfestival.org/wp-content/uploads/2024/01/Bliss_Logo_2024F3.png"

// var apiPrefix = "/wp-json/wp/v2"

// var blissfestFestivalLogoFilename = "blissfest-musical-festival-logo.png"

var blissfestLogoURI = "https://www.blissfestfestival.org/wp-content/uploads/2022/06/blissfest-musical-festival-logo.png"

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

func (bs *BlissfestService) GetShowclixTicketData() (*[]pkgmodels.PriceLevel, error) {
	resp, err := http.Get(fmt.Sprintf("%s%s/%d/all_levels", pkgmodels.ShowclixAPIURL, pkgmodels.ShowclixAPIEventPrefix, blissfestShowclixEventID))
	if err != nil {
		bs.logger.Error("error getting showclix ticket data", zap.Error(err))
		return nil, err
	}

	defer resp.Body.Close()
	result, err := io.ReadAll(resp.Body)
	if err != nil {
		bs.logger.Error("error reading showclix ticket data", zap.Error(err))
		return nil, err
	}
	var priceLevels map[int]pkgmodels.PriceLevel
	err = json.Unmarshal(result, &priceLevels)
	if err != nil {
		bs.logger.Error("error unmarshalling showclix ticket data", zap.Error(err), zap.ByteString("response", result))
		return nil, err
	}
	priceLevelSlice := []pkgmodels.PriceLevel{}
	for _, priceLevel := range priceLevels {
		priceLevelSlice = append(priceLevelSlice, priceLevel)
	}
	return &priceLevelSlice, nil
}

func (bs *BlissfestService) GetAdultWeekendPriceLevel() (*pkgmodels.PriceLevel, error) {
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

func (bs *BlissfestService) GetBlissfestLogoURI() string {
	return blissfestLogoURI
}
