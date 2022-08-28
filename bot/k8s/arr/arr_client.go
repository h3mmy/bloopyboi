package arr

import (
	"golift.io/starr"
	"golift.io/starr/sonarr"
	"golift.io/starr/radarr"
	"golift.io/starr/prowlarr"
	"golift.io/starr/readarr"
	"gitlab.com/h3mmy/bloopyboi/bot/internal/config"
)

var (
	supportedFeatures = []BloopyArrFeature{
		BloopyArrFeature{Name: "sonarr"},
		BloopyArrFeature{Name: "sonarrUHD"},
		BloopyArrFeature{Name: "radarr"},
		BloopyArrFeature{Name: "radarrUHD"},
		BloopyArrFeature{Name: "readarr"},
		BloopyArrFeature{Name: "readarrAudio"},
		BloopyArrFeature{Name: "prowlarr"},
	}

	featureMap = map[string]BloopyArrFeature{
		"sonarr": supportedFeatures[0],
		"sonarrUHD": supportedFeatures[1],
		"radarr": supportedFeatures[2],
		"radarrUHD": supportedFeatures[3],
		"readarr": supportedFeatures[4],
		"readarrAudio": supportedFeatures[5],
		"prowlarr": supportedFeatures[6],
	}
)

type BloopyArrFeature struct {
	Name				string
	ClientGenerator		func (params map[string]string) *starr.App
}

type ArrFeature interface {
	GetFeature()							*BloopyArrFeature
	GenerateClient(config.FeatureConfig)	*starr.App
}

type FeatureKeys struct {
	Sonarr			string
	Radarr			string
	SonarrUHD		string
	RadarrUHD		string
	Readarr			string
	ReadarrAudio	string
	Prowlarr		string
}

type BloopyArrClientSet map[string] *starr.App;

func NewArrClientSet(botConfig config.BotConfig) {

}

func GetBloopyArrClientSet(botConfig config.BotConfig) *BloopyArrClientSet {
	var clientSet BloopyArrClientSet
	for _, featName := range botConfig.GetConfiguredFeatureNames() {
		if isArrFeature(featName) {
			feat := featureMap[featName]
			featConfig, err := botConfig.GetFeatureConfig(featName)
			if err != nil {
				panic(err)
			}
			clientSet[featName] = feat.ClientGenerator(featConfig.Data)
		}
	}
	return &clientSet
}

func isArrFeature(featureName string) bool {
	if _, ok := featureMap[featureName]; ok {
		return true
	}
	return false
}

func GetSonarrClient(params map[string]string) *sonarr.Sonarr {
	starrConfig := starr.New(
		params["apiKey"],
		params["appURL"],
		starr.DefaultTimeout,
	)

	return sonarr.New(starrConfig)
}

func GetRadarrClient(params map[string]string) *radarr.Radarr {
	starrConfig := starr.New(
		params["apiKey"],
		params["appURL"],
		starr.DefaultTimeout,
	)

	return radarr.New(starrConfig)
}

func GetReadarrClient(params map[string]string) *readarr.Readarr {
	starrConfig := starr.New(
		params["apiKey"],
		params["appURL"],
		starr.DefaultTimeout,
	)

	return readarr.New(starrConfig)
}

func GetProwlarrClient(params map[string]string) *prowlarr.Prowlarr {
	starrConfig := starr.New(
		params["apiKey"],
		params["appURL"],
		starr.DefaultTimeout,
	)

	return prowlarr.New(starrConfig)
}
