package services

import (
	"github.com/bwmarrin/discordgo"
	"github.com/h3mmy/bloopyboi/bot/internal/log"
	"github.com/h3mmy/bloopyboi/bot/internal/models"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type DiscordService struct {
	meta           models.BloopyMeta
	logger         *zap.Logger
	discordSession *discordgo.Session
	// The handlers registered with this service, keyed by the command name
	handlerRegistry map[string]models.DiscordAppCommand
	// The commands registered with discord that will need to be de-registered on shutdown
	commandRegistry map[string]*discordgo.ApplicationCommand
}

func NewDiscordService(session *discordgo.Session) *DiscordService {
	lgr := log.NewZapLogger().
		Named("discord_service").
		With(zapcore.Field{
			Key:    ServiceLoggerFieldKey,
			Type:   zapcore.StringType,
			String: "discord",
		})
	return &DiscordService{
		meta:           models.NewBloopyMeta(),
		logger:         lgr,
		discordSession: session,
	}
}

func (d *DiscordService) GetMeta() models.BloopyMeta {
	return d.meta
}

// Primarily for backwards compatibility while I move things into a service
func (d *DiscordService) GetSession() *discordgo.Session {
	return d.discordSession
}

func (d *DiscordService) AddHandler(handler interface{}) func() {
	return d.discordSession.AddHandler(handler)
}

func (d *DiscordService) GetDataReady() bool {
	return d.discordSession.DataReady
}

func (d *DiscordService) RegisterAppCommand(command models.DiscordAppCommand) (*discordgo.ApplicationCommand, error) {
	d.handlerRegistry[command.GetAppCommand().Name] = command
	cmd, err := d.discordSession.ApplicationCommandCreate(d.discordSession.State.User.ID, "", command.GetAppCommand())
	if err != nil {
		d.logger.Error("error registering app command")
		return nil, err
	}
	d.commandRegistry[command.GetAppCommand().Name] = cmd
	return cmd, nil
}

// Proxies InteractionCreate events to the handlers in the svc handler registry
func (d *DiscordService) AddInteractionHandlerProxy() {
	d.discordSession.AddHandler(
		func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			switch i.Type {
			case discordgo.InteractionApplicationCommand:
				if h, ok := d.handlerRegistry[i.ApplicationCommandData().Name]; ok {
					h.GetAppCommandHandler()(s, i)
				} else {
					logger.Info("no handler registered for command", zap.String("command", i.ApplicationCommandData().Name))
				}
			case discordgo.InteractionMessageComponent:
				if h, ok := d.handlerRegistry[i.MessageComponentData().CustomID]; ok {
					h.GetAppCommandHandler()(s, i)
				}else {
					logger.Info("no handler registered for message component", zap.String("customID", i.MessageComponentData().CustomID))
				}
			case discordgo.InteractionModalSubmit:
				if h, ok := d.handlerRegistry[i.ModalSubmitData().CustomID]; ok {
					h.GetAppCommandHandler()(s, i)
				}else {
					logger.Info("no handler registered for modal submit data", zap.String("customID", i.ModalSubmitData().CustomID))
				}
			}
		})
}
