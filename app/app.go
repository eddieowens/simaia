package app

import (
	"github.com/eddieowens/simaia/app/config"
	log "github.com/sirupsen/logrus"
	"strings"
)

const Key = "App"

type App interface {
	Start() error
}

type app struct {
	Config *config.Config `inject:"Config"`
}

func (a *app) Start() error {
	format := &log.JSONFormatter{
		TimestampFormat: a.Config.Log.TimeFormat,
	}

	log.SetFormatter(format)
	log.SetLevel(resolveLevel(a.Config.Log.Level))
}

func resolveLevel(level string) log.Level {
	switch strings.ToLower(level) {
	case "trace":
		return log.TraceLevel
	case "debug":
		return log.DebugLevel
	case "info":
		return log.InfoLevel
	case "warn", "warning":
		return log.WarnLevel
	case "error":
		return log.ErrorLevel
	case "fatal":
		return log.FatalLevel
	case "panic":
		return log.PanicLevel
	default:
		log.Info(level, " is not a valid log level. Setting to info.")
		return log.InfoLevel
	}
}
