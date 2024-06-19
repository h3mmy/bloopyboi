package providers

import "github.com/h3mmy/bloopyboi/internal/models"

func GetBloopyServiceRegistry() models.ServiceRegistry {
	return models.NewBloopyServiceBroker()
}
