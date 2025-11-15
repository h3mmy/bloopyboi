package providers

import (
	"fmt"

	"github.com/h3mmy/bloopyboi/pkg/config"
	"golang.org/x/oauth2"
)

func GetGatewayConfig() *config.GatewayConfig {
	appConfig := config.GetConfig()
	return &config.GatewayConfig{
		HttpServerConfig: appConfig.HttpConfig,
		GrpcServerConfig: appConfig.GrpcConfig,
	}
}

func GetDiscordOauthConfig() *oauth2.Config {
	appConfig := config.GetConfig()

	redirectUrl := fmt.Sprintf("https://%s%s/discord/linked-roles/callback",
		appConfig.HttpConfig.Hostname,
		appConfig.HttpConfig.BaseUrl)

	return &oauth2.Config{
		ClientID:     fmt.Sprint(appConfig.DiscordConfig.AppID),
		ClientSecret: appConfig.DiscordConfig.ClientSecret,
		RedirectURL:  redirectUrl,
		Scopes:       []string{"identify", "role_connections.write"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://discord.com/oauth2/authorize",
			TokenURL: "https://discord.com/api/oauth2/token",
		},
	}
}
