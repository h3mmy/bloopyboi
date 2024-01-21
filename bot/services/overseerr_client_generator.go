package services

import (
	"net/http"

	overseerr_go "github.com/h3mmy/bloopyboi/bot/clients/overseerr"
	"github.com/h3mmy/bloopyboi/bot/internal/models"
)

type OverseerrClientGenerator struct {
	models.ClientGenerator
}

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

func (s *OverseerrClientGenerator) generateClient() *overseerr_go.APIClient {
	config := s.parseConfigFromArgs(s.Args)
	return overseerr_go.NewAPIClient(config)
}
