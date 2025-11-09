package models

import (
	"context"
)

// ImageAnalysis contains the results of analyzing an image.
type ImageAnalysis struct {
	// Keywords is a list of descriptive tags for the image content.
	Keywords []string
	// Sentiment is a score from -1.0 (negative) to 1.0 (positive).
	Sentiment float64
}

// ImageAnalyzer defines the interface for a service that can analyze
// an image and return keywords and sentiment.
type ImageAnalyzer interface {
	// AnalyzeImageFromURL downloads an image from a URL and analyzes its content.
	AnalyzeImageFromURL(ctx context.Context, url string) (*ImageAnalysis, error)
}
