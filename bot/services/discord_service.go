package services

// import (
// 	"github.com/h3mmy/bloopyboi/bot/discord"
// 	"github.com/h3mmy/bloopyboi/bot/internal/log"
// 	"go.uber.org/zap"
// 	"go.uber.org/zap/zapcore"
// )

// type DiscordService struct {
// 	logger        *zap.Logger
// 	discordClient *discord.DiscordClient
// }

// func NewDiscordService(discordClient *discord.DiscordClient) *DiscordService {
// 	lgr := log.NewZapLogger().With(zapcore.Field{
// 		Key:    ServiceLoggerFieldKey,
// 		Type:   zapcore.StringType,
// 		String: "discord",
// 	})
// 	return &DiscordService{
// 		logger:        lgr,
// 		discordClient: discordClient,
// 	}
// }
