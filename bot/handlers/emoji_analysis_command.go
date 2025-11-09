package handlers

import (
	"context"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/h3mmy/bloopyboi/internal/models"
)

const (
	AnalyzeEmojiCommandName = "analyze-emoji"
)

type AnalyzeEmojiCommand struct {
	discordSvc *services.DiscordService
}

func NewAnalyzeEmojiCommand(discordSvc *services.DiscordService) *AnalyzeEmojiCommand {
	return &AnalyzeEmojiCommand{
		discordSvc: discordSvc,
	}
}

func (c *AnalyzeEmojiCommand) GetAppCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        AnalyzeEmojiCommandName,
		Description: "Analyze an emoji using Google Vision AI",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "emoji",
				Description: "The emoji to analyze",
				Required:    true,
			},
		},
	}
}

func (c *AnalyzeEmojiCommand) GetMessageComponentHandlers() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return nil
}

func (c *AnalyzeEmojiCommand) GetAppCommandHandler() func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		emojiArg := i.ApplicationCommandData().Options[0].StringValue()
		if strings.HasPrefix(emojiArg, "<:") && strings.HasSuffix(emojiArg, ">") {
			trimmed := strings.Trim(emojiArg, "<:>")
			parts := strings.Split(trimmed, ":")
			if len(parts) == 2 {
				emojiID := parts[1]
				emojiURL := fmt.Sprintf("https://cdn.discordapp.com/emojis/%s.png", emojiID)

				analysis, err := c.discordSvc.GetImageAnalyzerService().AnalyzeImageFromURL(context.Background(), emojiURL)
				if err != nil {
					s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
						Type: discordgo.InteractionResponseChannelMessageWithSource,
						Data: &discordgo.InteractionResponseData{
							Content: fmt.Sprintf("Failed to analyze emoji: %v", err),
						},
					})
					return
				}

				// Format and send the results
				response := fmt.Sprintf("Analysis for %s:\nKeywords: %s\nSentiment: %.2f", emojiArg, strings.Join(analysis.Keywords, ", "), analysis.Sentiment)
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: response,
					},
				})
				return
			}
		}
	}
}

func (c *AnalyzeEmojiCommand) WithGuild(guildID string) models.DiscordAppCommand {
	return c
}

func (c *AnalyzeEmojiCommand) GetGuildID() string {
	return ""
}

func (c *AnalyzeEmojiCommand) GetAllowedRoles() []int64 {
	return nil
}
