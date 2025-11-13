package services

import (
	"net/http"

	overseerr_go "github.com/devopsarr/overseerr-go/overseerr"
	"github.com/h3mmy/bloopyboi/internal/models"
)

// OverseerrClientGenerator is a client generator for the Overseerr API.
type OverseerrClientGenerator struct {
	models.ClientGenerator
}

// parseConfigFromArgs parses the configuration from the arguments.
func (s *OverseerrClientGenerator) parseConfigFromArgs(args map[string]interface{}) *overseerr_go.Configuration {
	host := args["host"].(string)
	scheme := args["scheme"].(string)
	insecure := args["insecure"].(bool)
	overseerrConfig := overseerr_go.NewConfiguration()
	overseerrConfig.Host = host
	overseerrConfig.Scheme = scheme
	overseerrConfig.HTTPClient = &http.Client{
		Transport: GetTLSTransport(insecure),
	}
	return overseerrConfig
}

// generateClient generates a new Overseerr API client.
func (s *OverseerrClientGenerator) generateClient() *overseerr_go.APIClient {
	config := s.parseConfigFromArgs(s.Args)
	return overseerr_go.NewAPIClient(config)
}
