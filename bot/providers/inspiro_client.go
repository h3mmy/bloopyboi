package providers

import "gitlab.com/h3mmy/bloopyboi/bot/services"

func GetInspiroClient() *services.InspiroClient {
	return services.NewInspiroClient()
}
