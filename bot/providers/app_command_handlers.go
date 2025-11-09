package providers

import (
	"github.com/h3mmy/bloopyboi/bot/handlers"
	"github.com/h3mmy/bloopyboi/internal/models"
	"github.com/h3mmy/bloopyboi/pkg/config"
	"go.uber.org/zap"
)

func GetDiscordAppCommands(cfgs []config.DiscordGuildConfig) []models.DiscordAppCommand {
	handls := make([]models.DiscordAppCommand, 0, 3)
	handls = append(handls, handlers.NewInspiroCommand(GetInspiroService()))
	handls = append(handls, GetGuildAppCommands(cfgs)...)
	logger.Debug("got discord commands", zap.Int("count", len(handls)))
	return handls
}

func GetGuildAppCommands(cfgs []config.DiscordGuildConfig) []models.DiscordAppCommand {
	handls := make([]models.DiscordAppCommand, 0, 3)
	logger.Debug("getting configs for guilds", zap.Int("count", len(cfgs)))
	for _, guild := range cfgs {
		logger.Debug("getting guild commands", zap.Int("count", len(guild.GuildCommandConfig)))
		for _, v := range guild.GuildCommandConfig {
			logger.Debug("getting guild command", zap.String("name", v.Name), zap.Bool("enabled", v.Enabled))
			if v.Enabled {
				cmd := GetCommandWithConfig(guild.GuildId, v)
				if cmd != nil {
					handls = append(handls, cmd)
				}
			}
		}
	}
	logger.Debug("got guild commands", zap.Int("count", len(handls)))
	return handls
}

func GetCommandWithConfig(guildId string, cfg config.GuildCommandConfig) models.DiscordAppCommand {
	flogger := logger.With(zap.String("guild_app_command", cfg.Name), zap.String("guild_id", guildId))
	// get from repository TODO
	switch cfg.Name {
	case "blissfest":
		flogger.Debug("Checking if feature enabled")
		if IsFeatureEnabled(models.BlissfestFeatureKey) {
			return handlers.NewBlissfestCommand(GetBlissfestService()).WithGuild(guildId).WithRoles(cfg.Roles...)
		}
		flogger.Warn("blissfest guild command exists but feature is disabled")
	case "book":
		bookSvc, err := GetBookService()
		if err != nil {
			flogger.Error("failed to create book svc", zap.Error(err))
		} else {
			return handlers.NewBookCommand(bookSvc).WithRoles(cfg.Roles...).WithGuild(guildId)
		}
	case "requests":
		bookSvc, err := GetBookService()
		if err != nil {
			flogger.Error("failed to create book svc", zap.Error(err))
		} else {
			return handlers.NewUserRequestCommand(bookSvc).WithRoles(cfg.Roles...).WithGuild(guildId)
		}
	}
	flogger.Warn("not adding command", zap.String("name", cfg.Name))
	return nil
}
