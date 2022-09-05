package log

import (
	"os"

	"github.com/onrik/logrus/filename"
	"github.com/sirupsen/logrus"
)

// New function initialize logrus and return a new logger
// We use an abstraction so that our logs are consistent and if there's anything that needs change
// related to logs, we can just change here
func New() *logrus.Logger {
	// Filename is a hook for logrus that adds file name and line number to the log as well.
	// It's useful for indicating where the log was originated from
	filenameHook := filename.NewHook()

	log := &logrus.Logger{
		Hooks:     make(logrus.LevelHooks),
		Out:       os.Stdout,
		Formatter: &logrus.TextFormatter{},
		Level:     logrus.DebugLevel,
	}

	log.Hooks.Add(filenameHook)

	return log
}

// Set LogLevel from config
func ParseLevel(logLevel string) logrus.Level {
	switch logLevel {
	case "debug":
		return logrus.DebugLevel
	case "info":
		return logrus.InfoLevel
	case "warn":
		return logrus.WarnLevel
	case "error":
		return logrus.ErrorLevel
	default:
		return logrus.InfoLevel
	}
}

// Setup New CommonLogger
func DefaultBloopyFieldLogger() logrus.FieldLogger {
	myLogger := logrus.New()
	myLogger.Formatter = &logrus.TextFormatter{FullTimestamp: true, DisableColors: false}

	return myLogger.WithField("common", "group")
}
