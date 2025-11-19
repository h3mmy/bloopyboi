package commands

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

var NoticedReactionPool []string = []string{"👀","🙊","🙈","🙉","👁️","👄","🫦","✍🏽","🐸","🐢","🥁","🔬","🔭","⁉️","🆒"}

//TODO: Migrate to asynchandlers
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

	botMentioned := false
	// Filter only commands we care about (or not)
	// TODO: Replace with logic to detect engagement
	if len(m.Mentions) > 0 {
		// Just react to some mentions mysteriously
		if rand.Float32() < 0.5 {
			var err error
			if nsfwContext {
				switch rand.Intn(5) {
				case 0:
					err = s.MessageReactionAdd(m.ChannelID, m.ID, "imwetrn:1236826185783316552")
				case 1:
					err = s.MessageReactionAdd(m.ChannelID, m.ID, "animeass:1236826184663175188")
				case 2:
					err = s.MessageReactionAdd(m.ChannelID, m.ID, "peepeekun:1236821722460848139")
				case 3:
					err = s.MessageReactionAdd(m.ChannelID, m.ID, "rikkaSpank:1236824309083803790")
				case 4:
					err = s.MessageReactionAdd(m.ChannelID, m.ID, "frog_booty:1180327466775105617")
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
		reactn := NoticedReactionPool[rand(Intn(len(NoticedReactionPool)))]
		err := s.MessageReactionAdd(m.ChannelID, m.ID, reactn)
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
