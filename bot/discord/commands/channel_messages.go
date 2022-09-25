package commands

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"gitlab.com/h3mmy/bloopyboi/bot/services"
)

var (
	textResponseMap = map[string]string{"pong": "Ping!", "Pong!": "-_-"}
)

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.ToLower(m.Content) == "inspire" {
		logger.Debug(
			fmt.Sprintf(
				"Received Inspiration Request from %s with ID %s",
				m.Author.Username,
				m.Author.ID),
			)
		// Start typing indicator
		typingStop := make(chan bool, 1)
		go typeInChannel(typingStop, s, m.ChannelID)
		bttp := services.NewInspiroClient()
		embed := &discordgo.MessageEmbed{
			Author: &discordgo.MessageEmbedAuthor{},
			Image: &discordgo.MessageEmbedImage{
				URL: bttp.GetInspiroImageURL(),
			},
		}
		typingStop <- true
		s.ChannelMessageSendEmbed(m.ChannelID, embed)
	}

	resp, ok := textResponseMap[m.Content]
	if !ok {
		// Means nothing stored for canned Response
		return
	}
	logger.Debug(
		fmt.Sprintf(
			"Received Message from %s with ID %s",
			m.Author.Username,
			m.Author.ID),
		)
	// Send with typing indicators
	typingStop := make(chan bool, 1)
	go typeInChannel(typingStop, s, m.ChannelID)
	s.ChannelMessageSend(m.ChannelID, resp)
	typingStop <- true
}
