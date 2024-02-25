package config

import (
	"errors"

	"github.com/h3mmy/bloopyboi/bot/internal/log"
	"go.uber.org/zap/zapcore"
)

var (
	logger = log.NewZapLogger().With(
		zapcore.Field{
			Key:    "group",
			Type:   zapcore.StringType,
			String: "config",
		}).Sugar()
)

// App Config
type AppConfig struct {
	DiscordConfig *DiscordConfig `mapstructure:"discord"`
	Features      []FeatureConfig
	LogLevel      string         `mapstructure:"logLevel"`
	DBConfig      *PostgresConfig `mapstructure:"db"`
	FeatureMap    map[string]FeatureConfig
}

type DiscordConfig struct {
	Token string `mapstructure:"token"`
	AppName  string `mapstructure:"name"`
	AppID    int64  `mapstructure:"appId"`
	GuildConfigs []DiscordGuildConfig `mapstructure:"guilds"`
}

func (c *DiscordConfig) GetToken() string {
	return c.Token
}

// Guild Specific Config
type DiscordGuildConfig struct {
	GuildId     string `mapstructure:"id"`
	// Channel to be used for bot-specific announcements. If empty, no announcement will be sent.
	Announcement *AnnouncementConfig `mapstructure:"announcement"`
	GuildCommandConfig []GuildCommandConfig `mapstructure:"commands"`
}

// Config for dedicated announcement channel. If empty, no announcement will be sent.
// the bot MUST have MANAGE_CHANNEL permissions
// This will be a GUILD_ANNOUNCEMENT type channel https://discord.com/developers/docs/resources/channel
type AnnouncementConfig struct {
	// Channel to be used for bot-specific announcements.
	Channel struct {
		Name string `mapstructure:"name"`
		ID string `mapstructure:"id"`
	}
	NSFW bool `mapstructure:"nsfw"`
}

type GuildCommandConfig struct {
	Name    string `mapstructure:"name"`
	Enabled bool   `mapstructure:"enabled"`
	// Allowed channels for command to be used in. If empty, command can be used in any channel.
	// If not empty, command can only be used in channels listed here.
	// Channels are case sensitive.
	// Example:
	// 		"channels": [
	//			"#general",
	//			"#random"
	//		]
	// This command can only be used in #general and #random channels.
	// If empty, command can be used in any channel.
	// If not empty, command can only be used in channels listed here.
	// Channels are case sensitive.
	// Example:
	// 		"channels": [
	//			"#general",
	//			"#random"
	//		]
	// This command can only be used in #general and #random channels.
	// If empty, command can be used in any channel.
	// If not empty, command can only be used in channels listed here.
	// Channels are case sensitive.
	// Example:
	// 		"channels": [
	//			"#general",
	Channels []int64 `mapstructure:"channels"`
	// Roles allowed to use command. If empty, command can be used by anyone.
	Roles []int64 `mapstructure:"roles"`
}

// Feature Specific Config
type FeatureConfig struct {
	Name    string `mapstructure:"name"`
	Enabled bool   `mapstructure:"enabled"`
	Data    map[string]string
}

// Gets FeatureConfig for Key if exists
func (myConfig *AppConfig) GetFeatureConfigViaMap(name string) (FeatureConfig, error) {
	feat, ok := myConfig.FeatureMap[name]
	if ok {
		return feat, nil
	}
	logger.Error("Could not find config for feature", name)
	return FeatureConfig{}, errors.New("could not find config for feature")
}

// Deprecated until the map version works
func (myConfig *AppConfig) GetFeatureConfig(name string) (FeatureConfig, error) {
	for _, feat := range myConfig.Features {
		if feat.Name == name {
			return feat, nil
		}
	}
	logger.Error("Could not find config for feature", name)
	return FeatureConfig{}, errors.New("could not find config for feature")
}

// Builds FeatureMap. Faster to reference than array
func (myConfig *AppConfig) buildFeatureMap() error {
	logger.Debug("Building Feature Map")
	featMap := make(map[string]FeatureConfig)
	logger.Debug("Feature List", myConfig.Features)
	for _, feat := range myConfig.Features {
		logger.Debug("Adding feature ", feat)
		featMap[feat.Name] = feat
	}
	logger.Debug("Feature Map ", featMap)
	myConfig.FeatureMap = featMap
	return nil
}

// deprecated. Refactor to use FeatureMap Keys
func (myConfig *AppConfig) GetConfiguredFeatureNames() []string {
	var names []string
	for _, feat := range myConfig.Features {
		names = append(names, feat.Name)
	}
	return names
}
