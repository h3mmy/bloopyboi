package providers

import (
	"context"

	"github.com/h3mmy/bloopyboi/internal/models"
	"go.uber.org/zap"
)

type ImageAnalyzerType string

// TODO: prefix with AnalyzerType_
const GoogleVision ImageAnalyzerType = "google_vision"
const AnalyzerType_Local ImageAnalyzerType = "local"

func NewImageAnalyzer(analyzerType ImageAnalyzerType) models.ImageAnalyzer {
	switch analyzerType {
	case "google_vision":
		googleVisionAnalyzer, err := NewGoogleVisionAnalyzer(context.Background())
		if err != nil {
			logger.Error("failed to create google vision analyzer", zap.Error(err))
			return nil
		}
		return googleVisionAnalyzer
	case AnalyzerType_Local:
		logger.Warn("local analyzer type not implemented yet")
		return nil
	default:
		return nil
	}

}
