package asynchandlers

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/h3mmy/bloopyboi/bot/internal/log"
	"github.com/h3mmy/bloopyboi/bot/internal/models"
	"go.uber.org/zap"
)

type MessageReactor struct {
	meta   models.BloopyMeta
	logger *zap.Logger
}

func NewMessageReactor() *MessageReactor {
	bmeta := models.NewBloopyMeta()
	logger := log.NewZapLogger().Named("message_reactor")
	return &MessageReactor{
		meta:   bmeta,
		logger: logger,
	}
}

func (mr *MessageReactor) Handle(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's 1=a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	logger := mr.logger.With(zap.String("method", "Handle"), zap.String("messageID", m.ID))
	logger.Debug(fmt.Sprintf("Processing Message from %s with Content %s", m.Author.Username, m.Content))
	if mr.ShouldAddReaction(s, m.Message) {
		logger.Debug("Will add reaction")
	} else {
		logger.Debug("Will NOT add reaction")
	}
}

func (mr *MessageReactor) ShouldAddReaction(s *discordgo.Session, m *discordgo.Message) bool {
	logger := mr.logger.With(zap.String("method", "ShouldAddReaction"))
	if len(m.Mentions) > 0 {
		return true
	}
	if m.GuildID == "" {
		// Implies a DM
		return false
	}
	if m.Type == discordgo.MessageTypeReply {
		logger.Debug(
			"message is a reply type",
			zap.String("channelID", m.ChannelID),
			zap.String("messageID", m.ID),
		)
		// react to the referenced message
		err := mr.ReactToMessage(s, m.ReferencedMessage)
		if err != nil {
			logger.Warn("failed reacting to referenced message", zap.Error(err))
		}
		return true
	}
	lastChannelMessages, err := s.ChannelMessages(m.ChannelID, 1, m.ID, "", "")

	if err != nil {
		logger.Warn(
			"could not get last channel message",
			zap.String("channelID", m.ChannelID),
			zap.String("messageID", m.ID),
			zap.Error(err),
		)
		return false
	} else {
		lastMessage := lastChannelMessages[0]
		if lastMessage != nil {
			logger.Debug("last message is nil for some reason",
				zap.String("channelID", m.ChannelID),
				zap.String("messageID", m.ID),
			)
			return true
		}
		timeDiff := lastMessage.Timestamp.Sub(m.Timestamp)
		logger.Debug("time difference between messages", zap.Duration("timeDiff", timeDiff))
		if timeDiff < 7*time.Minute {
			return true
		}
	}
	return rand.Float64() < 0.6
}

func (mr *MessageReactor) ReactToMessage(s *discordgo.Session, m *discordgo.Message) error {
	logger := mr.logger.With(zap.String("method", "ReactToMessage"))
	guildEmojis, err := s.GuildEmojis(m.GuildID)
	if err != nil {
		logger.Warn("could not get emoji for guild", zap.String("guildID", m.GuildID))
	}
	if guildEmojis != nil {
		emj := mr.SelectGuildEmojiForReaction(guildEmojis)
		if emj.Available {
			err = s.MessageReactionAdd(m.ChannelID, m.ID, emj.APIName())
		} else {
			err = s.MessageReactionAdd(m.ChannelID, m.ID, "👁‍🗨")
		}
	} else {
		err = s.MessageReactionAdd(m.ChannelID, m.ID, "👁‍🗨")
	}
	return err
}

func (mr *MessageReactor) SelectGuildEmojiForReaction(emojiPool []*discordgo.Emoji) *discordgo.Emoji {
	return emojiPool[rand.Intn(len(emojiPool))]
}
