package handlers

import (
	"context"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/h3mmy/bloopyboi/bot/internal/log"
	"github.com/h3mmy/bloopyboi/bot/internal/models"
	"github.com/h3mmy/bloopyboi/bot/services"
	"go.uber.org/zap"
)

type UserRequestCommand struct {
	meta        models.BloopyMeta
	Name        string
	Description string
	logger      *zap.Logger
	bookSvc     *services.BookService
}

func NewUserRequestCommand(bookSvc *services.BookService) *UserRequestCommand {
	return &UserRequestCommand{
		meta:        models.NewBloopyMeta(),
		Name:        "requests",
		Description: "(Xperimental) Get your requests",
		bookSvc:     bookSvc,
		logger:      log.NewZapLogger().Named("requests_command"),
	}
}

func (c *UserRequestCommand) GetAppCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        c.Name,
		Description: c.Description,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "media_type",
				Description: "Type of media",
				Type:        discordgo.ApplicationCommandOptionString,
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  "Books",
						Value: "book",
					},
				},
				Required: true,
			},
		},
	}
}

func (c *UserRequestCommand) GetAppCommandHandler() func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		c.logger.Info("handling user request command")

		// Options in the order provided by the user.
		options := i.ApplicationCommandData().Options
		// Convert the slice into a map
		optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
		for _, opt := range options {
			optionMap[opt.Name] = opt
		}
		if mediaType, ok := optionMap["media_type"]; ok {
			if mediaType.StringValue() == "book" {
				discordUser := GetDiscordUserFromInteraction(i)
				allBookReqs, err := c.bookSvc.GetAllBookRequestsForUser(context.Background(), discordUser.ID)
				if err != nil {
					c.logger.Error("error getting book requests", zap.Error(err))
					// responsd to user
					err = s.InteractionRespond(i.Interaction,
						&discordgo.InteractionResponse{
							Type: discordgo.InteractionResponseChannelMessageWithSource,
							Data: &discordgo.InteractionResponseData{
								Content: fmt.Sprintf("... So there was a problem, but I can't read the note: `%s`", err.Error()),
							},
						},
					)
					if err != nil {
						c.logger.Error("error responding to interaction", zap.Error(err))
					}
				} else {
					c.logger.Info("got book requests for user", zap.Int("count", len(allBookReqs)), zap.String("username", discordUser.Username))
					err = s.InteractionRespond(i.Interaction,
						&discordgo.InteractionResponse{
							Type: discordgo.InteractionResponseChannelMessageWithSource,
							Data: &discordgo.InteractionResponseData{
								Content: fmt.Sprintf("Here are your book requests: %s", allBookReqs),
								Flags:   discordgo.MessageFlagsEphemeral, // only show to user who requested it
								Embeds: nil,
							},
						})
					if err != nil {
						c.logger.Error("error responding to interaction", zap.Error(err))
					}
				}

			}
		}
	}
}
