package bot

import (
	"github.com/alexliesenfeld/health"
	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gitlab.com/h3mmy/bloopyboi/bot/discord"
	"gitlab.com/h3mmy/bloopyboi/bot/internal/config"
	"gitlab.com/h3mmy/bloopyboi/bot/providers"
	"gorm.io/gorm"
)

type BloopyBoi struct {
	log					logrus.FieldLogger
	DB					*gorm.DB
	DiscordClient		*discord.DiscordClient
	Config				*config.BotConfig
	ReadinessChecker	*health.Checker
}

func New(log				logrus.FieldLogger,
		db					*gorm.DB,
		discordClient		*discord.DiscordClient,
		config				*config.BotConfig,
		readinessChecker	*health.Checker) *BloopyBoi {
			return &BloopyBoi{
				log:				log,
				DB:					db,
				DiscordClient:		discordClient,
				Config:				config,
				ReadinessChecker:	readinessChecker,
			}
}

func NewBoi(log				logrus.FieldLogger,
	discordClient		*discord.DiscordClient,
	readinessChecker	*health.Checker) *BloopyBoi {
		botConfig, err := config.GetConfig()
		if err != nil {
			log.Error("Unable to get Config: ", err)
		}
		dbMgr := providers.NewBloopyDBManager(botConfig)
		dbMgr.InitSqliteDatabase()

		return &BloopyBoi{
			log:				log,
			DB:					dbMgr.GetDB(),
			DiscordClient:		discordClient,
			Config:				botConfig,
			ReadinessChecker:	readinessChecker,
		}
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
