package providers

import (
	"time"

	pkgmodels "github.com/h3mmy/bloopyboi/pkg/models"
	"github.com/h3mmy/bloopyboi/bot/services"
)

func GetBlissfestService() *services.BlissfestService {
	location, _ := time.LoadLocation("America/Detroit")
	config := pkgmodels.BlissfestConfig{
		Start:    time.Date(2023, 7, 7, 9, 0, 0, 0, location),
		End:      time.Date(2023, 7, 7, 9, 0, 0, 0, location),
		Homepage: "https://www.blissfest.org/",
	}
	return services.NewBlissfestService(config)
}
