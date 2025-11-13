package handlers

import (
	"context"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/h3mmy/bloopyboi/bot/services"
	"github.com/h3mmy/bloopyboi/internal/models"
	log "github.com/h3mmy/bloopyboi/pkg/logs"
	"go.uber.org/zap"
)

// UserRequestCommand is a command that allows users to see their media requests.
// TODO: This is an experimental command and should be refactored.
type UserRequestCommand struct {
	meta        models.BloopyMeta
	Name        string
	Description string
	logger      *zap.Logger
	bookSvc     *services.BookService
	guildId     string
	roles       []int64
}

// NewUserRequestCommand creates a new UserRequestCommand.
func NewUserRequestCommand(bookSvc *services.BookService) *UserRequestCommand {
	return &UserRequestCommand{
		meta:        models.NewBloopyMeta(),
		Name:        string(Requests),
		Description: "(Xperimental) Get your requests",
		bookSvc:     bookSvc,
		logger:      log.NewZapLogger().Named("requests_command"),
		guildId:     "",
		roles:       []int64{},
	}
}

// WithGuild sets the guild ID for the command.
func (c *UserRequestCommand) WithGuild(guildId string) *UserRequestCommand {
	c.guildId = guildId
	return c
}

// WithRoles sets the allowed roles for the command.
func (c *UserRequestCommand) WithRoles(roles ...int64) *UserRequestCommand {
	c.roles = roles
	return c
}

// GetGuildID returns the guild ID for the command.
func (c *UserRequestCommand) GetGuildID() string {
	return c.guildId
}

// GetAllowedRoles returns the allowed roles for the command.
func (c *UserRequestCommand) GetAllowedRoles() []int64 {
	return c.roles
}

// GetAppCommand returns the application command for the UserRequest command.
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

// GetAppCommandHandler returns the handler for the UserRequest command.
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
					bookEmbeds := []*discordgo.MessageEmbed{}
					embedLimit := 3
					for i, req := range allBookReqs {
						if i < embedLimit {
							bookEmbeds = append(bookEmbeds, c.bookSvc.BuildBookRequestStatusAsEmbed(context.TODO(), req))
						} else {
							break
						}
					}
					err = s.InteractionRespond(i.Interaction,
						&discordgo.InteractionResponse{
							Type: discordgo.InteractionResponseChannelMessageWithSource,
							Data: &discordgo.InteractionResponseData{
								Content: fmt.Sprintf("You have %d total book requests. I can only show %d at a time", len(allBookReqs), embedLimit),
								Flags:   discordgo.MessageFlagsEphemeral, // only show to user who requested it
								Embeds:  bookEmbeds,
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

// GetMessageComponentHandlers returns the message component handlers for the UserRequest command.
func (c *UserRequestCommand) GetMessageComponentHandlers() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return nil
}
