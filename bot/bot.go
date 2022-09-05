package bot

import (
	"context"

	"github.com/alexliesenfeld/health"
	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gitlab.com/h3mmy/bloopyboi/bot/discord"
	"gitlab.com/h3mmy/bloopyboi/bot/internal/config"
	"gitlab.com/h3mmy/bloopyboi/bot/providers"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

const (
	botLogFieldKey = "bot"
)

type BloopyBoi struct {
	log					logrus.FieldLogger
	DB					*gorm.DB
	DiscordClient		*discord.DiscordClient
	Config				*config.BotConfig
	Status				*health.AvailabilityStatus
}

func New() *BloopyBoi {
	return &BloopyBoi{}
}

func (bot *BloopyBoi) WithLogger(logger	logrus.FieldLogger) *BloopyBoi {
	logger.Debug("Adding Logger to boi")
	return &BloopyBoi{
		log:			logger,
	}
}

func (bot *BloopyBoi) Start(ctx context.Context) error {
	if bot.log == nil {
		bot.log = providers.CommonLogger.WithField(botLogFieldKey, "BloopyBoi")
		bot.log.Info("No Logger Detected. Using default field logger")
	}
	errGroup, ctx := errgroup.WithContext(ctx)
	errGroup.Go(func() error {
		bot.log.Debug("Starting Discord Client...")
		return bot.initializeDiscord(ctx)
	})
	errGroup.Go(func() error {
		bot.log.Debug("Initializing Database...")
		return bot.initializeDatabase(ctx)
	})

	<- ctx.Done()

	bot.log.Info("Shutting down Boi. context should propogate")
	return nil
}

func (bot *BloopyBoi) initializeDatabase(ctx context.Context) error {

	botConfig, err := config.GetConfig()
		if err != nil {
			bot.log.Error("Unable to get Config: ", err)
		}
	dbMgr := providers.NewBloopyDBManager(botConfig)
	dbMgr, err = dbMgr.WithSqliteDatabase()
	if err != nil {
		bot.log.Error("Error Initializing DB for boi")
		return err
	}

	bot.DB, err = dbMgr.GetDB()
	if err != nil {
		bot.log.Error("Could not get DB for boi")
		return err
	}
	return nil
}

func (bot *BloopyBoi) initializeDiscord(ctx context.Context) error {

	discordClient, err := discord.NewDiscordClient(bot.log.WithField(botLogFieldKey, "Discord"))
	if err != nil {
		bot.log.Panicf("Error Creating Discord Client %v", err)
		return err
	}

	bot.DiscordClient = discordClient

	bot.log.Debug("Starting Discord Client...")
	return bot.DiscordClient.Start(ctx)

}

// createMessageEvent logs a given message event into the database.
func (bot *BloopyBoi) createMessageEvent(c string, m *discordgo.Message) {
	uuid := uuid.New().String()
	bot.DB.Create(&discord.MessageEvent{
		UUID:           uuid,
		AuthorId:       m.Author.ID,
		AuthorUsername: m.Author.Username,
		MessageId:      m.ID,
		Command:        c,
		ChannelId:      m.ChannelID,
		ServerID:       m.GuildID,
	})
}
