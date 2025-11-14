package handlers

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/h3mmy/bloopyboi/bot/services"
	"github.com/h3mmy/bloopyboi/internal/models"
	log "github.com/h3mmy/bloopyboi/pkg/logs"
	"go.uber.org/zap"
)

type InspiroCommand struct {
	meta        models.BloopyMeta
	Name        string
	Description string
	logger      *zap.Logger
	inspiroSvc  *services.InspiroService
}

func NewInspiroCommand(svc *services.InspiroService) *InspiroCommand {
	return &InspiroCommand{
		meta:        models.NewBloopyMeta(),
		Name:        string(Inspiro),
		Description: "Summons Inspiration",
		inspiroSvc:  svc,
		logger:      log.NewZapLogger().Named("inspiro_command"),
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
		err := s.InteractionRespond(i.Interaction, p.inspiroSvc.CreateInteractionResponse())
		if err != nil {
			p.logger.Error("Failed to respond to interaction", zap.Error(err), zap.String("command", "inspire"))
		}
	}
}

func (p *InspiroCommand) GetMessageComponentHandlers() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return nil
}

func (p *InspiroCommand) GetGuildID() string {
	// Is global command
	return ""
}

func (p *InspiroCommand) GetAllowedRoles() []int64 {
	return []int64{}
}
