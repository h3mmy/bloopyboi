package handlers

import (
	"context"
	"fmt"
	"regexp"
	"slices"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/h3mmy/bloopyboi/bot/services"
	"github.com/h3mmy/bloopyboi/internal/models"
	log "github.com/h3mmy/bloopyboi/pkg/logs"
	"go.uber.org/zap"
)

type AnalyzeEmojiCommand struct {
	meta               models.BloopyMeta
	Name               string
	Description        string
	guildId            string
	logger             *zap.Logger
	roles              []int64
	imageAnalyzingSvc  *services.ImageAnalyzerService
	discordEmojiRegexp *regexp.Regexp
}

func NewAnalyzeEmojiCommand(imageAnalyzerSvc *services.ImageAnalyzerService) *AnalyzeEmojiCommand {
	logger := log.NewZapLogger().Named("analyze_command")
	rexep, err := regexp.Compile(`<(a|):([A-Za-z0-9_~]+):([0-9]{18,20})>`)
	if err != nil {
		logger.Error("error compiling regexp", zap.Error(err))
	}
	return &AnalyzeEmojiCommand{
		meta:               models.NewBloopyMeta(),
		logger:             logger,
		Name:               string(Analyze),
		Description:        "Analyze an emoji using an ML model",
		imageAnalyzingSvc:  imageAnalyzerSvc,
		discordEmojiRegexp: rexep,
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
		Name:        c.Name,
		Description: c.Description,
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
		// examples of expected values
		// <:litolshy:1236822872811634699>
		// <a:batty_giggle:1071603174760251414>

		matchSlice := c.discordEmojiRegexp.FindStringSubmatch(emojiArg)
		c.logger.Debug("matched a discord emoji", zap.Any("match_slice", matchSlice))

		if len(matchSlice) == 0 || len(matchSlice) == 1 {
			response := fmt.Sprintf("🙁 failed to parse emojiID from string: %s\nGot matchSlice:%+v", emojiArg, matchSlice)

			errR := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: response,
				},
			})
			if errR != nil {
				c.logger.Error("error responding to interaction", zap.Error(errR))
			}
		}
		emojiID := matchSlice[len(matchSlice)-1]
		animated := slices.Contains(matchSlice, "a")
		extension := "png" // Could also be jpg or webp
		if animated {
			extension = "gif"
		}

		emojiURL := fmt.Sprintf("https://cdn.discordapp.com/emojis/%s.%s", emojiID, extension)

		analysis, err := c.imageAnalyzingSvc.AnalyzeImageFromURL(context.Background(), emojiURL)
		if err != nil {
			c.logger.Error("error analyzing image", zap.Error(err))
			errR := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: fmt.Sprintf("Failed to analyze emoji: %v", err),
				},
			})
			if errR != nil {
				c.logger.Error("error responding to interaction", zap.Error(errR))
			}
			return
		} else {
			c.logger.Debug("Got analysis results", zap.Any("analysis", analysis))
		}


		// Format and send the results
		response := fmt.Sprintf("Analysis for %s:\nKeywords: %s\nSafeSearchAnnotation: %s",
			emojiArg,
			strings.Join(analysis.GetKeywordsSortedByScore(), ", "),
			PrintJSON(analysis.SafeSearchAnalysis))

		errR := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: response,
			},
		})
		if errR != nil {
			c.logger.Error("error responding to interaction", zap.Error(errR))
		}

	}
}

func (c *AnalyzeEmojiCommand) GetGuildID() string {
	return c.guildId
}
