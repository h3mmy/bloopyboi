package asynchandlers

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/adrg/strutil/metrics"
	"github.com/bwmarrin/discordgo"
	"github.com/h3mmy/bloopyboi/internal/discord"
	"github.com/h3mmy/bloopyboi/internal/models"
	log "github.com/h3mmy/bloopyboi/pkg/logs"
	rake "github.com/afjoseph/RAKE.go"
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

	keywords := rake.RunRake(m.Content)
	
	logger.Debug("Extracted keywords", zap.Int("keyword_ct", len(keywords)))
	
	if len(keywords) == 0 {
		logger.Debug("no keywords found in message")
		return emojiPool
	}

	oc := metrics.NewOverlapCoefficient()

	var totalSimilarity float64
	var comparisons int

	for _, emoji := range emojiPool {
		for _, keyword := range keywords {
			totalSimilarity += oc.Compare(emoji.Name, keyword.Key)
			comparisons++
		}
	}

	similarityThreshold := totalSimilarity / float64(comparisons)
	logger.Debug("calculated similarity threshold", zap.Float64("threshold", similarityThreshold))

	revisedEmojiPool := []*discordgo.Emoji{}
	seen := make(map[string]bool)
	for _, emoji := range emojiPool {
		for _, keyword := range keywords {
			sim := oc.Compare(emoji.Name, keyword.Key)
			if sim > similarityThreshold && !seen[emoji.Name] {
				logger.Debug("adding emoji with similarity score",
					zap.String("emoji", emoji.Name),
					zap.String("keyword", keyword.Key),
					zap.Float64("similarity", sim),
				)
				revisedEmojiPool = append(revisedEmojiPool, emoji)
				seen[emoji.Name] = true
			}
		}
	}

	if len(revisedEmojiPool) == 0 {
		logger.Warn("no emoji similar enough")
	}
	return revisedEmojiPool
}
