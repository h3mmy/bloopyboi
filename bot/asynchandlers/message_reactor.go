package asynchandlers

import (
	"fmt"
	"math/rand"
	"slices"
	"time"

	"github.com/adrg/strutil/metrics"
	"github.com/bwmarrin/discordgo"
	"github.com/h3mmy/bloopyboi/internal/models"
	"github.com/h3mmy/bloopyboi/internal/discord"
	log "github.com/h3mmy/bloopyboi/pkg/logs"
	"github.com/kljensen/snowball"
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
		if rand.Float64() < 0.1 {
		  err := mr.ReactToMessage(s, m.Message)
		  if err != nil {
			  logger.Error("failed reacting to message", zap.Error(err))
		  }
		}
	} else {
		logger.Debug("Will NOT add reaction")
	}
}

func (mr *MessageReactor) ShouldAddReaction(s *discordgo.Session, m *discordgo.Message) bool {
	logger := mr.logger.With(zap.String("method", "ShouldAddReaction"))
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's 1=a good practice.
	if m.Author.ID == s.State.User.ID {
		return false
	}
	if len(m.Mentions) > 0 {
		return rand.Float64() < 0.55
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
		// s.ChannelMessage(m.ChannelID, m.ID)
		err := mr.ReactToMessage(s, m.ReferencedMessage)
		if err != nil {
			logger.Warn("failed reacting to referenced message", zap.Error(err))
			return false
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
		logger.Debug("found the last message?", zap.Int("lastChannelMessages size", len(lastChannelMessages)))
		lastMessage := lastChannelMessages[0]
		if lastMessage != nil {
			logger.Debug("last message is not nil",
				zap.String("channelID", m.ChannelID),
				zap.String("messageID", m.ID),
			)
			return rand.Float64() < 0.5
		}
		lastMsgTimestamp, err := discord.SnowflakeTimestamp(lastMessage.ID)
		if err != nil {
			logger.Warn(
				"error calculating snowflake timestamp",
				zap.String("messageID", lastMessage.ID),
				zap.Error(err),
			)
			lastMsgTimestamp = lastMessage.Timestamp
		}
		timeDiff := lastMsgTimestamp.Sub(m.Timestamp)
		logger.Debug("time difference between messages", zap.Duration("timeDiff", timeDiff))
		if timeDiff < 7*time.Minute {
			_ = mr.ReactToMessage(s, lastMessage)
			return rand.Float64() < 0.55
		}
	}
	return rand.Float64() < 0.4
}

func (mr *MessageReactor) ReactToMessage(s *discordgo.Session, m *discordgo.Message) error {
	logger := mr.logger.With(zap.String("method", "ReactToMessage"), zap.String("messageID", m.ID))
	guildEmojis, err := s.GuildEmojis(m.GuildID)
	if err != nil {
		logger.Warn("could not get emoji for guild", zap.String("guildID", m.GuildID))
	}
	if guildEmojis != nil {
		logger.Debug("Found Guild Emojis", zap.Int("count", len(guildEmojis)))
		emj := mr.SelectGuildEmojiForReaction(m, guildEmojis)
		if emj.Available {
			logger.Debug("selected emoji is available", zap.String("emoji", emj.APIName()))
			err = s.MessageReactionAdd(m.ChannelID, m.ID, emj.APIName())
		} else {
			logger.Debug("selected emoji is not available", zap.String("emoji", emj.APIName()))
			err = s.MessageReactionAdd(m.ChannelID, m.ID, "👁‍🗨")
		}
	} else {
		logger.Debug("No guild emojis found. Using default")
		err = s.MessageReactionAdd(m.ChannelID, m.ID, "👁‍🗨")
	}
	return err
}

func (mr *MessageReactor) SelectGuildEmojiForReaction(m *discordgo.Message, emojiPool []*discordgo.Emoji) *discordgo.Emoji {
	siftedEmojiPool := mr.FindSimilarEmoji(m, emojiPool)
	return siftedEmojiPool[rand.Intn(len(siftedEmojiPool))]
}

// This is extremely crude at the moment. I intend to use something like james-bowman/nlp to properly check semantic similarity in the future
func (mr *MessageReactor) FindSimilarEmoji(m *discordgo.Message, emojiPool []*discordgo.Emoji) []*discordgo.Emoji {
	logger := mr.logger.With(zap.String("method", "FindSimilarEmoji"), zap.String("messageID", m.ID))

	stemmed, err := snowball.Stem(m.Content, "english", true)
	if err != nil {
		logger.Error("error while stemming", zap.Error(err))
		stemmed = m.Content
	}
	logger.Debug("stemmed a thing", zap.String("post stemming", stemmed))
	oc := metrics.NewOverlapCoefficient()
	revisedEmojiPool := []*discordgo.Emoji{}
	highestSim := 0.0
	// def not efficient
	for _, emoji := range emojiPool {
		sim := oc.Compare(emoji.Name, stemmed)
		if sim >= float64(highestSim) {
			highestSim = sim
			if len(revisedEmojiPool) > 5 {
				revisedEmojiPool = slices.Delete(revisedEmojiPool, 0, 1)
			}
			logger.Debug(fmt.Sprintf("Adding with similarity score: .2%f", sim), zap.String("emoji", emoji.Name), zap.String("stemmed", stemmed))
			revisedEmojiPool = append(revisedEmojiPool, emoji)
		}
	}
	if len(revisedEmojiPool) == 0 {
		logger.Warn("no emoji similar enough. Returning OG pool")
		return emojiPool
	}
	return revisedEmojiPool
}
