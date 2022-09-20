package discord

import (
	"context"
	"fmt"
	"regexp"

	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
	bloopyCommands "gitlab.com/h3mmy/bloopyboi/bot/discord/commands"
	"gitlab.com/h3mmy/bloopyboi/bot/providers"
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
	log                logrus.FieldLogger
	botId              string
	api                *discordgo.Session
	registeredCommands []*discordgo.ApplicationCommand
}

// Constructs new Discord Client
func NewDiscordClient(logger logrus.FieldLogger) (*DiscordClient, error) {
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
		if h, ok := bloopyCommands.CommandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
	d.api.AddHandler(bloopyCommands.MessageCreate)
	d.api.AddHandler(bloopyCommands.DirectMessageCreate)
	d.api.AddHandler(bloopyCommands.DirectedMessageReceive)

	d.log.Debug("Registered Handlers...")

	d.api.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsDirectMessages
	// Open a websocket connection to Discord and begin listening.
	d.log.Info("Opening Websocket Connection")
	err := d.api.Open()
	if err != nil {
		return fmt.Errorf("While opening a connection: %w", err)
	}

	d.log.Info("Registering Commands")
	d.registeredCommands = make([]*discordgo.ApplicationCommand, len(bloopyCommands.Commands))
	for i, v := range bloopyCommands.Commands {
		// Leaving GuildId empty
		cmd, err := d.api.ApplicationCommandCreate(d.api.State.User.ID, "", v)
		if err != nil {
			d.log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		d.registeredCommands[i] = cmd
	}

	<-ctx.Done()

	d.log.Info("Received ctx.Done() Exiting...")

	d.log.Info("Removing registered commands...")
	for _, v := range d.registeredCommands {
		err := d.api.ApplicationCommandDelete(d.api.State.User.ID, "", v.ID)
		if err != nil {
			d.log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
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
