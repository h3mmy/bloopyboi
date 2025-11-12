package discord

import (
	"context"
	"fmt"
	"regexp"

	"github.com/bwmarrin/discordgo"
	"github.com/h3mmy/bloopyboi/bot/asynchandlers"
	bloopyCommands "github.com/h3mmy/bloopyboi/bot/discord/commands"
	"github.com/h3mmy/bloopyboi/bot/handlers"
	"github.com/h3mmy/bloopyboi/bot/providers"
	"github.com/h3mmy/bloopyboi/bot/services"
	"github.com/h3mmy/bloopyboi/internal/models"
	"github.com/h3mmy/bloopyboi/pkg/config"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

// customTimeFormat holds custom time format string.
const (
	// customTimeFormat = "2006-01-02T15:04:05Z"

	// discordBotMentionRegexFmt supports also nicknames (the exclamation mark).
	// Read more: https://discordjs.guide/miscellaneous/parsing-mention-arguments.html#how-discord-mentions-work
	discordBotMentionRegexFmt = "^<@!?%s>"
)

// DiscordManager is responsible for interfacing with the discord session
type DiscordManager struct {
	botMentionRegex *regexp.Regexp
	log             *zap.Logger
	botId           int64
	discordSvc      *services.DiscordService
	discordCfg      *config.DiscordConfig
}

// Constructs new Discord Manager
func NewDiscordManager(cfg *config.DiscordConfig, logger *zap.Logger) (*DiscordManager, error) {
	botID := cfg.AppID

	botMentionRegex, err := regexp.Compile(fmt.Sprintf(discordBotMentionRegexFmt, fmt.Sprintf("%d", botID)))
	if err != nil {
		return nil, fmt.Errorf("while compiling bot mention regex: %w", err)
	}

	// Create a new Discord session using the provided bot token.
	s, err := providers.NewDiscordServiceWithConfig(cfg)
	if err != nil {
		return nil, fmt.Errorf("error creating Discord Service: %w", err)
	}
	return &DiscordManager{
		botId:           botID,
		discordSvc:      s,
		log:             logger,
		botMentionRegex: botMentionRegex,
		discordCfg:      cfg,
	}, nil
}

// Initiates websocket connection with Discord and starts listening
func (d *DiscordManager) Start(ctx context.Context) error {
	messageReactor := asynchandlers.NewMessageReactor()
	d.log.Info("Starting Bot")
	d.discordSvc.AddInteractionHandlerProxy()
	d.discordSvc.AddHandler(bloopyCommands.DirectMessageCreate)
	d.discordSvc.AddHandler(bloopyCommands.DirectedMessageReceive)
	d.discordSvc.AddHandler(messageReactor.Handle)

	d.log.Debug("Using config", zap.Any("config", d.discordCfg))

	d.log.Debug("Registered some Handlers... and the proxy")

	d.discordSvc.SetIntents(discordgo.IntentsGuildMessages | discordgo.IntentsDirectMessages | discordgo.IntentDirectMessageReactions | discordgo.IntentGuildMessageReactions | discordgo.IntentGuildEmojis)
	// Open a websocket connection to Discord and begin listening.
	d.log.Info("Opening Websocket Connection")
	err := d.discordSvc.GetSession().Open()
	if err != nil {
		return fmt.Errorf("while opening a connection: %w", err)
	}
	// d.discordSvc.GetSession().LogLevel = discordgo.LogDebug

	d.log.Info("Registering App Commands")
	for _, v := range providers.GetDiscordAppCommands(d.discordCfg.GuildConfigs, d.discordSvc) {
		d.log.Debug("Registering command", zap.Any("command", v.GetAppCommand()))
		_, err := d.discordSvc.RegisterAppCommand(v)
		if err != nil {
			d.log.Sugar().Panicf("Cannot create '%v' command: %v", v.GetAppCommand().Name, err)
		}
		if v.GetMessageComponentHandlers() != nil {
			err = d.discordSvc.RegisterMessageComponentHandlers(v.GetMessageComponentHandlers())
			if err != nil {
				d.log.Error("wasn't expecting this to be possible", zap.Error(err))
			}
		}
	}

	d.log.Info("Initializing Experimental Handler")
	msgSendChan := make(chan *models.DiscordMessageSendRequest, 20)
	expHandler := getBloopyChanHandler(d.discordSvc, &msgSendChan)

	d.log.Debug("rotating through guild configs")
	for _, gcfg := range d.discordCfg.GuildConfigs {
		d.log.Debug("processing guild config", zap.Any("guildConfig", gcfg))
		if gcfg.RoleSelectionConfig != nil {
			d.log.Debug("importing role selection config")
			roleSelector := handlers.NewRoleSelectionHandler(gcfg.GuildId, gcfg.RoleSelectionConfig)
			d.discordSvc.AddHandler(roleSelector.HandleReactionAdd)
			d.discordSvc.AddHandler(roleSelector.HandleReactionRemove)
		}
	}

	ctx, cancelFn := context.WithCancel(ctx)
	defer cancelFn()

	errGroup, ctx := errgroup.WithContext(ctx)
	errGroup.Go(func() error {
		return expHandler.Start(ctx)
	})
	errGroup.Go(func() error {
		return bloopyCommands.StartChannelMessageActor(ctx, d.discordSvc.GetSession(), &msgSendChan)
	})

	// for _, gcfg := range d.discordCfg.GuildConfigs {
	// 	go d.discordSvc.IngestGuildEmojis(ctx, gcfg.GuildId)
	// }

	<-ctx.Done()

	d.log.Info("Received ctx.Done() Exiting...")

	d.log.Info("Removing registered commands...")
	d.discordSvc.DeleteAppCommands()

	d.log.Info("Closing Connection")
	err = d.discordSvc.GetSession().Close()
	if err != nil {
		return fmt.Errorf("while closing connection: %w", err)
	}

	d.log.Info("...Done")
	return nil
}

func getBloopyChanHandler(ds *services.DiscordService, msgSendChan *chan *models.DiscordMessageSendRequest) *handlers.MessageChanBlooper {
	s := ds.GetSession()
	createCh := bloopyCommands.NextMessageCreateC(s)
	reactACh := bloopyCommands.NextMessageReactionAddC(s)
	reactRCh := bloopyCommands.NextMessageReactionRemoveC(s)

	return handlers.NewMessageChanBlooper(ds, providers.GetInspiroService(), &createCh, &reactACh, &reactRCh, msgSendChan)
}

func (d *DiscordManager) IsReady() bool {
	return d.discordSvc.GetDataReady()
}

func (d *DiscordManager) GetDiscordService() *services.DiscordService {
	return d.discordSvc
}
