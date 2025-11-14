package providers

import (
	"context"
	"fmt"
	"io"
	"net/http"

	vision "cloud.google.com/go/vision/v2/apiv1"
	visionpb "cloud.google.com/go/vision/v2/apiv1/visionpb"
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
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil && err == nil {
			err = closeErr
		}
	}()

	imageBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read image body: %w", err)
	}

	// 2. Call the Vision API
	image := &visionpb.Image{
		Content: imageBytes,
	}

	// 3. Perform label and sentiment detection (simplified for this example)
	req := &visionpb.BatchAnnotateImagesRequest{
		Requests: []*visionpb.AnnotateImageRequest{
			{
				Image: image,
				Features: []*visionpb.Feature{
					{
						Type:       visionpb.Feature_LABEL_DETECTION,
						MaxResults: 10,
					},
					{
						Type: visionpb.Feature_SAFE_SEARCH_DETECTION,
					},
				},
			},
		},
	}

	res, err := a.client.BatchAnnotateImages(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to batch annotate images: %w", err)
	}

	if len(res.Responses) == 0 {
		return nil, fmt.Errorf("no responses from vision api")
	}

	response := res.Responses[0]

	analysis := &models.ImageAnalysis{
		Keywords:  []string{},
		Sentiment: 0.0,
	}

	for _, label := range response.LabelAnnotations {
		analysis.Keywords = append(analysis.Keywords, label.Description)
	}

	analysis.Sentiment = calculateSentiment(response.SafeSearchAnnotation)

	return analysis, nil
}

// calculateSentiment derives a sentiment score from safe search annotations.
func calculateSentiment(safeSearch *visionpb.SafeSearchAnnotation) float64 {
	var score float64

	score -= float64(safeSearch.Adult) * 0.25
	score -= float64(safeSearch.Spoof) * 0.25
	score -= float64(safeSearch.Medical) * 0.1
	score -= float64(safeSearch.Violence) * 0.25
	score -= float64(safeSearch.Racy) * 0.15

	return score / 5.0
}

// Close closes the underlying client connection.
func (a *GoogleVisionAnalyzer) Close() error {
	return a.client.Close()
}
