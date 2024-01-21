package providers

import "github.com/h3mmy/bloopyboi/bot/internal/models"

func GetBloopyServiceRegistry() models.ServiceRegistry {
	return models.NewBloopyServiceBroker()
}
