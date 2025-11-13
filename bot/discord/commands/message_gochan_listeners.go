package commands

import "github.com/bwmarrin/discordgo"

const defaultChannelBuffer = 20

// NextMessageCreateC returns a channel for the next MessageCreate event
func NextMessageCreateC(s *discordgo.Session) chan *discordgo.MessageCreate {
	out := make(chan *discordgo.MessageCreate, defaultChannelBuffer)
	s.AddHandler(func(_ *discordgo.Session, e *discordgo.MessageCreate) {
		out <- e
	})
	return out
}

// NextMessageReactionAddC returns a channel for the next MessageReactionAdd event
func NextMessageReactionAddC(s *discordgo.Session) chan *discordgo.MessageReactionAdd {
	out := make(chan *discordgo.MessageReactionAdd, defaultChannelBuffer)
	s.AddHandler(func(_ *discordgo.Session, e *discordgo.MessageReactionAdd) {
		out <- e
	})
	return out
}

// NextMessageReactionRemoveC returns a channel for the next MessageReactionRemove event
func NextMessageReactionRemoveC(s *discordgo.Session) chan *discordgo.MessageReactionRemove {
	out := make(chan *discordgo.MessageReactionRemove, defaultChannelBuffer)
	s.AddHandler(func(_ *discordgo.Session, e *discordgo.MessageReactionRemove) {
		out <- e
	})
	return out
}
