package providers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/h3mmy/bloopyboi/bot/services"
	"go.uber.org/zap"
)

// NewDiscordServiceWithToken creates a new DiscordService with a token
// Oauth tokens need to be prefixed with "Bearer " instead so this won't work for that
func NewDiscordServiceWithToken(token string) (*services.DiscordService, error) {
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		logger.Error("Failed to create discord session", zap.Error(err))
		return nil, err
	}
	return NewDiscordServiceWithSession(session), nil
}

func NewDiscordServiceWithSession(session *discordgo.Session) *services.DiscordService {
	return services.NewDiscordService(session)
}
