package handlers

import (
	"context"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func (h *DiscordHandlers) HandleAnalyzeEmojiCmd(
	s *discordgo.Session,
	m *discordgo.MessageCreate,
) {
	parts := strings.Split(m.Content, " ")
	if len(parts) != 2 {
		s.ChannelMessageSend(m.ChannelID, "Usage: !analyze-emoji <emoji>")
		return
	}

	emojiArg := parts[1]
	// Basic parsing for custom emoji, e.g., <:name:id>
	if strings.HasPrefix(emojiArg, "<:") && strings.HasSuffix(emojiArg, ">") {
		trimmed := strings.Trim(emojiArg, "<:>")
		parts := strings.Split(trimmed, ":")
		if len(parts) == 2 {
			emojiID := parts[1]
			emojiURL := fmt.Sprintf("https://cdn.discordapp.com/emojis/%s.png", emojiID)

			analysis, err := h.discordSvc.GetImageAnalyzerService().AnalyzeImageFromURL(context.Background(), emojiURL)
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Failed to analyze emoji: %v", err))
				return
			}

			// Format and send the results
			response := fmt.Sprintf("Analysis for %s:\nKeywords: %s\nSentiment: %.2f", emojiArg, strings.Join(analysis.Keywords, ", "), analysis.Sentiment)
			s.ChannelMessageSend(m.ChannelID, response)
			return
		}
	}

	s.ChannelMessageSend(m.ChannelID, "Invalid custom emoji format. Please provide an emoji in the format <:name:id>.")
}
