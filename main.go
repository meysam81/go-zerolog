/*
Package gozerolog provides a simple wrapper around the zerolog logger with additional
configuration options for log level and color output.

This package allows you to create a new logger instance with customizable options
such as log level and colorization. It also sets the default time format for log
timestamps.

Quick Start:

Default logger (info level, no color):

	logger := gozerolog.NewLogger()
	logger.Info().Msg("application started")

Debug logger with color:

	logger := gozerolog.NewLogger(
		gozerolog.WithLogLevel("debug"),
		gozerolog.WithColor(true),
	)
	logger.Debug().Str("component", "database").Msg("connection established")

Production logger (warn level):

	logger := gozerolog.NewLogger(gozerolog.WithLogLevel("warn"))
	logger.Warn().Err(err).Msg("retrying failed operation")

Available log levels: "debug", "info", "warn", "error", "critical"
*/
package gozerolog

import (
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

// LoggerOption is a function that configures the logger.
type LoggerOption func(*loggerConfig)

// loggerConfig holds the configuration options for the logger.
type loggerConfig struct {
	logLevel     string // log level as a string
	colorEnabled bool   // flag to enable or disable color
}

// WithLogLevel sets the log level for the logger.
// Available levels are "debug", "info", "warn", "error", and "critical".
func WithLogLevel(level string) LoggerOption {
	return func(c *loggerConfig) {
		c.logLevel = level
	}
}

// WithColor enables or disables color output for the logger.
func WithColor(enabled bool) LoggerOption {
	return func(c *loggerConfig) {
		c.colorEnabled = enabled
	}
}

// NewLogger creates a new zerolog.Logger instance with the given options.
// It returns a pointer to the configured logger.
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
