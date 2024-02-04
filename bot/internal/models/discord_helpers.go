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

type DiscordAppCommand interface {
	GetAppCommand() *discordgo.ApplicationCommand
	GetAppCommandHandler() func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

type DiscordCommand interface {
	GetCommand() *discordgo.ApplicationCommand
	GetCommandHandler() func(s *discordgo.Session, i *discordgo.InteractionCreate)
}
