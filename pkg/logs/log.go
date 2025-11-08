package logs

import (
	"context"
	"fmt"
	"time"

	"github.com/alexliesenfeld/health"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const DefaultHealthCheckLoggerName = "health_check"

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
	zapConfig := BaseZapConfig(zapcore.DebugLevel)

	zlogger, _ := zapConfig.Build()
	return zlogger.With(*zloggerCommonKey)
}

func NewTracingLogger() *otelzap.Logger {
	zapConfig := BaseZapConfig(zapcore.DebugLevel)

	zlogger, _ := zapConfig.Build()
	return otelzap.New(zlogger.With(zapcore.Field{Key: "tracing", Type: zapcore.BoolType, Interface: true}))
}

// LoggingInterceptor for HealthCheckers
func LoggingInterceptor(loggerName string) health.Interceptor {
	if loggerName == "" {
		loggerName = DefaultHealthCheckLoggerName
	}
	logger := NewZapLogger().Named(loggerName)
	return func(next health.InterceptorFunc) health.InterceptorFunc {
		return func(ctx context.Context, name string, state health.CheckState) health.CheckState {
			now := time.Now()
			result := next(ctx, name, state)
			logger.Info(fmt.Sprintf("checked component %s in %f seconds (result: %s)",
				name, time.Since(now).Seconds(), result.Status), zap.String("component", name))
			return result
		}
	}
}
