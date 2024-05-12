package services

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/h3mmy/bloopyboi/bot/internal/config"
	"github.com/h3mmy/bloopyboi/bot/internal/database"
	"github.com/h3mmy/bloopyboi/bot/internal/log"
	"github.com/h3mmy/bloopyboi/bot/internal/models"
	"github.com/h3mmy/bloopyboi/ent"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const DefaultIntents = discordgo.IntentsGuildMessages | discordgo.IntentsDirectMessages | discordgo.IntentDirectMessageReactions | discordgo.IntentGuildMessageReactions | discordgo.IntentGuildEmojis

type DiscordService struct {
	meta           models.BloopyMeta
	logger         *zap.Logger
	discordSession *discordgo.Session
	// The interaction handlers registered with this service, keyed by the command name
	interactionHandlerRegistry map[string]func(*discordgo.Session, *discordgo.InteractionCreate)
	// The commands registered with discord that will need to be de-registered on shutdown
	commandRegistry map[string]*discordgo.ApplicationCommand
	db              *ent.Client
	dbEnabled       bool
	config          *config.DiscordConfig
	intents         discordgo.Intent
}

func NewDiscordService() *DiscordService {
	lgr := log.NewZapLogger().
		Named("discord_service").
		With(zapcore.Field{
			Key:    ServiceLoggerFieldKey,
			Type:   zapcore.StringType,
			String: "discord",
		})
	return &DiscordService{
		meta:            models.NewBloopyMeta(),
		logger:          lgr,
		discordSession:  nil,
		dbEnabled:       false,
		db:              nil,
		intents:         DefaultIntents,
		interactionHandlerRegistry: make(map[string]func(*discordgo.Session, *discordgo.InteractionCreate)),
		commandRegistry: make(map[string]*discordgo.ApplicationCommand),
	}
}

func (d *DiscordService) WithSession(session *discordgo.Session) *DiscordService {
	d.discordSession = session
	d.discordSession.Identify.Intents = d.intents
	return d
}

// NewDiscordServiceWithToken creates a new DiscordService with a token
// Oauth tokens need to be prefixed with "Bearer " instead so this won't work for that
func (d *DiscordService) WithToken(token string) *DiscordService {
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		d.logger.Error("failed to create discord session", zap.Error(err))
		return nil
	}
	return d.WithSession(session)
}

func (d *DiscordService) WithConfig(cfg *config.DiscordConfig) *DiscordService {
	d = d.WithToken(cfg.GetToken())
	d.config = cfg
	return d
}

func (d *DiscordService) RefreshDBConnection() error {
	if d.dbEnabled {
		d.db.Close()
	}
	dbEnabled := true
	dbClient, err := database.Open()
	if err != nil {
		d.logger.Error("failed to open database", zap.Error(err))
		dbEnabled = false
	}
	d.db = dbClient
	d.dbEnabled = dbEnabled

	return err
}

func (d *DiscordService) GetMeta() models.BloopyMeta {
	return d.meta
}

// Primarily for backwards compatibility while I move things into a service
func (d *DiscordService) GetSession() *discordgo.Session {
	return d.discordSession
}

func (d *DiscordService) AddHandler(handler interface{}) func() {
	logger.Debug("Adding simple handler")
	return d.discordSession.AddHandler(handler)
}

func (d *DiscordService) GetDataReady() bool {
	return d.discordSession.DataReady
}

func (d *DiscordService) SetIntents(intents discordgo.Intent) {
	d.intents = intents
	if d.discordSession == nil {
		d.logger.Error("no discord session set")
		return
	}
	d.discordSession.Identify.Intents = intents
}

// Registers an app command with discord and adds the respective handler to the svc handler registry.
func (d *DiscordService) RegisterAppCommand(command models.DiscordAppCommand) (*discordgo.ApplicationCommand, error) {
	d.logger.Debug(fmt.Sprintf("adding handler for %s to registry", command.GetAppCommand().Name))
	d.interactionHandlerRegistry[command.GetAppCommand().Name] = command.GetAppCommandHandler()

	cmd, err := d.discordSession.ApplicationCommandCreate(d.discordSession.State.User.ID, command.GetGuildID(), command.GetAppCommand())
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
		d.interactionHandlerRegistry[k] = h
	}
	return nil
}

// Proxies InteractionCreate events to the handlers in the svc handler registry
func (d *DiscordService) AddInteractionHandlerProxy() {
	d.discordSession.AddHandler(
		func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			switch i.Type {
			case discordgo.InteractionApplicationCommand:
				if h, ok := d.interactionHandlerRegistry[i.ApplicationCommandData().Name]; ok {
					h(s, i)
				} else {
					logger.Info("no handler registered for command", zap.String("command", i.ApplicationCommandData().Name))
				}
			case discordgo.InteractionMessageComponent:
				if h, ok := d.interactionHandlerRegistry[i.MessageComponentData().CustomID]; ok {
					h(s, i)
				} else {
					logger.Info("no handler registered for message component", zap.String("customID", i.MessageComponentData().CustomID))
				}
			case discordgo.InteractionModalSubmit:
				if h, ok := d.interactionHandlerRegistry[i.ModalSubmitData().CustomID]; ok {
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
	allGlobalCmds, err := d.discordSession.ApplicationCommands(d.discordSession.State.User.ID, "")
	if err != nil {
		d.logger.Error("error getting global commands", zap.Error(err))
	} else {
		d.logger.Debug(fmt.Sprintf("found %d global commands", len(allGlobalCmds)))
		for _, cmd := range allGlobalCmds {
			flogger := d.logger.With(zap.String("command", cmd.Name), zap.String("commandID", cmd.ID))
			flogger.Debug(fmt.Sprintf("deleting global command: %v", cmd))
			err := d.discordSession.ApplicationCommandDelete(d.discordSession.State.User.ID, cmd.GuildID, cmd.ID)
			if err != nil {
				flogger.Error("error deleting global command", zap.Error(err))
			} else {
				if d.commandRegistry[cmd.Name] != nil {
					if d.commandRegistry[cmd.Name].ID == cmd.ID {
						delete(d.commandRegistry, cmd.Name)
					} else {
						logger.Warn("commands with same name and different IDs!", zap.String("command", cmd.Name), zap.String("commandID 1", cmd.ID), zap.String("commandID 2", d.commandRegistry[cmd.Name].ID))
					}
				} else {
					d.logger.Warn("deleted command was not found in registry. Likely leftover from a previous instance", zap.String("command", cmd.Name))
				}
			}
		}
	}
	d.logger.Debug("deleting app commands")
	for _, cmd := range d.commandRegistry {
		err := d.discordSession.ApplicationCommandDelete(d.discordSession.State.User.ID, cmd.GuildID, cmd.ID)
		if err != nil {
			d.logger.Error(fmt.Sprintf("Cannot delete '%s' command", cmd.Name), zap.Error(err))
		}
	}
}

// Gets all app commands registered with the discord session AND the discord Registry
// Uses service registry for retrieval IDs and errors are logged
func (d *DiscordService) GetCurrentAppCommands() []*discordgo.ApplicationCommand {
	var commands []*discordgo.ApplicationCommand
	for _, command := range d.commandRegistry {
		d.logger.Debug(fmt.Sprintf("retrieving command: %v", command))
		cmd, err := d.discordSession.ApplicationCommand(d.discordSession.State.User.ID, command.GuildID, command.ID)
		if err != nil {
			d.logger.Error("error retrieving command from discord", zap.String("commandID", command.ID), zap.Error(err))
		} else {
			d.logger.Debug(fmt.Sprintf("retrieved command: %v", cmd))
			commands = append(commands, cmd)
		}
	}
	return commands
}

func (d *DiscordService) SendMessage(messageRequest models.DiscordMessageSendRequest) {
	d.logger.Debug(fmt.Sprintf("sending message: %v", messageRequest))
	_, err := d.discordSession.ChannelMessageSendComplex(messageRequest.ChannelID, messageRequest.MessageComplex)
	if err != nil {
		d.logger.Error("error sending discord message", zap.Error(err))
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
