package arr

import (
	"fmt"
	"strings"

	"github.com/h3mmy/bloopyboi/bot/internal/config"
	"github.com/h3mmy/bloopyboi/bot/internal/log"
	"github.com/h3mmy/bloopyboi/bot/internal/models"
	"go.uber.org/zap"
	"golift.io/starr"
	"golift.io/starr/lidarr"
	"golift.io/starr/prowlarr"
	"golift.io/starr/radarr"
	"golift.io/starr/readarr"
	"golift.io/starr/sonarr"
)

type ArrClientMap map[string]starr.APIer
type ArrClientRegister map[starr.App]ArrClientMap

type ArrClientRegistry struct {
	meta models.BloopyMeta
	logger *zap.Logger
	registry ArrClientRegister
}

func NewArrClientRegistry(controllerName string) *ArrClientRegistry {
	mta := models.NewBloopyMeta(controllerName)
	lgr := log.NewZapLogger().Named("arr_client_registry").With
	return &ArrClientRegistry{
		meta: mta,
		logger: lgr,
		registry: make(map[starr.App]ArrClientMap),
	}
}

func BuildArrClient(cfg *config.ArrClientConfig) (starr.APIer, error) {
	params := cfg.ToParams()
	starrConfig := starr.New(
		params[config.ApiKey],
		params[config.AppURL],
		starr.DefaultTimeout,
	)
	switch strings.ToLower(cfg.Type) {
	case starr.Sonarr.Lower():
		return sonarr.New(starrConfig), nil
	case starr.Radarr.Lower():
		return radarr.New(starrConfig), nil
	case starr.Readarr.Lower():
		return readarr.New(starrConfig), nil
	case starr.Prowlarr.Lower():
		return prowlarr.New(starrConfig), nil
	case starr.Lidarr.Lower():
		return lidarr.New(starrConfig), nil
	}
	return nil, fmt.Errorf("Could not build client %s of type: %s", cfg.Name, cfg.Type)
}

func (s *ArrClientRegistry) AddClient(cfg *config.ArrClientConfig) error {
	logger := s.logger.With(zap.String("clientName", cfg.Name)).With(zap.String("clientType", cfg.Type))
	key := starr.App(cfg.Type)
	if val, ok := s.registry[key]; ok {
		logger.Debug("register exists for client type.")
		client, err := BuildArrClient(cfg)
		if err != nil {
			logger.Error("error building client", zap.Error(err))
			return err
		}
		if _, ok := val[cfg.Name]; ok {
			logger.Warn("entry already exists with client Name. Existing entry will be overwritten.")
		}
		val[cfg.Name] = client
		logger.Debug("added client to register")
	} else {
		logger.Debug("no existing entries for client type. adding new entry")
		cMap := make(map[string]starr.APIer)
		client, err := BuildArrClient(cfg)
		if err != nil {
			logger.Error("error building client", zap.Error(err))
			return err
		}
		cMap[cfg.Name] = client
		logger.Debug("added client to registry")
		s.registry[key] = cMap
		logger.Debug("Added new type to register")
	}

	return nil
}
