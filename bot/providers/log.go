package providers

import (
	"gitlab.com/h3mmy/bloopyboi/bot/internal/log"
)

const (
	ServiceLoggerFieldKey = "service_name"
)

var (
	logger = log.New()
	CommonLogger = log.DefaultBloopyFieldLogger()
)
