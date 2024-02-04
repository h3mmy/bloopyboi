package handlers

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/dustin/go-humanize"
	"github.com/h3mmy/bloopyboi/bot/internal/log"
	"github.com/h3mmy/bloopyboi/bot/internal/models"
	"github.com/h3mmy/bloopyboi/bot/services"
	"go.uber.org/zap"
)

type BlissfestCommand struct {
	meta        models.BloopyMeta
	Name        string
	Description string
	logger      *zap.Logger
	blissSvc    *services.BlissfestService
}

func NewBlissfestCommand(svc *services.BlissfestService) *BlissfestCommand {
	return &BlissfestCommand{
		meta:        models.NewBloopyMeta(),
		Name:        "blissfest",
		Description: "Gets blissfest related information",
		logger:      log.NewZapLogger().Named("blissfest_command"),
		blissSvc:    svc,
	}
}

func (p *BlissfestCommand) GetAppCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        strings.ToLower(p.Name),
		Description: p.Description,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionBoolean,
				Name:        "lineup",
				Description: "Try to fetch lineup image",
				Required:    true,
			},
		},
	}
}

func (p *BlissfestCommand) GetAppCommandHandler() func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		getLineUp := false
		// Access options in the order provided by the user.
		options := i.ApplicationCommandData().Options
		for _, opt := range options {
			if opt.Name == "lineup" {
				getLineUp = opt.BoolValue()
			}
		}

		bsvc := p.blissSvc
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
				Title: "Blissfest",
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

		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &resData,
		})
		if err != nil {
			p.logger.Error("error responding to interaction", zap.Error(err))
		}
	}
}

func (p *BlissfestCommand) GetMessageComponentHandlers() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return nil
}
