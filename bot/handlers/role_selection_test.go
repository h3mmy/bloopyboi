package handlers

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/h3mmy/bloopyboi/pkg/config"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestReconcileConfig_Create(t *testing.T) {
	// Test setup
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	session, err := discordgo.New("Bot testToken")
	assert.NoError(t, err)

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
	httpmock.RegisterResponder("GET", "https://discord.com/api/v9/channels/testChannel",
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewJsonResponse(200, discordgo.Channel{ID: "testChannel"})
		},
	)
	httpmock.RegisterResponder("GET", "https://discord.com/api/v9/channels/testChannel/messages?limit=2",
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewJsonResponse(200, []*discordgo.Message{})
		},
	)
	httpmock.RegisterResponder("POST", "https://discord.com/api/v9/channels/testChannel/messages",
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewJsonResponse(200, discordgo.Message{ID: "testMessage", ChannelID: "testChannel"})
		},
	)
	httpmock.RegisterResponder("PUT", "https://discord.com/api/v9/channels/testChannel/messages/testMessage/reactions/testEmoji/@me",
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewJsonResponse(204, nil)
		},
	)

	// Call the function
	err = roleSelectionHandler.ReconcileConfig(session)
	fmt.Println(httpmock.GetCallCountInfo())

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, 4, httpmock.GetTotalCallCount())
}

func TestReconcileConfig_Update(t *testing.T) {
	// Test setup
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	session, err := discordgo.New("Bot testToken")
	assert.NoError(t, err)

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
	httpmock.RegisterResponder("GET", "https://discord.com/api/v9/channels/testChannel",
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewJsonResponse(200, discordgo.Channel{ID: "testChannel"})
		},
	)
	httpmock.RegisterResponder("GET", "https://discord.com/api/v9/channels/testChannel/messages?limit=2",
		func(req *http.Request) (*http.Response, error) {
			messages := []*discordgo.Message{
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
			}
			return httpmock.NewJsonResponse(200, messages)
		},
	)
	httpmock.RegisterResponder("PATCH", "https://discord.com/api/v9/channels/testChannel/messages/testMessage",
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewJsonResponse(200, discordgo.Message{})
		},
	)
	httpmock.RegisterResponder("PUT", "https://discord.com/api/v9/channels/testChannel/messages/testMessage/reactions/testEmoji/@me",
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewJsonResponse(204, nil)
		},
	)

	// Call the function
	err = roleSelectionHandler.ReconcileConfig(session)
	fmt.Println(httpmock.GetCallCountInfo())

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, 4, httpmock.GetTotalCallCount())
}

func TestReconcileConfig_Delete(t *testing.T) {
	// Test setup
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	session, err := discordgo.New("Bot testToken")
	assert.NoError(t, err)

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
	httpmock.RegisterResponder("GET", "https://discord.com/api/v9/channels/testChannel",
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewJsonResponse(200, discordgo.Channel{ID: "testChannel"})
		},
	)
	httpmock.RegisterResponder("GET", "https://discord.com/api/v9/channels/testChannel/messages",
		func(req *http.Request) (*http.Response, error) {
			messages := []*discordgo.Message{
				{
					ID:        "testMessage",
					ChannelID: "testChannel",
					Embeds: []*discordgo.MessageEmbed{
						{
							Title: "Test Prompt",
						},
					},
				},
			}
			return httpmock.NewJsonResponse(200, messages)
		},
	)
	httpmock.RegisterResponder("DELETE", "https://discord.com/api/v9/channels/testChannel/messages/testMessage",
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewJsonResponse(204, nil)
		},
	)

	// Call the function
	err = roleSelectionHandler.ReconcileConfig(session)
	fmt.Println(httpmock.GetCallCountInfo())

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, 3, httpmock.GetTotalCallCount())
}
