package providers

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"

	vision "cloud.google.com/go/vision/apiv1"
	"github.com/h3mmy/bloopyboi/internal/models"
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

	return analysis, nil
}

// Close closes the underlying client connection.
func (a *GoogleVisionAnalyzer) Close() error {
	return a.client.Close()
}
