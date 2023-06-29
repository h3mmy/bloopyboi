package providers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/dustin/go-humanize"
)

var (
	AppCommands = []*discordgo.ApplicationCommand{
		{
			Name:        "inspire",
			Description: "Summons Inspiration",
		},
		{
			Name:        "blissfest",
			Description: "Gets blissfest related information",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionBoolean,
					Name:        "lineup",
					Description: "Try to fetch lineup image",
					Required:    true,
				},
			},
		},
	}
	AppCommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"inspire": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			bttp := GetInspiroClient()
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{
						{
							Author: &discordgo.MessageEmbedAuthor{},
							Image: &discordgo.MessageEmbedImage{
								URL: bttp.GetInspiroImageURL(),
							},
						},
					},
				},
			})
		},
		"blissfest": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			getLineUp := false
			// Access options in the order provided by the user.
			options := i.ApplicationCommandData().Options
			for _, opt := range options {
				if opt.Name == "lineup" {
					getLineUp = opt.BoolValue()
				}
			}

			bsvc := GetBlissfestService()
			var resData discordgo.InteractionResponseData

			if getLineUp {
				resData = discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{
						{
							Author: &discordgo.MessageEmbedAuthor{},
							Image: &discordgo.MessageEmbedImage{
								URL: bsvc.GetLineupImageURI(),
							},
						},
					},
					Title:   "Blissfest",
					// pending https://github.com/dustin/go-humanize/pull/92
					// Content: fmt.Sprintf("%s left", humanize.Time(bsvc.GetTimeUntilStart(nil))),
					Content: fmt.Sprintf("%s left", humanize.Time(*bsvc.GetStartTime())),
				}


			} else {
				resData = discordgo.InteractionResponseData{
					Title:   "Blissfest",
					Content: fmt.Sprintf("%s left", humanize.Time(*bsvc.GetStartTime())),
				}
			}

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &resData,
			})
		},
	}
)
