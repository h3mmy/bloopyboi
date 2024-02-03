package discord

import (
	"time"
)

// A MessageEvent is created when we receive a message that
// requires our attention
type MessageEvent struct {
	CreatedAt      time.Time
	UUID           string
	AuthorId       string
	AuthorUsername string
	MessageId      string
	Command        string
	ChannelId      string
	ServerID       string
}

// createMessageEvent logs a given message event into the database.
// func (dc *DiscordManager) createMessageEvent(c string, m *discordgo.Message) {
// 	uuid := uuid.New().String()
// 	bot.DB.Create(&MessageEvent{
// 		UUID:           uuid,
// 		AuthorId:       m.Author.ID,
// 		AuthorUsername: m.Author.Username,
// 		MessageId:      m.ID,
// 		Command:        c,
// 		ChannelId:      m.ChannelID,
// 		ServerID:       m.GuildID,
// 	})
// }
