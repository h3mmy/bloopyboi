package discord

import (
	"context"
	"fmt"
	"regexp"

	"github.com/bwmarrin/discordgo"
	bloopyCommands "github.com/h3mmy/bloopyboi/bot/discord/commands"
	"github.com/h3mmy/bloopyboi/bot/handlers"
	"github.com/h3mmy/bloopyboi/bot/internal/models"
	"github.com/h3mmy/bloopyboi/bot/providers"
	"github.com/h3mmy/bloopyboi/bot/services"
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
	botMentionRegex    *regexp.Regexp
	log                *zap.Logger
	botId              string
	discordSvc            *services.DiscordService
	registeredCommands []*discordgo.ApplicationCommand
}

// Constructs new Discord Manager
func NewDiscordManager(logger *zap.Logger) (*DiscordManager, error) {
	// Get token
	token := providers.GetBotToken()
	botID := providers.GetBotName()

	botMentionRegex, err := regexp.Compile(fmt.Sprintf(discordBotMentionRegexFmt, botID))
	if err != nil {
		return nil, fmt.Errorf("while compiling bot mention regex: %w", err)
	}

	// Create a new Discord session using the provided bot token.
	s, err := providers.NewDiscordServiceWithToken(token)
	if err != nil {
		return nil, fmt.Errorf("Error Creating Discord Session: %w", err)
	}
	return &DiscordManager{
		botId:           botID,
		discordSvc:         s,
		log:             logger,
		botMentionRegex: botMentionRegex,
	}, nil
}

// Initiates websocket connection with Discord and starts listening
func (d *DiscordManager) Start(ctx context.Context) error {
	d.log.Info("Starting Bot")
	d.discordSvc.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		switch i.Type {
		case discordgo.InteractionApplicationCommand:
			if h, ok := providers.AppCommandHandlers[i.ApplicationCommandData().Name]; ok {
				h(s, i)
			}
		case discordgo.InteractionMessageComponent:
			if h, ok := providers.MessageComponentHandlers[i.MessageComponentData().CustomID]; ok {
				h(s, i)
			}
		case discordgo.InteractionModalSubmit:
			if h, ok := providers.ModalSubmitHandlers[i.ModalSubmitData().CustomID]; ok {
				h(s, i)
			}
		}
	})
	d.discordSvc.AddHandler(bloopyCommands.DirectMessageCreate)
	d.discordSvc.AddHandler(bloopyCommands.DirectedMessageReceive)

	d.log.Debug("Registered Handlers...")

	d.discordSvc.GetSession().Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsDirectMessages | discordgo.IntentDirectMessageReactions | discordgo.IntentGuildMessageReactions | discordgo.IntentGuildEmojis
	// Open a websocket connection to Discord and begin listening.
	d.log.Info("Opening Websocket Connection")
	err := d.discordSvc.GetSession().Open()
	if err != nil {
		return fmt.Errorf("While opening a connection: %w", err)
	}

	d.log.Info("Registering Commands")
	d.registeredCommands = make([]*discordgo.ApplicationCommand, len(providers.AppCommands))
	for i, v := range providers.AppCommands {
		// Leaving GuildId empty
		cmd, err := d.discordSvc.GetSession().ApplicationCommandCreate(d.discordSvc.GetSession().State.User.ID, "", v)
		if err != nil {
			d.log.Sugar().Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		d.registeredCommands[i] = cmd
	}

	d.log.Info("Initializing Experimental Handler")
	msgSendChan := make(chan *models.DiscordMessageSendRequest, 20)
	expHandler := getBloopyChanHandler(d.discordSvc.GetSession(), &msgSendChan)

	ctx, cancelFn := context.WithCancel(ctx)
	defer cancelFn()

	errGroup, ctx := errgroup.WithContext(ctx)
	errGroup.Go(func() error {
		return expHandler.Start(ctx)
	})
	errGroup.Go(func() error {
		return bloopyCommands.StartChannelMessageActor(ctx, d.discordSvc.GetSession(), &msgSendChan)
	})

	<-ctx.Done()

	d.log.Info("Received ctx.Done() Exiting...")

	d.log.Info("Removing registered commands...")
	for _, v := range d.registeredCommands {
		err := d.discordSvc.GetSession().ApplicationCommandDelete(d.discordSvc.GetSession().State.User.ID, "", v.ID)
		if err != nil {
			d.log.Sugar().Panicf("Cannot delete '%v' command: %v", v.Name, err)
		}
	}

	d.log.Info("Closing Connection")
	err = d.discordSvc.GetSession().Close()
	if err != nil {
		return fmt.Errorf("while closing connection: %w", err)
	}

	d.log.Info("...Done")
	return nil
}

func getBloopyChanHandler(s *discordgo.Session, msgSendChan *chan *models.DiscordMessageSendRequest) *handlers.MessageChanBlooper {
	createCh := bloopyCommands.NextMessageCreateC(s)
	reactACh := bloopyCommands.NextMessageReactionAddC(s)
	reactRCh := bloopyCommands.NextMessageReactionRemoveC(s)

	return handlers.NewMessageChanBlooper(&createCh, &reactACh, &reactRCh, msgSendChan)
}

func (d *DiscordManager) IsReady() bool {
	return d.discordSvc.GetDataReady()
}
