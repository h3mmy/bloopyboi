package discord

import "time"

// A MessageEvent is created when we receive a message that
// requires our attention
type MessageEvent struct {
	CreatedAt        time.Time
	UUID             string `gorm:"primaryKey"`
	AuthorId         string
	AuthorUsername   string
	MessageId        string
	Command          string
	ChannelId        string
	ServerID         string
}
