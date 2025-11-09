package handlers

import (
	"sync"

	"github.com/bwmarrin/discordgo"
	"github.com/h3mmy/bloopyboi/internal/models"
	"github.com/h3mmy/bloopyboi/pkg/config"
	log "github.com/h3mmy/bloopyboi/pkg/logs"
	"go.uber.org/zap"
)

type SelectionPrompt = config.RoleSelectionPrompt

// RoleSelectionHandler is responsible for managing role selection through reactions.
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

// ReconcileConfig ensures that the role selection channel and messages are in the desired state.
func (r *RoleSelectionHandler) ReconcileConfig(s *discordgo.Session) error {
	r.reconciling.RLock()
	chList, err := s.GuildChannels(r.guildID)
	if err != nil {
		r.logger.Error("error getting channels for guild", zap.String("guildID", r.guildID), zap.Error(err))
		r.reconciling.RUnlock()
		return err
	}
	roleChExists := false
	var roleChannel *discordgo.Channel
	for _, ch := range chList {
		if ch.ID == r.config.Channel.ID {
			r.logger.Debug("role channel exists!", zap.String("channel", ch.Name))
			roleChExists = true
			roleChannel = ch
			break
		}
	}
	r.reconciling.RUnlock()
	if !roleChExists {
		r.logger.Error("role selection channel does not exist", zap.String("channel_id", r.config.Channel.ID))
		// create channel maybe? for now, we just exit.
		return nil
	}

	r.reconciling.Lock()
	defer r.reconciling.Unlock()

	messages, err := s.ChannelMessages(roleChannel.ID, 100, "", "", "")
	if err != nil {
		r.logger.Error("error getting messages for channel", zap.String("channelID", roleChannel.ID), zap.Error(err))
		return err
	}

	for _, p := range r.config.Prompts {
		messageExists := false
		for _, msg := range messages {
			if msg.Content == p.Message {
				messageExists = true
				if _, ok := r.prompts[msg.ID]; !ok {
					r.prompts[msg.ID] = p
				}
				break
			}
		}

		if !messageExists {
			r.logger.Info("creating prompt message", zap.String("message", p.Message))
			msg, err := s.ChannelMessageSend(roleChannel.ID, p.Message)
			if err != nil {
				r.logger.Error("error creating channel message", zap.String("channelName", roleChannel.Name), zap.String("channelId", roleChannel.ID), zap.Error(err))
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
				}
			}
		}
	}
	return nil
}

// HandleReactionAdd is called when a user adds a reaction to a message.
func (r *RoleSelectionHandler) HandleReactionAdd(s *discordgo.Session, m *discordgo.MessageReactionAdd) {
	if !r.initialized {
		err2 := r.ReconcileConfig(s)
		if err2 != nil {
			r.logger.Error("failed to reconcile config", zap.Error(err2))
		} else {
			r.initialized = true
		}
	}
	// Ignore reactions in non-target channel
	if m.ChannelID != r.config.Channel.ID {
		return
	}
	if pr, ok := r.prompts[m.MessageID]; !ok {
		r.logger.Debug("message is not registered prompt", zap.String("message", m.MessageID))
	} else {
		var focusRoleId string
		for _, op := range pr.Options {
			if op.EmojiID == m.Emoji.ID {
				focusRoleId = op.RoleID
			}
		}
		if focusRoleId == "" {
			// unrelated emoji?
			return
		}
		user, err := s.GuildMember(m.GuildID, m.UserID)
		if err != nil {
			r.logger.Error("error fetching guild member", zap.Error(err))
		}
		for _, roleID := range user.Roles {
			if roleID == focusRoleId {
				// user has role
				return
			}
		}
		err = s.GuildMemberRoleAdd(m.GuildID, m.UserID, focusRoleId)
		if err != nil {
			r.logger.Error("failed to add role", zap.String("roleId", focusRoleId), zap.String("user", user.User.Username), zap.Error(err))
		}
	}

}

// HandleReactionRemove is called when a user removes a reaction from a message.
func (r *RoleSelectionHandler) HandleReactionRemove(s *discordgo.Session, m *discordgo.MessageReactionRemove) {
	if !r.initialized {
		err2 := r.ReconcileConfig(s)
		if err2 != nil {
			r.logger.Error("failed to reconcile config", zap.Error(err2))
		} else {
			r.initialized = true
		}
	}
	// Ignore reactions in non-target channel
	if m.ChannelID != r.config.Channel.ID {
		return
	}
	if pr, ok := r.prompts[m.MessageID]; !ok {
		r.logger.Debug("message is not registered prompt", zap.String("message", m.MessageID))
	} else {
		var focusRoleId string
		for _, op := range pr.Options {
			if op.EmojiID == m.Emoji.ID {
				focusRoleId = op.RoleID
			}
		}
		if focusRoleId == "" {
			// unrelated emoji?
			return
		}
		user, err := s.GuildMember(m.GuildID, m.UserID)
		if err != nil {
			r.logger.Error("error fetching guild member", zap.Error(err))
		}
		for _, roleID := range user.Roles {
			if roleID == focusRoleId {
				err = s.GuildMemberRoleRemove(m.GuildID, m.UserID, focusRoleId)
				if err != nil {
					r.logger.Error("failed to remove role", zap.String("roleId", focusRoleId), zap.String("user", user.User.Username), zap.Error(err))
				}
				return
			}
		}
	}

}
