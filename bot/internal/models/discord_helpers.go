package models

import (
	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
)

type DiscordMessageSendRequest struct {
	EventID        uuid.UUID
	ChannelID      string
	MessageComplex *discordgo.MessageSend
}

type TypingIndicatorChange struct {
	IsTyping bool
	ChannelID string
}
