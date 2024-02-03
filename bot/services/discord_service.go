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

