package commands

import (
	"github.com/bwmarrin/discordgo"
	"gitlab.com/h3mmy/bloopyboi/bot/services"
)

var (
	Commands = []*discordgo.ApplicationCommand{
		{
			Name: "inspire",
			Description: "Summons Inspiration",
		},
	}
	CommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"inspire": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			bttp := services.NewInspiroClient()
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{
						{
							Author: &discordgo.MessageEmbedAuthor{},
							Image: &discordgo.MessageEmbedImage{
								URL: bttp.Inspiro_api.GetInspiro(),
							},
						},
					},
				},
			})
		},
	}
)
