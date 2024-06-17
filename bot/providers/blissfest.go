package providers

import (
	"fmt"
	"time"

	"github.com/h3mmy/bloopyboi/pkg/config"
	"github.com/h3mmy/bloopyboi/bot/services"
	pkgmodels "github.com/h3mmy/bloopyboi/internal/models"
	"go.uber.org/zap"
)

const BlissfestStartDateKey = "start_date"
const BlissfestHomepage = "https://www.blissfest.org/"

// Blissfest is in MI and bound to the US Eastern Timezone
// Since it always happens in the summer we can assume EDT
var blissfestTZ = time.FixedZone("UTC-4", -4*60*60)
var defaultBlissfestStartDate = time.Date(2024, 7, 12, 0, 0, 0, 0, blissfestTZ)

func GetBlissfestService() *services.BlissfestService {
	// Blissfest is in MI and bound to the US Eastern Timezone
	if IsFeatureEnabled(pkgmodels.BlissfestFeatureKey) {
		// check for provided start date
		cfg := GetFeatures()[pkgmodels.BlissfestFeatureKey]
		startDate := getBlissfestStartDate(cfg)

		year, month, day := startDate.Date()
		logger.Debug(fmt.Sprintf("startDate year: %d, month: %d, day: %d", year, month, day))
		// blissfest always starts at 9am on a Friday
		finalStartDate := time.Date(year, month, day, 9, 0, 0, 0, blissfestTZ)
		logger.Debug("finalized blissfest start date", zap.Time("startDate", finalStartDate))
		// everyone is supposed to be out by noon on the following Monday (3 days + 3 hours => 4500min => 75 hours)
		finalEndDate := finalStartDate.Add(75 * time.Hour)
		logger.Debug("finalized blissfest end date", zap.Time("endDate", finalEndDate))
		return services.NewBlissfestService(pkgmodels.BlissfestConfig{
			Start:    finalStartDate,
			End:      finalEndDate,
			Homepage: BlissfestHomepage,
		})
	}
	logger.Warn("blissfest feature not enabled")

	return nil
}

func getBlissfestStartDate(cfg config.FeatureConfig) time.Time {
	if cfg.Data != nil {
		startDateString, ok := cfg.Data[BlissfestStartDateKey]
		if ok {
			logger.Debug("parsing start date", zap.String("providedDate", startDateString))
			startDate, err := time.Parse("2006-01-02", startDateString)
			if err != nil {
				logger.Error("error parsing configured start date", zap.Error(err))
				startDate = defaultBlissfestStartDate
			}
			logger.Debug("finished parsing start date", zap.String("providedDate", startDateString), zap.Time("startDate", startDate))
			return startDate
		}
	}
	return defaultBlissfestStartDate
}
