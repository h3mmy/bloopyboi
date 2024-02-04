package handlers

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/h3mmy/bloopyboi/bot/internal/models"
	"github.com/h3mmy/bloopyboi/bot/services"
	"go.uber.org/zap"
)

type InspiroCommand struct {
	meta        models.BloopyMeta
	Name        string
	Description string
	logger      *zap.Logger
	inspiroSvc  *services.InspiroClient
}

func NewInspiroCommand(svc *services.InspiroClient) *InspiroCommand {
	return &InspiroCommand{
		meta:        models.NewBloopyMeta(),
		Name:        "inspire",
		Description: "Summons Inspiration",
		inspiroSvc:  svc,
		logger:      zap.L().Named("InspiroCommand"),
	}
}

func (p *InspiroCommand) GetAppCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        strings.ToLower(p.Name),
		Description: p.Description,
	}
}

func (p *InspiroCommand) GetAppCommandHandler() func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {

		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{
					{
						Author: &discordgo.MessageEmbedAuthor{},
						Image: &discordgo.MessageEmbedImage{
							URL: p.inspiroSvc.GetInspiroImageURL(),
						},
					},
				},
			},
		})
		if err != nil {
			p.logger.Error("Failed to respond to interaction", zap.Error(err), zap.String("command", "inspire"))
		}
	}
}

func CreateInsprioEmbed(svc *services.InspiroClient) *discordgo.MessageEmbed {

	return &discordgo.MessageEmbed{

			Author: &discordgo.MessageEmbedAuthor{},
			Image: &discordgo.MessageEmbedImage{
				URL: svc.GetInspiroImageURL(),
			},
		
	}

}
