package util

import (
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
	"gitlab.com/h3mmy/bloopyboi/bot/internal/config"
	"gitlab.com/h3mmy/bloopyboi/bot/internal/log"
)


var logger = log.New()
var InspiroFeatureName = "inspiro"
var InspiroAPIKey = "api_url"
var InspiroBackupURLKey = "backup_image_link"

// Basically a static var for this 'Object'
type InspiroConfig struct {
	API_url string
	Logger *logrus.Logger
	Backup_image_link string
}

// Unmarshal config from file
func GetInspiroConfig() InspiroConfig {
	botConfig, err := config.GetConfig()
	if err != nil {
		logger.Error("Error loading config", err)
	}
	inspiroCfg, err := botConfig.GetFeatureConfig(InspiroFeatureName)
	if err != nil {
		logger.Error("Error loading FeatureConfig", err)
	}
	return InspiroConfig{
		API_url: inspiroCfg.Data[InspiroAPIKey],
		Logger: logger,
		Backup_image_link: inspiroCfg.Data[InspiroBackupURLKey]}
}

// should add uri validation
type InspiroClient struct {
	config InspiroConfig
}

// 'Constructs' InspiroClient with declared Config
func NewInspiroClient(myConfig InspiroConfig) *InspiroClient {
	return &InspiroClient{
		config: myConfig,
	}
}

// 'Constructs' InspiroClient with transparent Config
func NewInspiroClientWithURI(apiUrl string, logger *logrus.Logger, backupLink string) *InspiroClient {
	if logger != nil {
		return &InspiroClient{
			config: InspiroConfig{API_url: apiUrl, Logger: logger, Backup_image_link: backupLink},
		}
	}
	return &InspiroClient{
		config: InspiroConfig{API_url: apiUrl, Logger: log.New(), Backup_image_link: backupLink},
	}
}

// returns raw uri as string without validation
func (inspiroClient *InspiroClient) getInspiro() string {

	image_link, err := http.Get(inspiroClient.config.API_url)
	defer image_link.Body.Close()

	if err != nil {
		result, err2 := io.ReadAll(image_link.Body)
		if err2 != nil {
			return string(result)
		}
		return err2.Error()
	}
	return err.Error()
}
