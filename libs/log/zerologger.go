package log

import (
	"io"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// ZeroLogger is a wrapper for the Zerolog library to implement log.Logger interface
type ZeroLogger struct {
	rootLogger   zerolog.Logger // Root logger instance
	callerLogger zerolog.Logger // Caller logger attaches the caller
}

// NewZeroLogger instantiates a new Zerologger with configs
func NewZeroLogger(logLevel string, output io.Writer) *ZeroLogger {
	newLogger := new(ZeroLogger)

	writer := zerolog.ConsoleWriter{Out: output}
	// Instantiate loggers
	newLogger.rootLogger = zerolog.New(writer).With().Timestamp().Logger() // Root Logger
	// NOTE: 3 stack frames need to be skipped since the error function(s) are wrapped
	newLogger.callerLogger = newLogger.rootLogger.With().CallerWithSkipFrameCount(3).Logger() // callerLogger
	newLogger.SetLogLevel(logLevel)

	return newLogger
}

// SetLogLevel the log level to display
func (zl *ZeroLogger) SetLogLevel(level string) {
	switch level {
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	default:
		log.Warn().Msgf("Unrecognized log level: %s, setting log level to INFO", level)
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
	zl.Debugf("Setting log level to \"%s\"", level)
}

// Debugf prints a formatted DEBUG level message
func (zl *ZeroLogger) Debugf(msg string, args ...interface{}) {
	zl.rootLogger.Debug().Msgf(msg, args...)
}

// Infof prints a formatted INFO level message
func (zl *ZeroLogger) Infof(msg string, args ...interface{}) {
	zl.rootLogger.Info().Msgf(msg, args...)
}

// Warnf prints a formatted WARN level message
func (zl *ZeroLogger) Warnf(msg string, args ...interface{}) {
	zl.callerLogger.Warn().Msgf(msg, args...)
}

// Errorf prints a formatted ERROR level message
func (zl *ZeroLogger) Errorf(msg string, args ...interface{}) {
	zl.callerLogger.Error().Msgf(msg, args...)
}

// Fatalf prints a formatted FATAL level message
func (zl *ZeroLogger) Fatalf(msg string, args ...interface{}) {
	zl.callerLogger.Fatal().Msgf(msg, args...)
}
