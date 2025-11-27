//nolint:all
package handlers

import (
	"context"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/h3mmy/bloopyboi/bot/services"
	"github.com/h3mmy/bloopyboi/internal/models"
	"github.com/h3mmy/bloopyboi/pkg/logs"
	"go.uber.org/zap"
)

type ImageAnalysisHandler struct {
	meta             models.BloopyMeta
	logger           *zap.Logger
	imageAnalyzerSvc *services.ImageAnalyzerService
	discordSvc       *services.DiscordService
}

func NewImageAnalysisHandler(imageAnalyzerSvc *services.ImageAnalyzerService, discordSvc *services.DiscordService) *ImageAnalysisHandler {
	return &ImageAnalysisHandler{
		meta:             models.NewBloopyMeta(),
		logger:           logs.NewZapLogger().Named("image_analysis_handler"),
		imageAnalyzerSvc: imageAnalyzerSvc,
		discordSvc:       discordSvc,
	}
}

// ProcessGuildEmojis should not be used anywhere for the time being
func (i *ImageAnalysisHandler) ProcessGuildEmojis(ctx context.Context, emoji []*discordgo.Emoji) error {
	ctx = context.WithValue(ctx, "handler_id", i.meta.Id)
	logr := i.logger.With(zap.Any("context", ctx))
	logr.Debug("processing guild emojis", zap.Int("count", len(emoji)))
	for _, e := range emoji {
		ctxf := context.WithValue(ctx, "emoji_id", e.ID)
		logr.Debug("processing emoji", zap.String("emoji_id", e.ID))
		analysis, err := i.AnalyzeDiscordEmoji(ctxf, e)
		if err != nil {
			logr.Error("failed to process emoji", zap.String("emoji_id", e.ID), zap.Error(err))
			continue
		}

		err = i.discordSvc.SaveDiscordEmojiWithAnalysis(ctx, analysis)
		if err != nil {
			logr.Error("failed to persist emoji analysis", zap.String("emoji_id", e.ID), zap.Error(err))
			continue
		}
	}
	return nil
}

func (i *ImageAnalysisHandler) AnalyzeDiscordEmoji(ctx context.Context, emoji *discordgo.Emoji) (analysis *models.DiscordEmojiAnalysisResult, err error) {

	imageURL := fmt.Sprintf("https://cdn.discordapp.com/emojis/%s.png", emoji.ID)
	if emoji.Animated {
		imageURL = fmt.Sprintf("https://cdn.discordapp.com/emojis/%s.gif", emoji.ID)
	}

	imageAnalysis, err := i.imageAnalyzerSvc.AnalyzeImageFromURL(ctx, imageURL)
	if err != nil {
		i.logger.Error("failed to analyze emoji image", zap.String("emoji_id", emoji.ID), zap.Error(err))
		return nil, err
	}
	return &models.DiscordEmojiAnalysisResult{
		ImageURI:       imageURL,
		Emoji:          emoji,
		AnalysisResult: imageAnalysis,
	}, nil
}
