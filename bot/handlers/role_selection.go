package handlers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/h3mmy/bloopyboi/internal/models"
	"github.com/h3mmy/bloopyboi/pkg/config"
	log "github.com/h3mmy/bloopyboi/pkg/logs"
	"go.uber.org/zap"
)

type SelectionPrompt = config.RoleSelectionPrompt

type RoleSelectionHandler struct {
	meta        models.BloopyMeta
	config      *config.RoleSelectionConfig
	guildID     string
	logger      *zap.Logger
	prompts     map[string]SelectionPrompt
	initialized bool
}

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

func (r *RoleSelectionHandler) ReconcileConfig(s *discordgo.Session) error {
	chList, err := s.GuildChannels(r.guildID)
	if err != nil {
		r.logger.Error("error getting channels for guild", zap.String("guildID", r.guildID), zap.Error(err))
		return err
	}
	roleChExists := false
	var roleChannel *discordgo.Channel
	for _, ch := range chList {
		if ch.ID == r.config.ChannelID {
			r.logger.Debug("role channel exists!", zap.String("channel", ch.Name))
			roleChExists = true
			roleChannel = ch
		}
	}
	if !roleChExists {
		r.logger.Debug("channel does not yet exist. Wat?!")
		// create channel
	}
	messagesCreated := 0
	if roleChannel.MessageCount == 0 {
		for _, p := range r.config.Prompts {
			msg, err := s.ChannelMessageSend(roleChannel.ID, p.Message)
			if err != nil {
				r.logger.Error("error creating channel message", zap.String("channelName", roleChannel.Name), zap.String("channelId", roleChannel.ID), zap.Error(err))
				continue
			}
			r.prompts[msg.ID] = p
			messagesCreated++
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
		}
	}
	return nil
}

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
	if m.ChannelID != r.config.ChannelID {
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
	if m.ChannelID != r.config.ChannelID {
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
		err = s.GuildMemberRoleRemove(m.GuildID, m.UserID, focusRoleId)
		if err != nil {
			r.logger.Error("failed to remove role", zap.String("roleId", focusRoleId), zap.String("user", user.User.Username), zap.Error(err))
		}
	}

}
