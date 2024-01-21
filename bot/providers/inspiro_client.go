package providers

import (
	"github.com/h3mmy/bloopyboi/bot/internal/config"
	"github.com/h3mmy/bloopyboi/bot/services"
)

func GetInspiroClient() *services.InspiroClient {
	conf, _ := config.GetConfig().GetFeatureConfig(services.InspiroFeatureName)
	return GetInspiroClientWithConfig(&services.InspiroConfig{
		API_url:           conf.Data[services.InspiroAPIKey],
		Backup_image_link: conf.Data[services.InspiroBackupURLKey],
	})

}

func GetInspiroClientWithConfig(config *services.InspiroConfig) *services.InspiroClient {
	return GetInspiroClientWithConfig(config)
}

func InspiroServiceWithConfig(config *services.InspiroConfig) *services.InspiroService {
	return services.NewInspiroServiceWithConfig(*config)
}
