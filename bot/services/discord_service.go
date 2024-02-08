package services

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/h3mmy/bloopyboi/bot/internal/database"
	"github.com/h3mmy/bloopyboi/bot/internal/log"
	"github.com/h3mmy/bloopyboi/bot/internal/models"
	"github.com/h3mmy/bloopyboi/ent"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type DiscordService struct {
	meta           models.BloopyMeta
	logger         *zap.Logger
	discordSession *discordgo.Session
	// The handlers registered with this service, keyed by the command name
	handlerRegistry map[string]func(*discordgo.Session, *discordgo.InteractionCreate)
	// The commands registered with discord that will need to be de-registered on shutdown
	commandRegistry map[string]*discordgo.ApplicationCommand
	db              *ent.Client
	dbEnabled       bool
}

func NewDiscordService(session *discordgo.Session) *DiscordService {
	dbEnabled := true
	lgr := log.NewZapLogger().
		Named("discord_service").
		With(zapcore.Field{
			Key:    ServiceLoggerFieldKey,
			Type:   zapcore.StringType,
			String: "discord",
		})
	dbClient, err := database.Open()
	if err != nil {
		lgr.Error("failed to open database", zap.Error(err))
		dbEnabled = false
	}
	return &DiscordService{
		meta:            models.NewBloopyMeta(),
		logger:          lgr,
		discordSession:  session,
		dbEnabled:       dbEnabled,
		db:              dbClient,
		handlerRegistry: make(map[string]func(*discordgo.Session, *discordgo.InteractionCreate)),
		commandRegistry: make(map[string]*discordgo.ApplicationCommand),
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

// Registers an app command with discord and adds the respective handler to the svc handler registry.
func (d *DiscordService) RegisterAppCommand(command models.DiscordAppCommand) (*discordgo.ApplicationCommand, error) {
	d.logger.Debug(fmt.Sprintf("adding handler for %s to registry", command.GetAppCommand().Name))
	d.handlerRegistry[command.GetAppCommand().Name] = command.GetAppCommandHandler()

	cmd, err := d.discordSession.ApplicationCommandCreate(d.discordSession.State.User.ID, "", command.GetAppCommand())
	if err != nil {
		d.logger.Error("error registering app command")
		return nil, err
	}
	d.commandRegistry[command.GetAppCommand().Name] = cmd
	return cmd, nil
}

// Adds additional handlers to the svc handler registry.
// Intended for use by MessageComponent handlers
func (d *DiscordService) RegisterMessageComponentHandlers(additionalHandlers map[string]func(*discordgo.Session, *discordgo.InteractionCreate)) error {
	for k, h := range additionalHandlers {
		d.logger.Debug(fmt.Sprintf("adding handler for %s to registry", k))
		d.handlerRegistry[k] = h
	}
	return nil
}

// Proxies InteractionCreate events to the handlers in the svc handler registry
func (d *DiscordService) AddInteractionHandlerProxy() {
	d.discordSession.AddHandler(
		func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			switch i.Type {
			case discordgo.InteractionApplicationCommand:
				if h, ok := d.handlerRegistry[i.ApplicationCommandData().Name]; ok {
					h(s, i)
				} else {
					logger.Info("no handler registered for command", zap.String("command", i.ApplicationCommandData().Name))
				}
			case discordgo.InteractionMessageComponent:
				if h, ok := d.handlerRegistry[i.MessageComponentData().CustomID]; ok {
					h(s, i)
				} else {
					logger.Info("no handler registered for message component", zap.String("customID", i.MessageComponentData().CustomID))
				}
			case discordgo.InteractionModalSubmit:
				if h, ok := d.handlerRegistry[i.ModalSubmitData().CustomID]; ok {
					h(s, i)
				} else {
					logger.Info("no handler registered for modal submit data", zap.String("customID", i.ModalSubmitData().CustomID))
				}
			}
		})
}

// De-registers all app commands registered with this service.
// Intended for use by the shutdown handler.
func (d *DiscordService) DeleteAppCommands() {
	d.logger.Debug("deleting app commands")
	for _, cmd := range d.commandRegistry {
		err := d.discordSession.ApplicationCommandDelete(d.discordSession.State.User.ID, "", cmd.ID)
		if err != nil {
			d.logger.Error(fmt.Sprintf("Cannot delete '%s' command", cmd.Name), zap.Error(err))
		}
	}
}

// func (d *DiscordService) saveDiscordUser(user *discordgo.User) error {
// 	if !d.dbEnabled {
// 		return nil
// 	}
// 	_, err := d.db.DiscordUser.
// 		Create().
// 		SetID(uuid.New()).
// 		SetUsername(user.Username).
// 		SetDiscordid(user.ID).
// 		SetEmail(user.Email).
// 		SetDiscriminator(user.Discriminator).
// 		Save(context.Background())
// 	return err
// }

// func (d *DiscordService) syncGuildUsers(guildId string) error {
// 	if !d.dbEnabled {
// 		return nil
// 	}
// 	guild, err := d.discordSession.Guild(guildId)
// 	if err != nil {
// 		d.logger.Error("error getting guild", zap.Error(err))
// 		return err
// 	}

// 	for _, member := range guild.Members {
// 		_, err := d.db.DiscordUser.
// 			Create().
// 			SetID(uuid.New()).
// 			SetUsername(member.User.Username).
// 			SetDiscordid(member.User.ID).
// 			SetEmail(member.User.Email).
// 			SetDiscriminator(member.User.Discriminator).
// 			Save(context.Background())
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }
