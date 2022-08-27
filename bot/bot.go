package bot

import (
	"github.com/alexliesenfeld/health"
	"github.com/sirupsen/logrus"
	"gitlab.com/h3mmy/bloopyboi/bot/discord"
	"gitlab.com/h3mmy/bloopyboi/bot/internal/config"
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
