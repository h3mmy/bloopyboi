package providers

import "gitlab.com/h3mmy/bloopyboi/bot/internal/models"

func GetBloopyServiceRegistry() models.ServiceRegistry {
	return models.NewBloopyServiceBroker()
}
