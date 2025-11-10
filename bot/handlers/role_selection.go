package handlers

import (
	"fmt"
	"sync"

	"github.com/bwmarrin/discordgo"
	"github.com/h3mmy/bloopyboi/internal/models"
	"github.com/h3mmy/bloopyboi/pkg/config"
	log "github.com/h3mmy/bloopyboi/pkg/logs"
	"go.uber.org/zap"
)

// SelectionPrompt is a type alias for config.RoleSelectionPrompt.
type SelectionPrompt = config.RoleSelectionPrompt

// RoleSelectionHandler is a handler that manages role selection through reactions.
type RoleSelectionHandler struct {
	meta        models.BloopyMeta
	config      *config.RoleSelectionConfig
	guildID     string
	logger      *zap.Logger
	prompts     map[string]SelectionPrompt
	initialized bool
	reconciling sync.RWMutex
}

// NewRoleSelectionHandler creates a new RoleSelectionHandler.
func NewRoleSelectionHandler(guildID string, config *config.RoleSelectionConfig) *RoleSelectionHandler {
	bmeta := models.NewBloopyMeta()
	logger := log.NewZapLogger().Named("role_selection_handler")
	logger.Debug("I'm alive! Sort of...")
	return &RoleSelectionHandler{
		meta:        bmeta,
		logger:      logger,
		config:      config,
		guildID:     guildID,
		prompts:     make(map[string]SelectionPrompt),
		initialized: false,
	}
}

// ReconcileConfig reconciles the role selection configuration with the Discord guild.
func (r *RoleSelectionHandler) ReconcileConfig(s *discordgo.Session) error {
	r.reconciling.Lock()
	defer r.reconciling.Unlock()

	roleChannel, err := s.Channel(r.config.Channel.ID)
	if err != nil {
		r.logger.Error("error getting channel", zap.String("channelID", r.config.Channel.ID), zap.Error(err))
		return err
	}
	if roleChannel == nil {
		r.logger.Error("role channel does not exist", zap.String("channelID", r.config.Channel.ID))
		return nil
	}

	limit := len(r.config.Prompts) * 2
	if limit > 100 {
		limit = 100
	}
	messages, err := s.ChannelMessages(roleChannel.ID, limit, "", "", "")
	if err != nil {
		r.logger.Error("error getting messages for channel", zap.String("channelID", roleChannel.ID), zap.Error(err))
		return err
	}
	// TODO: add logic for prompt removal when removed from configuration
	if len(messages) > len(r.config.Prompts)*2 {
		r.logger.Warn("more messages in channel than configured prompts",
			zap.Int("messageCount", len(messages)),
			zap.Int("promptCount", len(r.config.Prompts)))
	}

	for _, p := range r.config.Prompts {
		parsedFromConfig := promptToEmbed(p)

		var messageExists bool
		var messageFieldsMatch bool
		var existingMessage *discordgo.Message

		for _, m := range messages {
			if len(m.Embeds) > 0 && m.Embeds[0].Title == parsedFromConfig.Title {
				messageExists = true
				if len(m.Embeds[0].Fields) == len(parsedFromConfig.Fields) {
					for i, f := range m.Embeds[0].Fields {
						if f != parsedFromConfig.Fields[i] {
							messageFieldsMatch = false
							break
						}
					}
					messageFieldsMatch = true
				} else {
					messageFieldsMatch = false
				}
				existingMessage = m
				break
			}
		}

		if messageExists && !messageFieldsMatch {
			// TODO: Add logic to update the existing message
		}

		if !messageExists {
			msg, err := s.ChannelMessageSendComplex(roleChannel.ID, &discordgo.MessageSend{
				Embeds: []*discordgo.MessageEmbed{parsedFromConfig},
			})
			if err != nil {
				r.logger.Error("error creating channel message",
					zap.String("channelName", roleChannel.Name),
					zap.String("channelId", roleChannel.ID),
					zap.Error(err))
				continue
			}
			r.prompts[msg.ID] = p
			for _, op := range p.Options {
				err := s.MessageReactionAdd(msg.ChannelID, msg.ID, op.EmojiID)
				if err != nil {
					r.logger.Error("error adding reaction",
						zap.String("channelName", roleChannel.Name),
						zap.String("channelId", roleChannel.ID),
						zap.String("emojiID", op.EmojiID),
						zap.String("option", op.Description),
						zap.Error(err))
					continue
				}
			}
		} else {
			r.prompts[existingMessage.ID] = p
			for _, op := range p.Options {
				var reactionExists bool
				for _, reaction := range existingMessage.Reactions {
					if reaction.Emoji.ID == op.EmojiID {
						reactionExists = true
						break
					}
				}
				if !reactionExists {
					err := s.MessageReactionAdd(existingMessage.ChannelID, existingMessage.ID, op.EmojiID)
					if err != nil {
						r.logger.Error("error adding reaction",
							zap.String("channelName", roleChannel.Name),
							zap.String("channelId", roleChannel.ID),
							zap.String("emojiID", op.EmojiID),
							zap.String("option", op.Description),
							zap.Error(err))
						continue
					}
				}
			}
		}
	}
	r.initialized = true
	return nil
}

// handleReaction is a helper function to handle both reaction add and remove events.
// It returns the role ID associated with the reaction, the member that performed the reaction op, and an error if one occurred.
func (r *RoleSelectionHandler) handleReaction(s *discordgo.Session, mReaction *discordgo.MessageReaction, member *discordgo.Member) (string, *discordgo.Member, error) {
	r.reconciling.RLock()
	if !r.initialized {
		r.reconciling.RUnlock()
		if err := r.ReconcileConfig(s); err != nil {
			r.logger.Error("failed to reconcile config", zap.Error(err))
			return "", nil, err
		}
		r.reconciling.RLock()
	}

	if mReaction.ChannelID != r.config.Channel.ID {
		r.reconciling.RUnlock()
		return "", nil, nil // Not the channel we're watching
	}

	pr, ok := r.prompts[mReaction.MessageID]
	r.reconciling.RUnlock()
	if !ok {
		r.logger.Debug("message is not a registered prompt", zap.String("messageID", mReaction.MessageID))
		return "", nil, nil // Not a message we're watching
	}

	var focusRoleID string
	for _, op := range pr.Options {
		if op.EmojiID == mReaction.Emoji.ID || op.EmojiID == mReaction.Emoji.Name {
			focusRoleID = op.RoleID
			break
		}
	}

	if focusRoleID == "" {
		return "", nil, nil // Not an emoji we're watching
	}

	if member != nil {
		return focusRoleID, member, nil
	}
	guildID := mReaction.GuildID
	userID := mReaction.UserID

	user, err := s.GuildMember(guildID, userID)
	if err != nil {
		r.logger.Error("error fetching guild member", zap.Error(err))
		return "", nil, err
	}

	return focusRoleID, user, nil
}

// HandleReactionAdd handles a reaction add event.
func (r *RoleSelectionHandler) HandleReactionAdd(s *discordgo.Session, m *discordgo.MessageReactionAdd) {
	r.logger.Debug("processing ReactionAdd", zap.Any("message", m))
	focusRoleID, user, err := r.handleReaction(s, m.MessageReaction, m.Member)
	if err != nil {
		r.logger.Error("error finding associated role", zap.Any("message", m))
		return
	}

	if user == nil {
		r.logger.Warn("user is nil", zap.Any("message", m))
		user = m.Member
	}

	for _, roleID := range user.Roles {
		if roleID == focusRoleID {
			r.logger.Debug("user already has role", zap.String("roleId", focusRoleID), zap.String("user", user.User.Username))
			return // User already has the role
		}
	}

	if err := s.GuildMemberRoleAdd(m.GuildID, m.UserID, focusRoleID); err != nil {
		r.logger.Error("failed to add role", zap.String("roleId", focusRoleID), zap.String("user", user.User.Username), zap.Error(err))
	} else {
		r.logger.Debug("added role to user", zap.String("roleId", focusRoleID), zap.String("user", user.User.Username))
	}
}

// HandleReactionRemove handles a reaction remove event.
func (r *RoleSelectionHandler) HandleReactionRemove(s *discordgo.Session, m *discordgo.MessageReactionRemove) {
	r.logger.Debug("processing ReactionRemove", zap.Any("message", m))
	focusRoleID, user, err := r.handleReaction(s, m.MessageReaction, nil)
	if err != nil {
		r.logger.Error("error finding associated role", zap.Any("message", m))
		return
	}

	var hasRole bool
	for _, roleID := range user.Roles {
		if roleID == focusRoleID {
			r.logger.Debug("user has role", zap.String("roleId", focusRoleID), zap.String("user", user.User.Username))
			hasRole = true
			break
		}
	}

	if hasRole {
		if err := s.GuildMemberRoleRemove(m.GuildID, m.UserID, focusRoleID); err != nil {
			r.logger.Error("failed to remove role", zap.String("roleId", focusRoleID), zap.String("user", user.User.Username), zap.Error(err))
		} else {
			r.logger.Debug("removed role from user", zap.String("roleId", focusRoleID), zap.String("user", user.User.Username))
		}
	}
}

func promptToEmbed(p SelectionPrompt) *discordgo.MessageEmbed {
	fields := []*discordgo.MessageEmbedField{}
	for _, opt := range p.Options {
		var emoj string
		// snowflakes are 18 or 19 digits
		if len(opt.EmojiID) > 17 {
			emoj = fmt.Sprintf("<:%s:%s>", "custom", opt.EmojiID)
		} else {
			emoj = opt.EmojiID
		}
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:  emoj,
			Value: fmt.Sprintf("<@&%s> - %s", opt.RoleID, opt.Description),
		})
	}
	embed := &discordgo.MessageEmbed{
		Title:  p.Message,
		Fields: fields,
	}
	return embed
}
