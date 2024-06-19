package providers

import (
	"github.com/h3mmy/bloopyboi/bot/services"
	"github.com/h3mmy/bloopyboi/pkg/config"
)

func GetInspiroService() *services.InspiroService {
	conf, _ := config.GetConfig().GetFeatureConfig(services.InspiroFeatureName)
	return GetInspiroServiceWithConfig(&services.InspiroConfig{
		API_url:           conf.Data[services.InspiroAPIKey],
		Backup_image_link: conf.Data[services.InspiroBackupURLKey],
	})
}

func GetInspiroServiceWithConfig(config *services.InspiroConfig) *services.InspiroService {
	return services.NewInspiroService(InspiroClientWithConfig(config))
}

func InspiroClientWithConfig(config *services.InspiroConfig) *services.InspiroClient {
	return services.NewInspiroClientWithConfig(*config)
}
