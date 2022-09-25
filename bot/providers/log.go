package providers

import (
	"gitlab.com/h3mmy/bloopyboi/bot/internal/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	ServiceLoggerFieldKey = "service_name"
)

var (
	logger = log.NewZapLogger()
)

func NewZapLogger() *zap.Logger {
	l, err := zapcore.ParseLevel(GetLogLevel())
	if err != nil {
		logger.Error("error parsing loglevel. Defaulting to InfoLevel.")
		l = zapcore.InfoLevel
	}
	cfg := log.BaseZapConfig(l)
	lgr, _ := cfg.Build()

	return lgr
}
