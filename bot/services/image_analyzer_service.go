package services

import (
	"context"
	"fmt"

	"github.com/h3mmy/bloopyboi/internal/models"
	log "github.com/h3mmy/bloopyboi/pkg/logs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ImageAnalyzerService struct {
	meta          models.BloopyMeta
	logger        *zap.Logger
	imageAnalyzer models.ImageAnalyzer
}

func NewImageAnalyzerService(imageAnalyzer models.ImageAnalyzer) *ImageAnalyzerService {
	meta := models.NewBloopyMeta()
	lgr := log.NewZapLogger().
		Named("image_analyzer_service").
		With(
			zapcore.Field{Type: zapcore.StringType, Key: "bloopy_id", String: meta.Id.String()},
			zapcore.Field{Type: zapcore.TimeFullType, Key: "created_at", Interface: meta.CreatedAt},
			zapcore.Field{Type: zapcore.StringType, Key: ServiceLoggerFieldKey, String: "image_analyzer_service"},
		)
	return &ImageAnalyzerService{
		meta:          meta,
		logger:        lgr,
		imageAnalyzer: imageAnalyzer,
	}
}

func (s *ImageAnalyzerService) AnalyzeImageFromURL(ctx context.Context, url string) (*models.ImageAnalysis, error) {
	// This is a dummy implementation.
	// In the future, this could call a real image analysis service.
	fmt.Printf("Analyzing image from URL: %s\n", url)
	return s.imageAnalyzer.AnalyzeImageFromURL(ctx, url)
}
