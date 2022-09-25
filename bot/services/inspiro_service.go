package services

import (
	"fmt"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
	"gitlab.com/h3mmy/bloopyboi/bot/internal/config"
	"gitlab.com/h3mmy/bloopyboi/bot/internal/log"
	"gitlab.com/h3mmy/bloopyboi/bot/providers"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger              = log.New()
	InspiroFeatureName  = "inspiro"
	InspiroAPIKey       = "api_url"
	InspiroBackupURLKey = "backup_image_link"
)

// Basically a static var for this 'Object'
type InspiroConfig struct {
	API_url           string
	Logger            *logrus.Logger
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
		API_url:           inspiroCfg.Data[InspiroAPIKey],
		Logger:            logger,
		Backup_image_link: inspiroCfg.Data[InspiroBackupURLKey]}
}

// should add uri validation
type InspiroService struct {
	config InspiroConfig
	logger *zap.Logger
}

// 'Constructs' InspiroService with declared Config
func NewInspiroServiceWithConfig(myConfig InspiroConfig) *InspiroService {
	lgr := log.NewZapLogger().With(zapcore.Field{
		Key: providers.ServiceLoggerFieldKey,
	})
	return &InspiroService{
		config: myConfig,
		logger: lgr,
	}
}

// Abstracted 'Constructor'
func NewInspiroService() *InspiroService {
	return NewInspiroServiceWithConfig(GetInspiroConfig())
}

// returns raw uri as string without validation
func (inspiroService *InspiroService) GetInspiro() string {

	image_link, err := http.Get(inspiroService.config.API_url)

	if err != nil {
		return err.Error()
	}
	defer image_link.Body.Close()

	result, err := io.ReadAll(image_link.Body)
	if err != nil {
		inspiroService.logger.Sugar().Error("IO Error while reading body", err)
		return err.Error()
	}
	inspiroService.logger.Debug(fmt.Sprintf("Got Link %s", result))
	return string(result)
}
