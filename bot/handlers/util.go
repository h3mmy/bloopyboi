package handlers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/h3mmy/bloopyboi/ent"
	"github.com/h3mmy/bloopyboi/internal/models"
)

// TODO: This function is commented out and should be removed or implemented.
// typeInChannel sets the typing indicator for a channel. The indicator is cleared
// when a message is sent.
// func typeInChannel(channel chan bool, s *discordgo.Session, channelID string) {
// 	for {
// 		select {
// 		case <-channel:
// 			return
// 		default:
// 			if err := s.ChannelTyping(channelID); err != nil {
// 				fmt.Println("unable to set typing indicator: ", err)
// 			}
// 			time.Sleep(time.Second * 5)
// 		}
// 	}
// }

// GetDiscordUserFromInteraction returns the user from an interaction.
func GetDiscordUserFromInteraction(i *discordgo.InteractionCreate) *discordgo.User {
	if i.User != nil {
		// This field is only filled when the slash command was invoked in a DM
		// if it was invoked in a guild, the `Member` field will be filled instead
		return i.User
	} else {
		return i.Member.User
	}
}

// GetBookRequestsAsEmbeds returns a slice of embeds for a slice of book requests.
func GetBookRequestsAsEmbeds(requests []*ent.MediaRequest) []*discordgo.MessageEmbed {
	var embeds []*discordgo.MessageEmbed
	for _, request := range requests {
		book, err := request.Edges.BookOrErr()
		if err != nil {
			continue
		}
		var colorCode models.ColorCode
		switch request.Status {
		case "Pending":
			colorCode = models.ColorCodeInfo
		case "Approved":
			colorCode = models.ColorCodeSuccess
		}
		embed := &discordgo.MessageEmbed{
			Title:       book.Title,
			Description: book.Description,
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:  "Status",
					Value: string(request.Status),
				},
				{
					Name:  "Requested",
					Value: request.CreateTime.Format("2006-01-02 15:04:05"),
				},
			},
			Color: int(colorCode),
		}
		embeds = append(embeds, embed)
	}
	return embeds
}
