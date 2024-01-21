package providers

import (
	"time"

	"github.com/h3mmy/bloopyboi/bot/internal/models"
	"github.com/h3mmy/bloopyboi/bot/services"
)

func GetBlissfestService() *services.BlissfestService {
	location, _ := time.LoadLocation("America/Detroit")
	config := models.BlissfestConfig{
		Start:    time.Date(2023, 7, 7, 9, 0, 0, 0, location),
		End:      time.Date(2023, 7, 7, 9, 0, 0, 0, location),
		Homepage: "https://www.blissfest.org/",
	}
	return services.NewBlissfestService(config)
}
