package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var zloggerCommonKey = &zap.Field{Key: "group", Type: zapcore.StringType, String: "common"}

// Setup New Common Zap Logger
func BaseZapConfig(level zapcore.Level) *zap.Config {
	zapconfig := zap.NewDevelopmentConfig()
	zapconfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	zapconfig.Level = zap.NewAtomicLevelAt(level)
	return &zapconfig
}

// Generate new zap logger
func NewZapLogger() *zap.Logger {
	zapConfig := zap.NewDevelopmentConfig()

	zlogger, _ := zapConfig.Build()
	return zlogger.With(*zloggerCommonKey)
}
