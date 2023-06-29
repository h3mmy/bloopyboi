package services

import (
	"gitlab.com/h3mmy/bloopyboi/bot/internal/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type InspiroClient struct {
	inspiroService *InspiroService
	logger *zap.Logger
}

// Creates New InspiroClient with specified Service
func NewInspiroHttpClient(inspiro *InspiroService) *InspiroClient {
	lgr := log.NewZapLogger().With(zapcore.Field{
		Key: ServiceLoggerFieldKey,
		Type: zapcore.StringType,
		String: "inspiro",
	})
	return &InspiroClient{
		inspiroService: inspiro,
		logger: lgr,
	}
}

// Creates New InspiroClient with default service
func NewInspiroClient() *InspiroClient {
	return NewInspiroHttpClient(NewInspiroService())
}

func (ic *InspiroClient) GetInspiroImageURL() string {
	ic.logger.Debug("Getting Inspiro Image")
	return ic.inspiroService.GetInspiro()
}

func (ic *InspiroClient) GetService() *InspiroService {
	return ic.inspiroService
}
