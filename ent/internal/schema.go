// Code generated by ent, DO NOT EDIT.

//go:build tools
// +build tools

// Package internal holds a loadable version of the latest schema.
package internal

const Schema = "{\"Schema\":\"github.com/h3mmy/bloopyboi/ent/schema\",\"Package\":\"github.com/h3mmy/bloopyboi/ent\",\"Schemas\":[{\"name\":\"Book\",\"config\":{\"Table\":\"\"},\"edges\":[{\"name\":\"book_author\",\"type\":\"BookAuthor\",\"ref_name\":\"books\",\"inverse\":true},{\"name\":\"media_request\",\"type\":\"MediaRequest\",\"unique\":true}],\"fields\":[{\"name\":\"id\",\"type\":{\"Type\":4,\"Ident\":\"uuid.UUID\",\"PkgPath\":\"github.com/google/uuid\",\"PkgName\":\"uuid\",\"Nillable\":false,\"RType\":{\"Name\":\"UUID\",\"Ident\":\"uuid.UUID\",\"Kind\":17,\"PkgPath\":\"github.com/google/uuid\",\"Methods\":{\"ClockSequence\":{\"In\":[],\"Out\":[{\"Name\":\"int\",\"Ident\":\"int\",\"Kind\":2,\"PkgPath\":\"\",\"Methods\":null}]},\"Domain\":{\"In\":[],\"Out\":[{\"Name\":\"Domain\",\"Ident\":\"uuid.Domain\",\"Kind\":8,\"PkgPath\":\"github.com/google/uuid\",\"Methods\":null}]},\"ID\":{\"In\":[],\"Out\":[{\"Name\":\"uint32\",\"Ident\":\"uint32\",\"Kind\":10,\"PkgPath\":\"\",\"Methods\":null}]},\"MarshalBinary\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"MarshalText\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"NodeID\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}]},\"Scan\":{\"In\":[{\"Name\":\"\",\"Ident\":\"interface {}\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"String\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"Time\":{\"In\":[],\"Out\":[{\"Name\":\"Time\",\"Ident\":\"uuid.Time\",\"Kind\":6,\"PkgPath\":\"github.com/google/uuid\",\"Methods\":null}]},\"URN\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"UnmarshalBinary\":{\"In\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"UnmarshalText\":{\"In\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"Value\":{\"In\":[],\"Out\":[{\"Name\":\"Value\",\"Ident\":\"driver.Value\",\"Kind\":20,\"PkgPath\":\"database/sql/driver\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"Variant\":{\"In\":[],\"Out\":[{\"Name\":\"Variant\",\"Ident\":\"uuid.Variant\",\"Kind\":8,\"PkgPath\":\"github.com/google/uuid\",\"Methods\":null}]},\"Version\":{\"In\":[],\"Out\":[{\"Name\":\"Version\",\"Ident\":\"uuid.Version\",\"Kind\":8,\"PkgPath\":\"github.com/google/uuid\",\"Methods\":null}]}}}},\"unique\":true,\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"title\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"description\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"position\":{\"Index\":2,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"goodreads_id\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"unique\":true,\"optional\":true,\"position\":{\"Index\":3,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"google_volume_id\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"unique\":true,\"position\":{\"Index\":4,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"isbn_10\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"position\":{\"Index\":5,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"isbn_13\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"position\":{\"Index\":6,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"publisher\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"position\":{\"Index\":7,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"image_url\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"position\":{\"Index\":8,\"MixedIn\":false,\"MixinIndex\":0}}]},{\"name\":\"BookAuthor\",\"config\":{\"Table\":\"\"},\"edges\":[{\"name\":\"books\",\"type\":\"Book\"}],\"fields\":[{\"name\":\"id\",\"type\":{\"Type\":4,\"Ident\":\"uuid.UUID\",\"PkgPath\":\"github.com/google/uuid\",\"PkgName\":\"uuid\",\"Nillable\":false,\"RType\":{\"Name\":\"UUID\",\"Ident\":\"uuid.UUID\",\"Kind\":17,\"PkgPath\":\"github.com/google/uuid\",\"Methods\":{\"ClockSequence\":{\"In\":[],\"Out\":[{\"Name\":\"int\",\"Ident\":\"int\",\"Kind\":2,\"PkgPath\":\"\",\"Methods\":null}]},\"Domain\":{\"In\":[],\"Out\":[{\"Name\":\"Domain\",\"Ident\":\"uuid.Domain\",\"Kind\":8,\"PkgPath\":\"github.com/google/uuid\",\"Methods\":null}]},\"ID\":{\"In\":[],\"Out\":[{\"Name\":\"uint32\",\"Ident\":\"uint32\",\"Kind\":10,\"PkgPath\":\"\",\"Methods\":null}]},\"MarshalBinary\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"MarshalText\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"NodeID\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}]},\"Scan\":{\"In\":[{\"Name\":\"\",\"Ident\":\"interface {}\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"String\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"Time\":{\"In\":[],\"Out\":[{\"Name\":\"Time\",\"Ident\":\"uuid.Time\",\"Kind\":6,\"PkgPath\":\"github.com/google/uuid\",\"Methods\":null}]},\"URN\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"UnmarshalBinary\":{\"In\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"UnmarshalText\":{\"In\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"Value\":{\"In\":[],\"Out\":[{\"Name\":\"Value\",\"Ident\":\"driver.Value\",\"Kind\":20,\"PkgPath\":\"database/sql/driver\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"Variant\":{\"In\":[],\"Out\":[{\"Name\":\"Variant\",\"Ident\":\"uuid.Variant\",\"Kind\":8,\"PkgPath\":\"github.com/google/uuid\",\"Methods\":null}]},\"Version\":{\"In\":[],\"Out\":[{\"Name\":\"Version\",\"Ident\":\"uuid.Version\",\"Kind\":8,\"PkgPath\":\"github.com/google/uuid\",\"Methods\":null}]}}}},\"unique\":true,\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"full_name\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"unique\":true,\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0}}]},{\"name\":\"DiscordGuild\",\"config\":{\"Table\":\"\"},\"edges\":[{\"name\":\"members\",\"type\":\"DiscordUser\"}],\"fields\":[{\"name\":\"id\",\"type\":{\"Type\":4,\"Ident\":\"uuid.UUID\",\"PkgPath\":\"github.com/google/uuid\",\"PkgName\":\"uuid\",\"Nillable\":false,\"RType\":{\"Name\":\"UUID\",\"Ident\":\"uuid.UUID\",\"Kind\":17,\"PkgPath\":\"github.com/google/uuid\",\"Methods\":{\"ClockSequence\":{\"In\":[],\"Out\":[{\"Name\":\"int\",\"Ident\":\"int\",\"Kind\":2,\"PkgPath\":\"\",\"Methods\":null}]},\"Domain\":{\"In\":[],\"Out\":[{\"Name\":\"Domain\",\"Ident\":\"uuid.Domain\",\"Kind\":8,\"PkgPath\":\"github.com/google/uuid\",\"Methods\":null}]},\"ID\":{\"In\":[],\"Out\":[{\"Name\":\"uint32\",\"Ident\":\"uint32\",\"Kind\":10,\"PkgPath\":\"\",\"Methods\":null}]},\"MarshalBinary\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"MarshalText\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"NodeID\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}]},\"Scan\":{\"In\":[{\"Name\":\"\",\"Ident\":\"interface {}\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"String\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"Time\":{\"In\":[],\"Out\":[{\"Name\":\"Time\",\"Ident\":\"uuid.Time\",\"Kind\":6,\"PkgPath\":\"github.com/google/uuid\",\"Methods\":null}]},\"URN\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"UnmarshalBinary\":{\"In\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"UnmarshalText\":{\"In\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"Value\":{\"In\":[],\"Out\":[{\"Name\":\"Value\",\"Ident\":\"driver.Value\",\"Kind\":20,\"PkgPath\":\"database/sql/driver\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"Variant\":{\"In\":[],\"Out\":[{\"Name\":\"Variant\",\"Ident\":\"uuid.Variant\",\"Kind\":8,\"PkgPath\":\"github.com/google/uuid\",\"Methods\":null}]},\"Version\":{\"In\":[],\"Out\":[{\"Name\":\"Version\",\"Ident\":\"uuid.Version\",\"Kind\":8,\"PkgPath\":\"github.com/google/uuid\",\"Methods\":null}]}}}},\"unique\":true,\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"discordid\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"unique\":true,\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"name\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":2,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"description\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"position\":{\"Index\":3,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"rules_channel_id\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"position\":{\"Index\":4,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"public_updates_channel_id\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"position\":{\"Index\":5,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"nsfw_level\",\"type\":{\"Type\":12,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"position\":{\"Index\":6,\"MixedIn\":false,\"MixinIndex\":0}}]},{\"name\":\"DiscordMessage\",\"config\":{\"Table\":\"\"},\"edges\":[{\"name\":\"author\",\"type\":\"DiscordUser\",\"ref_name\":\"discord_messages\",\"inverse\":true}],\"fields\":[{\"name\":\"create_time\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"immutable\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":0}},{\"name\":\"update_time\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"update_default\":true,\"position\":{\"Index\":1,\"MixedIn\":true,\"MixinIndex\":0}},{\"name\":\"id\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"raw\",\"type\":{\"Type\":3,\"Ident\":\"discordgo.Message\",\"PkgPath\":\"github.com/bwmarrin/discordgo\",\"PkgName\":\"discordgo\",\"Nillable\":false,\"RType\":{\"Name\":\"Message\",\"Ident\":\"discordgo.Message\",\"Kind\":25,\"PkgPath\":\"github.com/bwmarrin/discordgo\",\"Methods\":{\"ContentWithMentionsReplaced\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"ContentWithMoreMentionsReplaced\":{\"In\":[{\"Name\":\"\",\"Ident\":\"*discordgo.Session\",\"Kind\":22,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"GetCustomEmojis\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]*discordgo.Emoji\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}]},\"Reference\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"*discordgo.MessageReference\",\"Kind\":22,\"PkgPath\":\"\",\"Methods\":null}]},\"UnmarshalJSON\":{\"In\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]}}}},\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0}}]},{\"name\":\"DiscordUser\",\"config\":{\"Table\":\"\"},\"edges\":[{\"name\":\"guilds\",\"type\":\"DiscordGuild\",\"ref_name\":\"members\",\"inverse\":true},{\"name\":\"discord_messages\",\"type\":\"DiscordMessage\"},{\"name\":\"media_requests\",\"type\":\"MediaRequest\"}],\"fields\":[{\"name\":\"id\",\"type\":{\"Type\":4,\"Ident\":\"uuid.UUID\",\"PkgPath\":\"github.com/google/uuid\",\"PkgName\":\"uuid\",\"Nillable\":false,\"RType\":{\"Name\":\"UUID\",\"Ident\":\"uuid.UUID\",\"Kind\":17,\"PkgPath\":\"github.com/google/uuid\",\"Methods\":{\"ClockSequence\":{\"In\":[],\"Out\":[{\"Name\":\"int\",\"Ident\":\"int\",\"Kind\":2,\"PkgPath\":\"\",\"Methods\":null}]},\"Domain\":{\"In\":[],\"Out\":[{\"Name\":\"Domain\",\"Ident\":\"uuid.Domain\",\"Kind\":8,\"PkgPath\":\"github.com/google/uuid\",\"Methods\":null}]},\"ID\":{\"In\":[],\"Out\":[{\"Name\":\"uint32\",\"Ident\":\"uint32\",\"Kind\":10,\"PkgPath\":\"\",\"Methods\":null}]},\"MarshalBinary\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"MarshalText\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"NodeID\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}]},\"Scan\":{\"In\":[{\"Name\":\"\",\"Ident\":\"interface {}\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"String\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"Time\":{\"In\":[],\"Out\":[{\"Name\":\"Time\",\"Ident\":\"uuid.Time\",\"Kind\":6,\"PkgPath\":\"github.com/google/uuid\",\"Methods\":null}]},\"URN\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"UnmarshalBinary\":{\"In\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"UnmarshalText\":{\"In\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"Value\":{\"In\":[],\"Out\":[{\"Name\":\"Value\",\"Ident\":\"driver.Value\",\"Kind\":20,\"PkgPath\":\"database/sql/driver\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"Variant\":{\"In\":[],\"Out\":[{\"Name\":\"Variant\",\"Ident\":\"uuid.Variant\",\"Kind\":8,\"PkgPath\":\"github.com/google/uuid\",\"Methods\":null}]},\"Version\":{\"In\":[],\"Out\":[{\"Name\":\"Version\",\"Ident\":\"uuid.Version\",\"Kind\":8,\"PkgPath\":\"github.com/google/uuid\",\"Methods\":null}]}}}},\"unique\":true,\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"discordid\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"unique\":true,\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"username\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"unique\":true,\"position\":{\"Index\":2,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"email\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"position\":{\"Index\":3,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"discriminator\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"position\":{\"Index\":4,\"MixedIn\":false,\"MixinIndex\":0}}]},{\"name\":\"MediaRequest\",\"config\":{\"Table\":\"\"},\"edges\":[{\"name\":\"discord_users\",\"type\":\"DiscordUser\",\"ref_name\":\"media_requests\",\"inverse\":true},{\"name\":\"book\",\"type\":\"Book\",\"ref_name\":\"media_request\",\"unique\":true,\"inverse\":true}],\"fields\":[{\"name\":\"create_time\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"immutable\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":0}},{\"name\":\"update_time\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"update_default\":true,\"position\":{\"Index\":1,\"MixedIn\":true,\"MixinIndex\":0}},{\"name\":\"id\",\"type\":{\"Type\":4,\"Ident\":\"uuid.UUID\",\"PkgPath\":\"github.com/google/uuid\",\"PkgName\":\"uuid\",\"Nillable\":false,\"RType\":{\"Name\":\"UUID\",\"Ident\":\"uuid.UUID\",\"Kind\":17,\"PkgPath\":\"github.com/google/uuid\",\"Methods\":{\"ClockSequence\":{\"In\":[],\"Out\":[{\"Name\":\"int\",\"Ident\":\"int\",\"Kind\":2,\"PkgPath\":\"\",\"Methods\":null}]},\"Domain\":{\"In\":[],\"Out\":[{\"Name\":\"Domain\",\"Ident\":\"uuid.Domain\",\"Kind\":8,\"PkgPath\":\"github.com/google/uuid\",\"Methods\":null}]},\"ID\":{\"In\":[],\"Out\":[{\"Name\":\"uint32\",\"Ident\":\"uint32\",\"Kind\":10,\"PkgPath\":\"\",\"Methods\":null}]},\"MarshalBinary\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"MarshalText\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"NodeID\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}]},\"Scan\":{\"In\":[{\"Name\":\"\",\"Ident\":\"interface {}\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"String\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"Time\":{\"In\":[],\"Out\":[{\"Name\":\"Time\",\"Ident\":\"uuid.Time\",\"Kind\":6,\"PkgPath\":\"github.com/google/uuid\",\"Methods\":null}]},\"URN\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"UnmarshalBinary\":{\"In\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"UnmarshalText\":{\"In\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"Value\":{\"In\":[],\"Out\":[{\"Name\":\"Value\",\"Ident\":\"driver.Value\",\"Kind\":20,\"PkgPath\":\"database/sql/driver\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"Variant\":{\"In\":[],\"Out\":[{\"Name\":\"Variant\",\"Ident\":\"uuid.Variant\",\"Kind\":8,\"PkgPath\":\"github.com/google/uuid\",\"Methods\":null}]},\"Version\":{\"In\":[],\"Out\":[{\"Name\":\"Version\",\"Ident\":\"uuid.Version\",\"Kind\":8,\"PkgPath\":\"github.com/google/uuid\",\"Methods\":null}]}}}},\"unique\":true,\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"status\",\"type\":{\"Type\":6,\"Ident\":\"models.MediaRequestStatus\",\"PkgPath\":\"github.com/h3mmy/bloopyboi/internal/models\",\"PkgName\":\"models\",\"Nillable\":false,\"RType\":{\"Name\":\"MediaRequestStatus\",\"Ident\":\"models.MediaRequestStatus\",\"Kind\":24,\"PkgPath\":\"github.com/h3mmy/bloopyboi/internal/models\",\"Methods\":{\"Values\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]string\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}]}}}},\"enums\":[{\"N\":\"requested\",\"V\":\"requested\"},{\"N\":\"approved\",\"V\":\"approved\"},{\"N\":\"success\",\"V\":\"success\"},{\"N\":\"rejected\",\"V\":\"rejected\"},{\"N\":\"cancelled\",\"V\":\"cancelled\"},{\"N\":\"error\",\"V\":\"error\"}],\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"priority\",\"type\":{\"Type\":12,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":50,\"default_kind\":2,\"position\":{\"Index\":2,\"MixedIn\":false,\"MixinIndex\":0}}]}],\"Features\":[\"sql/lock\",\"namedges\",\"sql/upsert\",\"schema/snapshot\"]}"
