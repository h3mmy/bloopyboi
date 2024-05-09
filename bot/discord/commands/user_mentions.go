package commands

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

// Listens for messages specifically addressing bot
func DirectedMessageReceive(s *discordgo.Session, m *discordgo.MessageCreate) {
	directMessage := (m.GuildID == "")

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's 1=a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	logger.Debug(fmt.Sprintf("Processing Message from %s with Content %s", m.Author.Username, m.Content))
	var nsfwContext bool
	channel, err := s.Channel(m.ChannelID)

	if err != nil {
		logger.Warn("could not get information for channel", zap.String("channelID", m.ChannelID))
	} else {
		nsfwContext = channel.NSFW
	}

	guildEmojis, err := s.GuildEmojis(m.GuildID)
	if err != nil {
		logger.Warn("could not get emoji for guild", zap.String("guildID", m.GuildID))
	}

	botMentioned := false
	// Filter only commands we care about (or not)
	// TODO: Replace with logic to detect engagement
	if len(m.Mentions) > 0 || rand.Float32() < 0.5 {
		// Just react to some mentions mysteriously
		if rand.Float32() < 0.5 {
			var err error
			if nsfwContext {
				err = s.MessageReactionAdd(m.ChannelID, m.ID, "imwetrn:1236826185783316552")
			} else {
				if guildEmojis != nil {
					emj := selectGuildEmojiForReaction(guildEmojis)
					if emj.Available {
						err = s.MessageReactionAdd(m.ChannelID, m.ID, emj.APIName())
					} else {
						err = s.MessageReactionAdd(m.ChannelID, m.ID, "👁‍🗨")
					}
				} else{
				err = s.MessageReactionAdd(m.ChannelID, m.ID, "👁‍🗨")
			}
			}
			if err != nil {
				logger.Warn(fmt.Sprintf("Error adding reaction to message %s from user %s", m.ID, m.Author.Username))
			}
		}
		for _, user := range m.Mentions {
			if user.ID == s.State.User.ID {
				botMentioned = true
				break
			}
		}
	}
	if strings.Contains(strings.ToLower(m.Content), "bloopyboi") {
		logger.Sugar().Debug("Detected BloopyBoi in message from ", m.Author.Username)
		err := s.MessageReactionAdd(m.ChannelID, m.ID, "👀")
		if err != nil {
			logger.Sugar().Warn(err)
		}
	}

	if !directMessage && botMentioned {
		logger.Sugar().Debug(fmt.Sprintf("Detected Channel Mention in message from %s with UserID %s and Content: ", m.Author.Username, m.Author.ID), m.Content)
		err := s.MessageReactionAdd(m.ChannelID, m.ID, emojiZoop)
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
