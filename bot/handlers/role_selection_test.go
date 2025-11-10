package handlers

import (
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/h3mmy/bloopyboi/pkg/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockDiscordSession is a mock of the discordgo.Session.
type MockDiscordSession struct {
	mock.Mock
}

func (m *MockDiscordSession) Channel(channelID string) (*discordgo.Channel, error) {
	args := m.Called(channelID)
	return args.Get(0).(*discordgo.Channel), args.Error(1)
}

func (m *MockDiscordSession) ChannelMessages(channelID string, limit int, beforeID, afterID, aroundID string) ([]*discordgo.Message, error) {
	args := m.Called(channelID, limit, beforeID, afterID, aroundID)
	return args.Get(0).([]*discordgo.Message), args.Error(1)
}

func (m *MockDiscordSession) ChannelMessageSendComplex(channelID string, data *discordgo.MessageSend) (*discordgo.Message, error) {
	args := m.Called(channelID, data)
	return args.Get(0).(*discordgo.Message), args.Error(1)
}

func (m *MockDiscordSession) ChannelMessageEditEmbeds(channelID, messageID string, embeds []*discordgo.MessageEmbed) (*discordgo.Message, error) {
	args := m.Called(channelID, messageID, embeds)
	return args.Get(0).(*discordgo.Message), args.Error(1)
}

func (m *MockDiscordSession) ChannelMessageDelete(channelID, messageID string) error {
	args := m.Called(channelID, messageID)
	return args.Error(0)
}

func (m *MockDiscordSession) MessageReactionAdd(channelID, messageID, emojiID string) error {
	args := m.Called(channelID, messageID, emojiID)
	return args.Error(0)
}

func (m *MockDiscordSession) GuildMemberRoleAdd(guildID, userID, roleID string) error {
	args := m.Called(guildID, userID, roleID)
	return args.Error(0)
}

func (m *MockDiscordSession) GuildMemberRoleRemove(guildID, userID, roleID string) error {
	args := m.Called(guildID, userID, roleID)
	return args.Error(0)
}

func (m *MockDiscordSession) GuildMember(guildID, userID string) (*discordgo.Member, error) {
	args := m.Called(guildID, userID)
	return args.Get(0).(*discordgo.Member), args.Error(1)
}

func TestReconcileConfig_Create(t *testing.T) {
	// Test setup
	mockSession := new(MockDiscordSession)
	roleSelectionHandler := NewRoleSelectionHandler("testGuild", &config.RoleSelectionConfig{
		Channel: struct {
			Name string `mapstructure:"name"`
			ID   string `mapstructure:"id"`
		}{
			ID: "testChannel",
		},
		Prompts: []config.RoleSelectionPrompt{
			{
				Message: "Test Prompt",
				Options: []struct {
					EmojiID     string `mapstructure:"emojiID"`
					Description string `mapstructure:"description"`
					RoleID      string `mapstructure:"roleId"`
				}{
					{
						RoleID:      "testRole",
						EmojiID:     "testEmoji",
						Description: "Test Description",
					},
				},
			},
		},
	})

	// Mock responses
	mockSession.On("Channel", "testChannel").Return(&discordgo.Channel{ID: "testChannel"}, nil)
	mockSession.On("ChannelMessages", "testChannel", 2, "", "", "").Return([]*discordgo.Message{}, nil)
	mockSession.On("ChannelMessageSendComplex", "testChannel", mock.Anything).Return(&discordgo.Message{ID: "testMessage", ChannelID: "testChannel"}, nil)
	mockSession.On("MessageReactionAdd", "testChannel", "testMessage", "testEmoji").Return(nil)

	// Call the function
	err := roleSelectionHandler.ReconcileConfig(mockSession)

	// Assertions
	assert.NoError(t, err)
	mockSession.AssertCalled(t, "ChannelMessageSendComplex", "testChannel", mock.Anything)
}

func TestReconcileConfig_Update(t *testing.T) {
	// Test setup
	mockSession := new(MockDiscordSession)
	roleSelectionHandler := NewRoleSelectionHandler("testGuild", &config.RoleSelectionConfig{
		Channel: struct {
			Name string `mapstructure:"name"`
			ID   string `mapstructure:"id"`
		}{
			ID: "testChannel",
		},
		Prompts: []config.RoleSelectionPrompt{
			{
				Message: "Test Prompt",
				Options: []struct {
					EmojiID     string `mapstructure:"emojiID"`
					Description string `mapstructure:"description"`
					RoleID      string `mapstructure:"roleId"`
				}{
					{
						RoleID:      "testRole",
						EmojiID:     "testEmoji",
						Description: "Test Description Updated",
					},
				},
			},
		},
	})

	// Mock responses
	mockSession.On("Channel", "testChannel").Return(&discordgo.Channel{ID: "testChannel"}, nil)
	mockSession.On("ChannelMessages", "testChannel", 2, "", "", "").Return([]*discordgo.Message{
		{
			ID:        "testMessage",
			ChannelID: "testChannel",
			Embeds: []*discordgo.MessageEmbed{
				{
					Title: "Test Prompt",
					Fields: []*discordgo.MessageEmbedField{
						{
							Name:  "testEmoji",
							Value: "<@&testRole> - Test Description",
						},
					},
				},
			},
		},
	}, nil)
	mockSession.On("ChannelMessageEditEmbeds", "testChannel", "testMessage", mock.AnythingOfType("[]*discordgo.MessageEmbed")).Return(&discordgo.Message{}, nil)
	mockSession.On("MessageReactionAdd", "testChannel", "testMessage", "testEmoji").Return(nil)

	// Call the function
	err := roleSelectionHandler.ReconcileConfig(mockSession)

	// Assertions
	assert.NoError(t, err)
	mockSession.AssertCalled(t, "ChannelMessageEditEmbeds", "testChannel", "testMessage", mock.AnythingOfType("[]*discordgo.MessageEmbed"))
}

func TestReconcileConfig_Delete(t *testing.T) {
	// Test setup
	mockSession := new(MockDiscordSession)
	roleSelectionHandler := NewRoleSelectionHandler("testGuild", &config.RoleSelectionConfig{
		Channel: struct {
			Name string `mapstructure:"name"`
			ID   string `mapstructure:"id"`
		}{
			ID: "testChannel",
		},
		Prompts: []config.RoleSelectionPrompt{},
	})

	// Mock responses
	mockSession.On("Channel", "testChannel").Return(&discordgo.Channel{ID: "testChannel"}, nil)
	mockSession.On("ChannelMessages", "testChannel", 0, "", "", "").Return([]*discordgo.Message{
		{
			ID: "testMessage",
			Embeds: []*discordgo.MessageEmbed{
				{
					Title: "Test Prompt",
				},
			},
		},
	}, nil)
	mockSession.On("ChannelMessageDelete", "testChannel", "testMessage").Return(nil)

	// Call the function
	err := roleSelectionHandler.ReconcileConfig(mockSession)

	// Assertions
	assert.NoError(t, err)
	mockSession.AssertCalled(t, "ChannelMessageDelete", "testChannel", "testMessage")
}
