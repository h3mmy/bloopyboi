package arr

import (
	"fmt"
	"strings"

	"github.com/h3mmy/bloopyboi/bot/internal/config"
	"golift.io/starr"
	"golift.io/starr/lidarr"
	"golift.io/starr/prowlarr"
	"golift.io/starr/radarr"
	"golift.io/starr/readarr"
	"golift.io/starr/sonarr"
)

type ArrClientMap map[string]starr.APIer
type ArrClientRegistry map[starr.App]ArrClientMap

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

func BuildArrClientRegistry(cfg *config.AppConfig) ArrClientRegistry {
	registry:= make(map[starr.App]ArrClientMap)
	for _, arrCfg := range *cfg.Arrs {
		key := starr.App(arrCfg.Type)
		if val, ok := registry[key]; ok {
			client, err := BuildArrClient(&arrCfg)
			if err != nil {
				// log
			}
			val[arrCfg.Name] = client
		} else {
			cMap := make(map[string]starr.APIer)
			client, err := BuildArrClient(&arrCfg)
			if err != nil {
				// log
			}
			cMap[arrCfg.Name] = client
			registry[key] = cMap
		}
	}
	return registry
}
