package handlers

import (
	"context"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/h3mmy/bloopyboi/bot/services"
	"github.com/h3mmy/bloopyboi/internal/models"
	log "github.com/h3mmy/bloopyboi/pkg/logs"
	"go.uber.org/zap"
)

const (
	AnalyzeEmojiCommandName = "analyze-emoji"
)

type AnalyzeEmojiCommand struct {
	meta        models.BloopyMeta
	Name        string
	Description string
	guildId string
	logger      *zap.Logger
	roles       []int64
	discordSvc  *services.DiscordService
}

func NewAnalyzeEmojiCommand(discordSvc *services.DiscordService) *AnalyzeEmojiCommand {
	return &AnalyzeEmojiCommand{
		meta:       models.NewBloopyMeta(),
		logger:     log.NewZapLogger().Named("analyze_emoji_command"),
		Name: AnalyzeEmojiCommandName,
		Description: "Analyze an emoji using an ML model",
		discordSvc: discordSvc,
	}
}

// WithGuild sets the guild ID for the command.
func (c *AnalyzeEmojiCommand) WithGuild(guildId string) *AnalyzeEmojiCommand {
	c.guildId = guildId
	return c
}

// WithRoles sets the allowed roles for the command.
func (b *AnalyzeEmojiCommand) WithRoles(roles ...int64) *AnalyzeEmojiCommand {
	b.roles = roles
	return b
}

// GetAllowedRoles returns the allowed roles for the command.
func (b *AnalyzeEmojiCommand) GetAllowedRoles() []int64 {
	return b.roles
}

func (c *AnalyzeEmojiCommand) GetAppCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        AnalyzeEmojiCommandName,
		Description: "Analyze an emoji using an ML model",
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

func (c *AnalyzeEmojiCommand) GetGuildID() string {
	return c.guildId
}
