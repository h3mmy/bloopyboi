package handlers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/h3mmy/bloopyboi/internal/models"
	"github.com/h3mmy/bloopyboi/ent"
)

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

func GetDiscordUserFromInteraction(i *discordgo.InteractionCreate) *discordgo.User {
	if i.User != nil {
		// This field is only filled when the slash command was invoked in a DM
		// if it was invoked in a guild, the `Member` field will be filled instead
		return i.User
	} else {
		return i.Member.User
	}
}

func GetBookRequestsAsEmbeds(requests []*ent.MediaRequest) []*discordgo.MessageEmbed {
	var embeds []*discordgo.MessageEmbed
	for _, request := range requests {
		books, err := request.Edges.BookOrErr()
		if err != nil || len(books) == 0 {
			continue
		}
		book := books[0]
		var colorCode models.ColorCode
		if request.Status == "Pending" {
			colorCode = models.ColorCodeInfo
		} else if request.Status == "Approved" {
			colorCode = models.ColorCodeSuccess
		}
		embed := &discordgo.MessageEmbed{
			Title:       book.Title,
			Description: book.Description,
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:  "Status",
					Value: request.Status,
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
