package providers

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"

	vision "cloud.google.com/go/vision/apiv1"
	"github.com/h3mmy/bloopyboi/internal/models"
	visionpb "google.golang.org/genproto/googleapis/cloud/vision/v1"

	"google.golang.org/api/option"
)

// GoogleVisionAnalyzer is an ImageAnalyzer that uses the Google Cloud Vision API.
type GoogleVisionAnalyzer struct {
	client *vision.ImageAnnotatorClient
}

// NewGoogleVisionAnalyzer creates a new analyzer client.
// It expects Google Application Credentials to be configured in the environment.
func NewGoogleVisionAnalyzer(ctx context.Context, opts ...option.ClientOption) (*GoogleVisionAnalyzer, error) {
	client, err := vision.NewImageAnnotatorClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create vision client: %w", err)
	}
	return &GoogleVisionAnalyzer{client: client}, nil
}

// AnalyzeImageFromURL implements the ImageAnalyzer interface.
func (a *GoogleVisionAnalyzer) AnalyzeImageFromURL(ctx context.Context, url string) (*models.ImageAnalysis, error) {
	// 1. Download the image
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to download image from %s: %w", url, err)
	}
	defer resp.Body.Close()

	imageBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read image body: %w", err)
	}

	// 2. Call the Vision API
	image, err := vision.NewImageFromReader(bytes.NewReader(imageBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to create vision image: %w", err)
	}

	// 3. Perform label and sentiment detection (simplified for this example)
	labels, err := a.client.DetectLabels(ctx, image, nil, 10)
	if err != nil {
		return nil, fmt.Errorf("failed to detect labels: %w", err)
	}

	analysis := &models.ImageAnalysis{
		Keywords:  []string{},
		Sentiment: 0.0, // Placeholder for sentiment logic
	}

	for _, label := range labels {
		analysis.Keywords = append(analysis.Keywords, label.Description)
	}

	// Sentiment can be derived from face annotations, safe search, or dominant colors.
	// This logic can be added here.
	safeSearch, err := a.client.DetectSafeSearch(ctx, image, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to detect safe search: %w", err)
	}
	analysis.Sentiment = calculateSentiment(safeSearch)

	return analysis, nil
}

// calculateSentiment derives a sentiment score from safe search annotations.
func calculateSentiment(safeSearch *visionpb.SafeSearchAnnotation) float64 {
	// This is a simple example. A more sophisticated model could be used.
	// We'll assign negative scores for "bad" categories and positive for "good".
	// The likelihoods are enums from UNKNOWN to VERY_LIKELY.
	var score float64

	// Negative contributors
	score -= float64(safeSearch.Adult) * 0.25
	score -= float64(safeSearch.Spoof) * 0.25
	score -= float64(safeSearch.Medical) * 0.1
	score -= float64(safeSearch.Violence) * 0.25
	score -= float64(safeSearch.Racy) * 0.15

	// Normalize the score to be between -1.0 and 1.0
	// The max negative score is -5 (VERY_LIKELY for all) * 0.25 (avg weight) = -1.25 -> not quite
	// Let's do it differently. The enum values are 0-5.
	// Let's normalize the score to be within -1 and 1
	// The maximum possible negative score is -(5*0.25 + 5*0.25 + 5*0.1 + 5*0.25 + 5*0.15) = -(1.25 + 1.25 + 0.5 + 1.25 + 0.75) = -5
	// So we can divide by 5 to normalize.
	return score / 5.0
}

// Close closes the underlying client connection.
func (a *GoogleVisionAnalyzer) Close() error {
	return a.client.Close()
}
