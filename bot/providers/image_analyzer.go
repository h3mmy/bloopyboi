package providers

import (
	"context"

	"github.com/h3mmy/bloopyboi/internal/models"
	"go.uber.org/zap"
)

type ImageAnalyzerType string

const GoogleVision ImageAnalyzerType = "google_vision"

func NewImageAnalyzer(analyzerType ImageAnalyzerType) models.ImageAnalyzer {
	switch analyzerType {
	case "google_vision":
		googleVisionAnalyzer, err := NewGoogleVisionAnalyzer(context.Background())
		if err != nil {
			logger.Error("failed to create google vision analyzer", zap.Error(err))
			return nil
		}
		return googleVisionAnalyzer
	default:
		return nil
	}

}
