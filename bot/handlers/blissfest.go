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
	// GuildID for which this command will be active
	// For global commands, set to ""
	guildId string
	roles   []int64
}

func NewBlissfestCommand(svc *services.BlissfestService) *BlissfestCommand {
	return &BlissfestCommand{
		meta:        models.NewBloopyMeta(),
		Name:        "blissfest",
		Description: "Gets blissfest related information",
		logger:      log.NewZapLogger().Named("blissfest_command"),
		blissSvc:    svc,
		guildId:     "",
		roles:       []int64{},
	}
}

func (p *BlissfestCommand) WithGuild(guildId string) *BlissfestCommand {
	p.logger.Debug("setting guild", zap.String("guildId", guildId))
	p.guildId = guildId
	return p
}

func (p *BlissfestCommand) WithRoles(roles ...int64) *BlissfestCommand {
	p.roles = roles
	return p
}

func (p *BlissfestCommand) GetAllowedRoles() []int64 {
	return p.roles
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
		p.logger.Debug("received interaction", zap.String("interactionID", i.ID), zap.String("username", GetDiscordUserFromInteraction(i).Username))
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

		resEmbeds := []*discordgo.MessageEmbed{}

		adultWeekendPriceLevel, err := bsvc.GetAdultWeekendPriceLevel()
		var adultWeekendPriceLevelEmbed *discordgo.MessageEmbed
		if err != nil {
			p.logger.Warn("error getting adult weekend price level. Not including in response", zap.Error(err))
		} else {
			adultWeekendPriceLevelEmbed = &discordgo.MessageEmbed{
				Title: "Adult Weekend (18+) Ticket Info",
				Fields: []*discordgo.MessageEmbedField{
					{
						Name:  "Active",
						Value: fmt.Sprintf("%t", adultWeekendPriceLevel.Active == "1"),
					},
					{
						Name:  "Price",
						Value: adultWeekendPriceLevel.Price, //fmt.Sprintf("%.2f",adultWeekendPriceLevel.Price),
					},
					{
						Name:  "Transaction Limit",
						Value: adultWeekendPriceLevel.TransactionLimit, //fmt.Sprintf("%d", adultWeekendPriceLevel.TransactionLimit),
					},
				},
			}
			resEmbeds = append(resEmbeds, adultWeekendPriceLevelEmbed)
		}

		if getLineUp {
			resEmbeds = append(resEmbeds, &discordgo.MessageEmbed{
				Title:  fmt.Sprintf("%d Blissfest Lineup", bsvc.GetStartTime().Year()),
				Author: &discordgo.MessageEmbedAuthor{},
				Image: &discordgo.MessageEmbedImage{
					URL: bsvc.GetLineupImageURI(),
				},
			})
		}
		if len(resEmbeds) > 0 {
			resData = discordgo.InteractionResponseData{
				Embeds: resEmbeds,
				Title:  "Blissfest",
				// pending https://github.com/dustin/go-humanize/pull/92
				// Content: fmt.Sprintf("%s left", humanize.Time(bsvc.GetTimeUntilStart(nil))),
				Content: fmt.Sprintf("blissfest starts %s", humanize.Time(*bsvc.GetStartTime())),
			}
		} else {
			resData = discordgo.InteractionResponseData{
				Title:   "Blissfest",
				// pending https://github.com/dustin/go-humanize/pull/92
				// Content: fmt.Sprintf("%s left", humanize.Time(bsvc.GetTimeUntilStart(nil))),
				Content: fmt.Sprintf("blissfest start %s", humanize.Time(*bsvc.GetStartTime())),
			}
		}
		p.logger.Debug("finished constructing response", zap.Bool("getLineup", getLineUp))

		err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
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

func (p *BlissfestCommand) GetGuildID() string {
	return p.guildId
}
