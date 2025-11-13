package config

type DiscordConfig struct {
	Token        string               `mapstructure:"token"`
	AppName      string               `mapstructure:"name"`
	AppID        int64                `mapstructure:"appId"`
	GuildConfigs []DiscordGuildConfig `mapstructure:"guilds"`
}

// Guild Specific Config
type DiscordGuildConfig struct {
	GuildId string `mapstructure:"id"`
	// Channel to be used for bot-specific announcements. If empty, no announcement will be sent.
	Announcement        *AnnouncementConfig  `mapstructure:"announcement"`
	GuildCommandConfig  []GuildCommandConfig `mapstructure:"commands"`
	RoleSelectionConfig *RoleSelectionConfig `mapstructure:"roleSelection"`
}

// RoleSelectionConfig is intended to configure role selection prompts
// Eventually, this should be stateful and customizable, but UI needs to happen first
type RoleSelectionConfig struct {
	Channel struct {
		Name string `mapstructure:"name"`
		ID   string `mapstructure:"id"`
	}
	Prompts []RoleSelectionPrompt `mapstructure:"prompts"`
}

type RoleSelectionPrompt struct {
	Message string `mapstructure:"message"`
	Options []struct {
		EmojiID     string `mapstructure:"emojiID"`
		Description string `mapstructure:"description"`
		RoleID      string `mapstructure:"roleId"`
	}
}

// Config for dedicated announcement channel. If empty, no announcement will be sent.
// the bot MUST have MANAGE_CHANNEL permissions
// This will be a GUILD_ANNOUNCEMENT type channel https://discord.com/developers/docs/resources/channel
type AnnouncementConfig struct {
	// Channel to be used for bot-specific announcements.
	Channel struct {
		Name string `mapstructure:"name"`
		ID   string `mapstructure:"id"`
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
