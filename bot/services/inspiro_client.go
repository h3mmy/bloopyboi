package services

import (
	"fmt"
	"io"
	"net/http"

	"github.com/h3mmy/bloopyboi/pkg/config"
	log "github.com/h3mmy/bloopyboi/pkg/logs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger              = log.NewZapLogger()
	InspiroFeatureName  = "inspiro"
	InspiroAPIKey       = "api_url"
	InspiroBackupURLKey = "backup_image_link"
)

// Basically a static var for this 'Object'
type InspiroConfig struct {
	API_url           string
	Logger            *zap.Logger
	Backup_image_link string
}

// Unmarshal config from file
func GetInspiroConfig() InspiroConfig {
	AppConfig := config.GetConfig()
	inspiroCfg, err := AppConfig.GetFeatureConfig(InspiroFeatureName)
	if err != nil {
		logger.Sugar().Error("Error loading FeatureConfig", err)
	}
	return InspiroConfig{
		API_url:           inspiroCfg.Data[InspiroAPIKey],
		Logger:            logger,
		Backup_image_link: inspiroCfg.Data[InspiroBackupURLKey]}
}

// should add uri validation
type InspiroClient struct {
	config InspiroConfig
	logger *zap.Logger
}

// 'Constructs' InspiroClient with declared Config
func NewInspiroClientWithConfig(myConfig InspiroConfig) *InspiroClient {
	lgr := log.NewZapLogger().With(zapcore.Field{
		Key:    ServiceLoggerFieldKey,
		Type:   zapcore.StringType,
		String: "inspiro",
	})
	return &InspiroClient{
		config: myConfig,
		logger: lgr,
	}
}

// Abstracted 'Constructor'
func NewInspiroClient() *InspiroClient {
	return NewInspiroClientWithConfig(GetInspiroConfig())
}

// returns raw uri as string without validation
func (inspiroService *InspiroClient) GetInspiro() string {

	image_link, err := http.Get(inspiroService.config.API_url)

	if err != nil {
		return err.Error()
	}
	defer func() {
		if err := image_link.Body.Close(); err != nil {
			inspiroService.logger.Error("failed to close http response body", zap.Error(err))
		}
	}()

	result, err := io.ReadAll(image_link.Body)
	if err != nil {
		inspiroService.logger.Sugar().Error("IO Error while reading body", err)
		return err.Error()
	}
	inspiroService.logger.Debug(fmt.Sprintf("Got Link %s", result))
	return string(result)
}
