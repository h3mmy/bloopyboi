package commands

import (
	"strings"
	"github.com/bwmarrin/discordgo"
	"gitlab.com/h3mmy/bloopyboi/bot/util"
)

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}


	// If the message is "ping" reply with "Pong!"
	if strings.ToLower(m.Content) == "inspire" {
		// Start typing indicator
		typingStop := make(chan bool, 1)
		go typeInChannel(typingStop, s, m.ChannelID)
		bttp := util.NewBloopyClient()
		embed := &discordgo.MessageEmbed{
			Author: &discordgo.MessageEmbedAuthor{},
			Image: &discordgo.MessageEmbedImage{
				URL: bttp.Inspiro_api.GetInspiro(),
			},
		}
		typingStop <- true
		s.ChannelMessageSendEmbed(m.ChannelID, embed)
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "pong" {
		// Start typing indicator
		typingStop := make(chan bool, 1)
		go typeInChannel(typingStop, s, m.ChannelID)
		s.ChannelMessageSend(m.ChannelID, "Ping!")
		typingStop <- true
	}

	if m.Content == "Pong!" {
		// Start typing indicator
		typingStop := make(chan bool, 1)
		go typeInChannel(typingStop, s, m.ChannelID)
		s.ChannelMessageSend(m.ChannelID, "-_-")
		typingStop <- true
	}
}
