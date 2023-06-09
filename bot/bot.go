package bot

import (
	"context"
	"fmt"

	"github.com/alexliesenfeld/health"
	"gitlab.com/h3mmy/bloopyboi/bot/discord"
	"gitlab.com/h3mmy/bloopyboi/bot/internal/config"
	"gitlab.com/h3mmy/bloopyboi/bot/internal/models"
	"gitlab.com/h3mmy/bloopyboi/bot/providers"
	"gitlab.com/h3mmy/bloopyboi/bot/services"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/sync/errgroup"
)

const (
	botLogFieldKey = "bot"
)

type BloopyBoi struct {
	log             *zap.Logger
	DiscordClient   *discord.DiscordClient
	Config          *config.BotConfig
	Status          *health.AvailabilityStatus
	ServiceRegistry models.ServiceRegistry
}

func New() *BloopyBoi {
	return &BloopyBoi{}
}

func (bot *BloopyBoi) WithLogger(logger *zap.Logger) *BloopyBoi {
	logger.Debug("Adding Logger to boi")
	return &BloopyBoi{
		log: logger,
	}
}

func (bot *BloopyBoi) Start(ctx context.Context) error {
	if bot.log == nil {
		bot.log = providers.NewZapLogger().With(
			zapcore.Field{
				Key:    botLogFieldKey,
				Type:   zapcore.StringType,
				String: "BloopyBoi",
			})
		bot.log.Info("No Logger Detected. Using default field logger")
	}

	bot.log.Debug(fmt.Sprintf("FeatureMap contains %d entries", len(providers.GetFeatures())))
	bot.log.Debug(fmt.Sprintf("Experimental is enabled: %v", providers.IsFeaturedConfigured("experimental")))

	errGroup, ctx := errgroup.WithContext(ctx)
	errGroup.Go(func() error {
		bot.log.Debug("Starting Discord Client...")
		return bot.initializeDiscord(ctx)
	})
	// errGroup.Go(func() error {
	// 	bot.log.Debug("Starting K8s Service")
	// 	return bot.initializeK8sService(ctx)
	// })

	<-ctx.Done()

	bot.log.Info("Shutting down Boi. context should propogate")
	return nil
}

func (bot *BloopyBoi) initializeDiscord(ctx context.Context) error {

	discordClient, err := discord.NewDiscordClient(bot.log.With(zapcore.Field{
		Key:    botLogFieldKey,
		Type:   zapcore.StringType,
		String: "Discord",
	}))
	if err != nil {
		bot.log.Sugar().Panicf("Error Creating Discord Client %v", err)
		return err
	}

	bot.DiscordClient = discordClient

	bot.log.Debug("Starting Discord Client...")
	return bot.DiscordClient.Start(ctx)

}

func (bot *BloopyBoi) initializeK8sService(ctx context.Context) error {
	k8sService := services.NewK8sService()
	for _, ns := range k8sService.ListNamespaces(ctx) {
		bot.log.Sugar().Info(ns)
	}
	return nil
}

func (bot *BloopyBoi) initializeAuthentikService(ctx context.Context) error {
	return nil
}
