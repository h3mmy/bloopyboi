package discord

import "github.com/bwmarrin/discordgo"

type RoleConnectionDataKey string

const (
	RCKey_Reacts RoleConnectionDataKey = "ct_reacts"
	RCKey_Msgs   RoleConnectionDataKey = "ct_msgs"
)

var RoleConnectionMetadata = []*discordgo.ApplicationRoleConnectionMetadata{
	{
		Type:                     discordgo.ApplicationRoleConnectionMetadataIntegerGreaterThanOrEqual,
		Key:                      string(RCKey_Reacts),
		Name:                     "Number of reactions in server",
		NameLocalizations:        map[discordgo.Locale]string{},
		Description:              "Total reaction count in server",
		DescriptionLocalizations: map[discordgo.Locale]string{},
	},
	{
		Type:                     discordgo.ApplicationRoleConnectionMetadataIntegerGreaterThanOrEqual,
		Key:                      string(RCKey_Msgs),
		Name:                     "Number of messages in server",
		NameLocalizations:        map[discordgo.Locale]string{},
		Description:              "Total message count in server",
		DescriptionLocalizations: map[discordgo.Locale]string{},
	},
}
