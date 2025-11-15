package models

import (
	"context"
	"sort"
)

type Likelihood int32

const (
	// Unknown likelihood.
	Likelihood_UNKNOWN Likelihood = 0
	// It is very unlikely.
	Likelihood_VERY_UNLIKELY Likelihood = 1
	// It is unlikely.
	Likelihood_UNLIKELY Likelihood = 2
	// It is possible.
	Likelihood_POSSIBLE Likelihood = 3
	// It is likely.
	Likelihood_LIKELY Likelihood = 4
	// It is very likely.
	Likelihood_VERY_LIKELY Likelihood = 5
)

type SafeSearchAnnotation struct {

	// Represents the adult content likelihood for the image. Adult content may
	// contain elements such as nudity, pornographic images or cartoons, or
	// sexual activities.
	Adult Likelihood `json:"adult,omitempty"`
	// Spoof likelihood. The likelihood that an modification
	// was made to the image's canonical version to make it appear
	// funny or offensive.
	Spoof Likelihood `json:"spoof,omitempty"`
	// Likelihood that this is a medical image.
	Medical Likelihood `json:"medical,omitempty"`
	// Likelihood that this image contains violent content. Violent content may
	// include death, serious harm, or injury to individuals or groups of
	// individuals.
	Violence Likelihood `json:"violence,omitempty"`
	// Likelihood that the request image contains racy content. Racy content may
	// include (but is not limited to) skimpy or sheer clothing, strategically
	// covered nudity, lewd or provocative poses, or close-ups of sensitive
	// body areas.
	Racy Likelihood `json:"racy,omitempty"`
	// contains filtered or unexported fields
}

type EntityLabelAnnotation struct {
	// The language code for the locale in which the entity textual
	// `description` is expressed.
	Locale string `json:"locale,omitempty"`
	// Entity textual description, expressed in its `locale` language.
	Description string `json:"description,omitempty"`
	// Overall score of the result. Range [0, 1].
	Score float32 `json:"score,omitempty"`
	// The relevancy of the ICA (Image Content Annotation) label to the
	// image. For example, the relevancy of "tower" is likely higher to an image
	// containing the detected "Eiffel Tower" than to an image containing a
	// detected distant towering building, even though the confidence that
	// there is a tower in each image may be the same. Range [0, 1].
	Topicality float32 `json:"topicality,omitempty"`
	// contains filtered or unexported fields
}

// ImageAnalysis contains the results of analyzing an image.
type ImageAnalysis struct {
	Labels             []EntityLabelAnnotation
	SafeSearchAnalysis *SafeSearchAnnotation
}

func (ia *ImageAnalysis) GetKeywords() []string {
	var keywords []string
	for _, label := range ia.Labels {
		keywords = append(keywords, label.Description)
	}
	return keywords
}

func (ia *ImageAnalysis) GetKeywordsSortedByScore() []string {
	var keywords []string
	sortingLabels := ia.Labels
	sort.Slice(ia.Labels, func(i, j int) bool {
		return sortingLabels[i].Score > sortingLabels[j].Score
	})
	for _, label := range sortingLabels {
		keywords = append(keywords, label.Description)
	}
	return keywords
}

// ImageAnalyzer defines the interface for a service that can analyze
// an image and return keywords and sentiment.
type ImageAnalyzer interface {
	// AnalyzeImageFromURL downloads an image from a URL and analyzes its content.
	AnalyzeImageFromURL(ctx context.Context, url string) (*ImageAnalysis, error)
}
