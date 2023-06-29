package commands

import (
	"context"

	"github.com/bwmarrin/discordgo"
	"gitlab.com/h3mmy/bloopyboi/bot/internal/models"
)

// NextMessageReactionRemoveC returns a channel for the next MessageReactionRemove event
func StartChannelMessageActor(ctx context.Context, s *discordgo.Session, msCh *chan *models.DiscordMessageSendRequest) error {
	for {
		select {
		case msg := <-*msCh:
			s.ChannelMessageSendComplex(msg.ChannelID, msg.MessageComplex)
		case <-ctx.Done():
			return nil
		}
	}
}

// func StartTypingIndicatorChannel(ctx context.Context, s *discordgo.Session, typChan *chan *models.TypingIndicatorChange) error {
// 	for {
// 		select {
// 		case typ := <-*typChan:
// 			typeInChannel(*typChan, s, typ.ChannelID)
// 		}
// 	}
// }
