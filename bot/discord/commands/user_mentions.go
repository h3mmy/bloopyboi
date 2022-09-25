package commands

import (
	"github.com/bwmarrin/discordgo"
	"fmt"
)


// Listens for messages specifically addressing bot
func DirectedMessageReceive(s *discordgo.Session, m *discordgo.MessageCreate) {
	directMessage := (m.GuildID == "")

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	logger.Debug(fmt.Sprintf("Processing Message from %s with Content %s", m.Author.Username, m.Content))
	botMentioned := false
	// Filter only commands we care about
	if len(m.Mentions) > 0 {
		for _, user := range m.Mentions {
			if user.ID == s.State.User.ID {
				botMentioned = true
				break
			}
		}
	}

	if !directMessage && botMentioned {
		logger.Sugar().Debug(fmt.Sprintf("Detected Channel Mention in message from %s with UserID %s and Content: ", m.Author.Username, m.Author.ID), m.Content)
		err := s.MessageReactionAdd(m.ChannelID, m.ID,emojiZoop)
		if err != nil {
			logger.Sugar().Error(err)
			return
		}
	} else if directMessage {
		logger.Sugar().Debug("Detected Direct Message from ", m.Author.Username)
		channel, err := s.UserChannelCreate(m.Author.ID)
		if err != nil {
			logger.Sugar().Error(err)
			return
		}
		_, err = s.ChannelMessageSend(channel.ID, "Aye, I have received your direct message. Afraid this is still in development")
		if err != nil {
			logger.Sugar().Error(err)
			return
		}
	}

}
