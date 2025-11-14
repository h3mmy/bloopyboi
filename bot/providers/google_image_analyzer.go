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

	// 3. Perform label and safe search detection
	req := &visionpb.BatchAnnotateImagesRequest{
		Requests: []*visionpb.AnnotateImageRequest{
			{
				Image: image,
				Features: []*visionpb.Feature{
					{
						Type:       visionpb.Feature_LABEL_DETECTION,
						MaxResults: 15,
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
		Labels: []models.EntityLabelAnnotation{},
		SafeSearchAnalysis: &models.SafeSearchAnnotation{},
	}

	for _, label := range response.LabelAnnotations {
		analysis.Labels = append(analysis.Labels, models.EntityLabelAnnotation{
			Locale:       label.GetLocale(),
			Description:  label.GetDescription(),
			Score:        label.GetScore(),
			Topicality:   label.GetTopicality(),
		})
	}

	analysis.SafeSearchAnalysis = &models.SafeSearchAnnotation{
		Adult:     models.Likelihood(response.SafeSearchAnnotation.Adult),
		Spoof:     models.Likelihood(response.SafeSearchAnnotation.Spoof),
		Medical:   models.Likelihood(response.SafeSearchAnnotation.Medical),
		Violence:  models.Likelihood(response.SafeSearchAnnotation.Violence),
		Racy:      models.Likelihood(response.SafeSearchAnnotation.Racy),
	}

	return analysis, nil
}

// Close closes the underlying client connection.
func (a *GoogleVisionAnalyzer) Close() error {
	return a.client.Close()
}
