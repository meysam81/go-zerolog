package gozerolog

import (
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

type LoggerOption func(*loggerConfig)

type loggerConfig struct {
	logLevel     string
	colorEnabled bool
}

func WithLogLevel(level string) LoggerOption {
	return func(c *loggerConfig) {
		c.logLevel = level
	}
}

func WithColor(enabled bool) LoggerOption {
	return func(c *loggerConfig) {
		c.colorEnabled = enabled
	}
}

func NewLogger(opts ...LoggerOption) *zerolog.Logger {
	config := &loggerConfig{
		logLevel:     "info",
		colorEnabled: false,
	}

	for _, opt := range opts {
		opt(config)
	}

	zerolog.TimeFieldFormat = time.RFC3339

	level := zerolog.InfoLevel
	switch strings.ToLower(config.logLevel) {
	case "debug":
		level = zerolog.DebugLevel
	case "info":
		level = zerolog.InfoLevel
	case "warn":
		level = zerolog.WarnLevel
	case "error":
		level = zerolog.ErrorLevel
	case "critical":
		level = zerolog.FatalLevel
	}

	l := zerolog.New(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC3339,
		NoColor:    !config.colorEnabled,
	}).Level(level).With().Caller().Timestamp().Logger()

	return &l
}
