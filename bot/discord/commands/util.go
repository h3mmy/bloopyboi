package commands

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	log "github.com/h3mmy/bloopyboi/pkg/logs"
	"go.uber.org/zap/zapcore"
)

const (
	emojiZoop = ":zoop:872664064889872415"
)

var (
	logger = log.NewZapLogger().With(
		zapcore.Field{
			Key:    "bot",
			Type:   zapcore.StringType,
			String: "Discord.Commands",
		})
)

// typeInChannel sets the typing indicator for a channel. The indicator is cleared
// when a message is sent.
func typeInChannel(channel chan bool, s *discordgo.Session, channelID string) {
	for {
		select {
		case <-channel:
			return
		default:
			if err := s.ChannelTyping(channelID); err != nil {
				fmt.Println("unable to set typing indicator: ", err)
			}
			time.Sleep(time.Second * 5)
		}
	}
}
