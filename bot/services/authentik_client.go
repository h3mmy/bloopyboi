package services

import (
	"fmt"
	"net/http"
	"net/url"

	httptransport "github.com/go-openapi/runtime/client"
	"gitlab.com/h3mmy/bloopyboi/bot/internal/models"
	api "goauthentik.io/api/v3"
)

type AuthentikClientGenerator struct {
	models.ClientGenerator
}

// Generates Authentik Client
func (s *AuthentikClientGenerator) generateClient() *api.APIClient {

	apiURL := s.Args["url"].(string)
	token := s.Args["token"].(string)
	insecure := s.Args["insecure"].(bool)

	akURL, err := url.Parse(apiURL)
	if err != nil {
		panic(err)
	}

	config := api.NewConfiguration()
	config.UserAgent = fmt.Sprintf("serviceaccount:%s:%s", "bloopyboi", "authentik")
	config.Host = akURL.Host
	config.Scheme = akURL.Scheme

	config.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %s", token))
	config.HTTPClient = &http.Client{
		Transport: GetTLSTransport(insecure),
	}
	apiClient := api.NewAPIClient(config)

	return apiClient
}

// GetTLSTransport Get a TLS transport instance, that skips verification if configured via environment variables.
func GetTLSTransport(insecure bool) http.RoundTripper {
	tlsTransport, err := httptransport.TLSTransport(httptransport.TLSClientOptions{
		InsecureSkipVerify: insecure,
	})
	if err != nil {
		panic(err)
	}
	return tlsTransport
}
