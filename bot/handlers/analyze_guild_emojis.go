//nolint:all
package handlers

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/bwmarrin/discordgo"
	"github.com/h3mmy/bloopyboi/bot/services"
	"github.com/h3mmy/bloopyboi/ent"
	"github.com/h3mmy/bloopyboi/ent/emojikeywordscore"
	"github.com/h3mmy/bloopyboi/ent/keyword"
	"github.com/h3mmy/bloopyboi/internal/models"
	"github.com/h3mmy/bloopyboi/pkg/database"
	"github.com/h3mmy/bloopyboi/pkg/logs"
	"go.uber.org/zap"
)

type ImageAnalysisHandler struct {
	meta             models.BloopyMeta
	logger           *zap.Logger
	imageAnalyzerSvc *services.ImageAnalyzerService
	db               *ent.Client
}

func NewImageAnalysisHandler(imageAnalyzerSvc *services.ImageAnalyzerService, db *ent.Client) *ImageAnalysisHandler {
	return &ImageAnalysisHandler{
		meta:             models.NewBloopyMeta(),
		logger:           logs.NewZapLogger().Named("image_analysis_handler"),
		imageAnalyzerSvc: imageAnalyzerSvc,
		db:               db,
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

		err = i.SaveDiscordEmojiWithAnalysis(ctx, analysis)
		if err != nil {
			logr.Error("failed to persist emoji analysis", zap.String("emoji_id", e.ID), zap.Error(err))
			continue
		}
	}
	return nil
}

func (i *ImageAnalysisHandler) SaveDiscordEmojiWithAnalysis(ctx context.Context, analysis *models.DiscordEmojiAnalysisResult) error {
	logr := i.logger.With(zap.Any("context", ctx))
	logr.Debug("saving discord emoji with analysis", zap.Any("analysis", analysis))
	tx, err := i.db.Tx(ctx)
	if err != nil {
		return fmt.Errorf("starting a transaction: %w", err)
	}

	emjID, err := tx.Emoji.
		Create().
		SetEmojiID(analysis.Emoji.ID).
		SetAnimated(analysis.Emoji.Animated).
		SetName(analysis.Emoji.Name).
		SetImageURI(analysis.ImageURI).
		SetAdultLikelihood(int(analysis.AnalysisResult.SafeSearchAnalysis.Adult)).
		SetRacyLikelihood(int(analysis.AnalysisResult.SafeSearchAnalysis.Racy)).
		SetSpoofLikelihood(int(analysis.AnalysisResult.SafeSearchAnalysis.Spoof)).
		SetViolenceLikelihood(int(analysis.AnalysisResult.SafeSearchAnalysis.Violence)).
		SetMedicalLikelihood(int(analysis.AnalysisResult.SafeSearchAnalysis.Medical)).
		OnConflict(sql.ConflictColumns("emoji_id")).
		UpdateNewValues().
		ID(ctx)

	if err != nil {
		logr.Error("failed to create emoji", zap.String("emoji_id", analysis.Emoji.ID), zap.Error(err))
		return database.Rollback(tx, err)
	}

	keywordBuilders := make([]*ent.KeywordCreate, len(analysis.AnalysisResult.Labels))
	keywordScoreMap := make(map[string]models.EntityLabelAnnotation)
	keywordList := make([]string, len(analysis.AnalysisResult.Labels))

	for i, keyword := range analysis.AnalysisResult.Labels {
		keywordBuilders[i] = tx.Keyword.Create().
			SetKeyword(keyword.Description).
			AddEmojiIDs(emjID)
		keywordScoreMap[keyword.Description] = keyword
		keywordList[i] = keyword.Description
	}

	err = tx.Keyword.CreateBulk(
		keywordBuilders...,
	).OnConflictColumns(keyword.FieldKeyword).
		UpdateNewValues().
		Exec(ctx)

	if err != nil {
		logr.Error("failed to upsert keywords", zap.String("emoji_id", analysis.Emoji.ID), zap.Error(err))
		return database.Rollback(tx, err)
	}

	keywords, err := tx.Keyword.
		Query().
		Where(keyword.KeywordIn(keywordList...)).
		All(ctx)

	if err != nil {
		logr.Error("failed to query upserted keywords", zap.Error(err))
		return database.Rollback(tx, err)
	}
	keywordScoreBuilders := make([]*ent.EmojiKeywordScoreCreate, len(keywords))
	for i, keyw := range keywords {
		keywordScoreBuilders[i] = tx.EmojiKeywordScore.Create().
			SetEmojiID(emjID).
			SetKeywordID(keyw.ID).
			SetScore(keywordScoreMap[keyw.Keyword].Score).
			SetTopicality(keywordScoreMap[keyw.Keyword].Topicality)
	}

	err = tx.EmojiKeywordScore.CreateBulk(
		keywordScoreBuilders...,
	).
		OnConflictColumns(
			emojikeywordscore.FieldEmojiID,
			emojikeywordscore.FieldKeywordID,
		).
		UpdateNewValues().
		Exec(ctx)

	if err != nil {
		logr.Error("failed to upsert keyword scores", zap.Error(err))
		return database.Rollback(tx, err)
	}

	return tx.Commit()
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
