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
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

// customTimeFormat holds custom time format string.
const (
	customTimeFormat = "2006-01-02T15:04:05Z"

	// discordBotMentionRegexFmt supports also nicknames (the exclamation mark).
	// Read more: https://discordjs.guide/miscellaneous/parsing-mention-arguments.html#how-discord-mentions-work
	discordBotMentionRegexFmt = "^<@!?%s>"
)

type DiscordClient struct {
	botMentionRegex    *regexp.Regexp
	log                *zap.Logger
	botId              string
	api                *discordgo.Session
	registeredCommands []*discordgo.ApplicationCommand
}

// Constructs new Discord Client
func NewDiscordClient(logger *zap.Logger) (*DiscordClient, error) {
	// Get token
	token := providers.GetBotToken()
	botID := providers.GetBotName()

	botMentionRegex, err := regexp.Compile(fmt.Sprintf(discordBotMentionRegexFmt, botID))
	if err != nil {
		return nil, fmt.Errorf("while compiling bot mention regex: %w", err)
	}

	// Create a new Discord session using the provided bot token.
	s, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, fmt.Errorf("Error Creating Discord Session: %w", err)
	}
	return &DiscordClient{
		botId:           botID,
		api:             s,
		log:             logger,
		botMentionRegex: botMentionRegex,
	}, nil
}

// Initiates websocket connection with Discord and starts listening
func (d *DiscordClient) Start(ctx context.Context) error {
	d.log.Info("Starting Bot")
	d.api.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
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
	d.api.AddHandler(bloopyCommands.DirectMessageCreate)
	d.api.AddHandler(bloopyCommands.DirectedMessageReceive)

	d.log.Debug("Registered Handlers...")

	d.api.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsDirectMessages | discordgo.IntentDirectMessageReactions | discordgo.IntentGuildMessageReactions | discordgo.IntentGuildEmojis
	// Open a websocket connection to Discord and begin listening.
	d.log.Info("Opening Websocket Connection")
	err := d.api.Open()
	if err != nil {
		return fmt.Errorf("While opening a connection: %w", err)
	}

	d.log.Info("Registering Commands")
	d.registeredCommands = make([]*discordgo.ApplicationCommand, len(providers.AppCommands))
	for i, v := range providers.AppCommands {
		// Leaving GuildId empty
		cmd, err := d.api.ApplicationCommandCreate(d.api.State.User.ID, "", v)
		if err != nil {
			d.log.Sugar().Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		d.registeredCommands[i] = cmd
	}

	d.log.Info("Initializing Experimental Handler")
	msgSendChan := make(chan *models.DiscordMessageSendRequest, 20)
	expHandler := getBloopyChanHandler(d.api, &msgSendChan)

	ctx, cancelFn := context.WithCancel(ctx)
	defer cancelFn()

	errGroup, ctx := errgroup.WithContext(ctx)
	errGroup.Go(func() error {
		return expHandler.Start(ctx)
	})
	errGroup.Go(func() error {
		return bloopyCommands.StartChannelMessageActor(ctx, d.api, &msgSendChan)
	})

	<-ctx.Done()

	d.log.Info("Received ctx.Done() Exiting...")

	d.log.Info("Removing registered commands...")
	for _, v := range d.registeredCommands {
		err := d.api.ApplicationCommandDelete(d.api.State.User.ID, "", v.ID)
		if err != nil {
			d.log.Sugar().Panicf("Cannot delete '%v' command: %v", v.Name, err)
		}
	}

	d.log.Info("Closing Connection")
	err = d.api.Close()
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

func (d *DiscordClient) IsReady() bool {
	return d.api.DataReady
}
