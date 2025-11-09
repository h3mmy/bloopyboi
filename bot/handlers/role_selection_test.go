package handlers

import (
	"errors"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/h3mmy/bloopyboi/pkg/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockDiscordSession is a mock of DiscordSession for testing.
type MockDiscordSession struct {
	mock.Mock
}

func (m *MockDiscordSession) GuildChannels(guildID string, options ...discordgo.RequestOption) ([]*discordgo.Channel, error) {
	args := m.Called(guildID, options)
	return args.Get(0).([]*discordgo.Channel), args.Error(1)
}

func (m *MockDiscordSession) ChannelMessages(channelID string, limit int, beforeID, afterID, aroundID string, options ...discordgo.RequestOption) ([]*discordgo.Message, error) {
	args := m.Called(channelID, limit, beforeID, afterID, aroundID, options)
	return args.Get(0).([]*discordgo.Message), args.Error(1)
}

func (m *MockDiscordSession) ChannelMessageSend(channelID string, content string, options ...discordgo.RequestOption) (*discordgo.Message, error) {
	args := m.Called(channelID, content)
	return args.Get(0).(*discordgo.Message), args.Error(1)
}

func (m *MockDiscordSession) MessageReactionAdd(channelID, messageID, emojiID string, options ...discordgo.RequestOption) error {
	args := m.Called(channelID, messageID, emojiID, options)
	return args.Error(0)
}

func (m *MockDiscordSession) GuildMember(guildID, userID string, options ...discordgo.RequestOption) (*discordgo.Member, error) {
	args := m.Called(guildID, userID, options)
	return args.Get(0).(*discordgo.Member), args.Error(1)
}

func (m *MockDiscordSession) GuildMemberRoleAdd(guildID, userID, roleID string, options ...discordgo.RequestOption) error {
	args := m.Called(guildID, userID, roleID)
	return args.Error(0)
}

func (m *MockDiscordSession) GuildMemberRoleRemove(guildID, userID, roleID string, options ...discordgo.RequestOption) error {
	args := m.Called(guildID, userID, roleID)
	return args.Error(0)
}

func TestNewRoleSelectionHandler(t *testing.T) {
	handler := NewRoleSelectionHandler("guild1", &config.RoleSelectionConfig{})
	assert.NotNil(t, handler)
	assert.Equal(t, "guild1", handler.guildID)
}

func TestReconcileConfig_ChannelNotFound(t *testing.T) {
	mockSession := new(MockDiscordSession)
	handler := NewRoleSelectionHandler("guild1", &config.RoleSelectionConfig{
		Channel: config.RoleSelectionChannel{ID: "channel1"},
	})

	mockSession.On("GuildChannels", "guild1", mock.Anything).Return([]*discordgo.Channel{}, nil)

	err := handler.ReconcileConfig(mockSession)
	assert.NoError(t, err)
}

func TestReconcileConfig_Success(t *testing.T) {
	mockSession := new(MockDiscordSession)
	handler := NewRoleSelectionHandler("guild1", &config.RoleSelectionConfig{
		Channel: config.RoleSelectionChannel{ID: "channel1"},
		Prompts: []config.RoleSelectionPrompt{
			{
				Message: "React to get a role",
				Options: []config.RoleSelectionOption{
					{EmojiID: "👍", RoleID: "role1"},
				},
			},
		},
	})

	mockSession.On("GuildChannels", "guild1", mock.Anything).Return([]*discordgo.Channel{{ID: "channel1"}}, nil)
	mockSession.On("ChannelMessages", "channel1", 100, "", "", "", mock.Anything).Return([]*discordgo.Message{}, nil)
	mockSession.On("ChannelMessageSend", "channel1", "React to get a role").Return(&discordgo.Message{ID: "message1", ChannelID: "channel1"}, nil)
	mockSession.On("MessageReactionAdd", "channel1", "message1", "👍", mock.Anything).Return(nil)

	err := handler.ReconcileConfig(mockSession)
	assert.NoError(t, err)
	mockSession.AssertCalled(t, "ChannelMessageSend", "channel1", "React to get a role")
}

func TestHandleReactionAdd(t *testing.T) {
	mockSession := new(MockDiscordSession)
	handler := NewRoleSelectionHandler("guild1", &config.RoleSelectionConfig{
		Channel: config.RoleSelectionChannel{ID: "channel1"},
		Prompts: []config.RoleSelectionPrompt{
			{
				Message: "React to get a role",
				Options: []config.RoleSelectionOption{
					{EmojiID: "👍", RoleID: "role1"},
				},
			},
		},
	})
	handler.prompts["message1"] = handler.config.Prompts[0]
	handler.initialized = true

	reaction := &discordgo.MessageReactionAdd{
		MessageReaction: &discordgo.MessageReaction{
			ChannelID: "channel1",
			MessageID: "message1",
			Emoji:     discordgo.Emoji{ID: "👍"},
			UserID:    "user1",
			GuildID:   "guild1",
		},
	}

	mockSession.On("GuildMember", "guild1", "user1", mock.Anything).Return(&discordgo.Member{User: &discordgo.User{}, Roles: []string{}}, nil)
	mockSession.On("GuildMemberRoleAdd", "guild1", "user1", "role1").Return(nil)

	handler.HandleReactionAdd(mockSession, reaction)

	mockSession.AssertCalled(t, "GuildMemberRoleAdd", "guild1", "user1", "role1")
}

func TestHandleReactionRemove(t *testing.T) {
	mockSession := new(MockDiscordSession)
	handler := NewRoleSelectionHandler("guild1", &config.RoleSelectionConfig{
		Channel: config.RoleSelectionChannel{ID: "channel1"},
		Prompts: []config.RoleSelectionPrompt{
			{
				Message: "React to get a role",
				Options: []config.RoleSelectionOption{
					{EmojiID: "👍", RoleID: "role1"},
				},
			},
		},
	})
	handler.prompts["message1"] = handler.config.Prompts[0]
	handler.initialized = true

	reaction := &discordgo.MessageReactionRemove{
		MessageReaction: &discordgo.MessageReaction{
			ChannelID: "channel1",
			MessageID: "message1",
			Emoji:     discordgo.Emoji{ID: "👍"},
			UserID:    "user1",
			GuildID:   "guild1",
		},
	}

	mockSession.On("GuildMember", "guild1", "user1", mock.Anything).Return(&discordgo.Member{User: &discordgo.User{}, Roles: []string{"role1"}}, nil)
	mockSession.On("GuildMemberRoleRemove", "guild1", "user1", "role1").Return(nil)

	handler.HandleReactionRemove(mockSession, reaction)

	mockSession.AssertCalled(t, "GuildMemberRoleRemove", "guild1", "user1", "role1")
}

func TestHandleReactionAdd_RoleAlreadyExists(t *testing.T) {
	mockSession := new(MockDiscordSession)
	handler := NewRoleSelectionHandler("guild1", &config.RoleSelectionConfig{
		Channel: config.RoleSelectionChannel{ID: "channel1"},
		Prompts: []config.RoleSelectionPrompt{
			{
				Message: "React to get a role",
				Options: []config.RoleSelectionOption{
					{EmojiID: "👍", RoleID: "role1"},
				},
			},
		},
	})
	handler.prompts["message1"] = handler.config.Prompts[0]
	handler.initialized = true

	reaction := &discordgo.MessageReactionAdd{
		MessageReaction: &discordgo.MessageReaction{
			ChannelID: "channel1",
			MessageID: "message1",
			Emoji:     discordgo.Emoji{ID: "👍"},
			UserID:    "user1",
			GuildID:   "guild1",
		},
	}

	mockSession.On("GuildMember", "guild1", "user1", mock.Anything).Return(&discordgo.Member{User: &discordgo.User{}, Roles: []string{"role1"}}, nil)

	handler.HandleReactionAdd(mockSession, reaction)

	mockSession.AssertNotCalled(t, "GuildMemberRoleAdd", "guild1", "user1", "role1")
}

func TestHandleReactionRemove_RoleNotFound(t *testing.T) {
	mockSession := new(MockDiscordSession)
	handler := NewRoleSelectionHandler("guild1", &config.RoleSelectionConfig{
		Channel: config.RoleSelectionChannel{ID: "channel1"},
		Prompts: []config.RoleSelectionPrompt{
			{
				Message: "React to get a role",
				Options: []config.RoleSelectionOption{
					{EmojiID: "👍", RoleID: "role1"},
				},
			},
		},
	})
	handler.prompts["message1"] = handler.config.Prompts[0]
	handler.initialized = true

	reaction := &discordgo.MessageReactionRemove{
		MessageReaction: &discordgo.MessageReaction{
			ChannelID: "channel1",
			MessageID: "message1",
			Emoji:     discordgo.Emoji{ID: "👍"},
			UserID:    "user1",
			GuildID:   "guild1",
		},
	}

	mockSession.On("GuildMember", "guild1", "user1", mock.Anything).Return(&discordgo.Member{User: &discordgo.User{}, Roles: []string{}}, nil)

	handler.HandleReactionRemove(mockSession, reaction)

	mockSession.AssertNotCalled(t, "GuildMemberRoleRemove", "guild1", "user1", "role1")
}

func TestHandleReactionAdd_GuildMemberError(t *testing.T) {
	mockSession := new(MockDiscordSession)
	handler := NewRoleSelectionHandler("guild1", &config.RoleSelectionConfig{
		Channel: config.RoleSelectionChannel{ID: "channel1"},
		Prompts: []config.RoleSelectionPrompt{
			{
				Message: "React to get a role",
				Options: []config.RoleSelectionOption{
					{EmojiID: "👍", RoleID: "role1"},
				},
			},
		},
	})
	handler.prompts["message1"] = handler.config.Prompts[0]
	handler.initialized = true

	reaction := &discordgo.MessageReactionAdd{
		MessageReaction: &discordgo.MessageReaction{
			ChannelID: "channel1",
			MessageID: "message1",
			Emoji:     discordgo.Emoji{ID: "👍"},
			UserID:    "user1",
			GuildID:   "guild1",
		},
	}

	mockSession.On("GuildMember", "guild1", "user1", mock.Anything).Return((*discordgo.Member)(nil), errors.New("error"))

	handler.HandleReactionAdd(mockSession, reaction)

	mockSession.AssertNotCalled(t, "GuildMemberRoleAdd", "guild1", "user1", "role1")
}
