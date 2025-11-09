package handlers

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/h3mmy/bloopyboi/bot/services"
	"github.com/h3mmy/bloopyboi/internal/models"
	log "github.com/h3mmy/bloopyboi/pkg/logs"
	"go.uber.org/zap"
)

// InspiroCommand is a command that summons inspiration.
type InspiroCommand struct {
	meta        models.BloopyMeta
	Name        string
	Description string
	logger      *zap.Logger
	inspiroSvc  *services.InspiroService
}

// NewInspiroCommand creates a new InspiroCommand.
func NewInspiroCommand(svc *services.InspiroService) *InspiroCommand {
	return &InspiroCommand{
		meta:        models.NewBloopyMeta(),
		Name:        "inspire",
		Description: "Summons Inspiration",
		inspiroSvc:  svc,
		logger:      log.NewZapLogger().Named("inspiro_command"),
	}
}

// GetAppCommand returns the application command for the Inspiro command.
func (p *InspiroCommand) GetAppCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        strings.ToLower(p.Name),
		Description: p.Description,
	}
}

// GetAppCommandHandler returns the handler for the Inspiro command.
func (p *InspiroCommand) GetAppCommandHandler() func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		err := s.InteractionRespond(i.Interaction, p.inspiroSvc.CreateInteractionResponse())
		if err != nil {
			p.logger.Error("Failed to respond to interaction", zap.Error(err), zap.String("command", "inspire"))
		}
	}
}

// GetMessageComponentHandlers returns the message component handlers for the Inspiro command.
func (p *InspiroCommand) GetMessageComponentHandlers() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return nil
}

// GetGuildID returns the guild ID for the command.
func (p *InspiroCommand) GetGuildID() string {
	// Is global command
	return ""
}

// GetAllowedRoles returns the allowed roles for the command.
func (p *InspiroCommand) GetAllowedRoles() []int64 {
	return []int64{}
}
