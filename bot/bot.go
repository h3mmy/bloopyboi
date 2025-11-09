package bot

import (
	"context"
	"fmt"

	"github.com/alexliesenfeld/health"
	"github.com/h3mmy/bloopyboi/bot/discord"
	"github.com/h3mmy/bloopyboi/bot/providers"
	"github.com/h3mmy/bloopyboi/internal/models"
	"github.com/h3mmy/bloopyboi/pkg/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/sync/errgroup"
)

const (
	botLogFieldKey = "bot"
)

type BloopyBoi struct {
	log             *zap.Logger
	DiscordManager  *discord.DiscordManager
	Config          *config.AppConfig
	Status          *health.AvailabilityStatus
	ServiceRegistry models.ServiceRegistry
	Running         bool
	ImageAnalyzer   models.ImageAnalyzer
}

func New() *BloopyBoi {
	return &BloopyBoi{}
}

func (bot *BloopyBoi) WithLogger(logger *zap.Logger) *BloopyBoi {
	logger.Debug("Adding Logger to boi")
	bot.log = logger
	return bot
}

func (bot *BloopyBoi) Run(ctx context.Context) error {
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
	bot.log.Debug(fmt.Sprintf("Experimental is enabled: %v", providers.IsFeatureEnabled("experimental")))

	errGroup, ctx := errgroup.WithContext(ctx)
	errGroup.Go(func() error {
		bot.log.Debug("Starting Discord Client...")
		return bot.initializeDiscord(ctx)
	})
	bot.Running = true
	// errGroup.Go(func() error {
	// 	bot.log.Debug("Starting K8s Service")
	// 	return bot.initializeK8sService(ctx)
	// })
	go func() {
		err := errGroup.Wait()
		if err != nil {
			bot.log.Error("Error in bot errGroup", zap.Error(err))
		}
		bot.log.Info("bot.Run monitor gofunc exiting")
		bot.Running = false
	}()
	<-ctx.Done()

	bot.log.Info("Shutting down Boi. context should propogate")
	return nil
}

func (bot *BloopyBoi) initializeDiscord(ctx context.Context) error {

	discordConfig := providers.GetDiscordConfig()

	discordManager, err := discord.NewDiscordManager(discordConfig, bot.log.With(zapcore.Field{
		Key:    botLogFieldKey,
		Type:   zapcore.StringType,
		String: "Discord",
	}))
	if err != nil {
		bot.log.Sugar().Panicf("Error Creating Discord Client %v", err)
		return err
	}

	bot.DiscordManager = discordManager

	bot.log.Debug("Starting Discord Manager...")
	return bot.DiscordManager.Start(ctx)

}

// func (bot *BloopyBoi) initializeK8sService(ctx context.Context) error {
// 	k8sService := services.NewK8sService()
// 	for _, ns := range k8sService.ListNamespaces(ctx) {
// 		bot.log.Sugar().Info(ns)
// 	}
// 	return nil
// }

// func (bot *BloopyBoi) initializeAuthentikService(ctx context.Context) error {
// 	return nil
// }

func (bot *BloopyBoi) Ping(ctx context.Context) error {
	return nil
}

func (bot *BloopyBoi) GetStatus(ctx context.Context) *health.AvailabilityStatus {
	return bot.Status
}

func (bot *BloopyBoi) GetReadinessChecker() health.Checker {
	discordReady := func() bool {
		return bot.Running
	}

	return providers.NewReadinessChecker(discordReady)
}
