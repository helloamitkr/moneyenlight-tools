package logger

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

// Init initializes the global logger
func Init() {
	// Set output to stdout
	log.Out = os.Stdout

	// Use JSON formatter for structured logging
	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		PrettyPrint:     false,
	})

	// Set log level (default: Info)
	level := strings.ToLower(os.Getenv("LOG_LEVEL"))
	switch level {
	case "debug":
		log.SetLevel(logrus.DebugLevel)
	case "warn":
		log.SetLevel(logrus.WarnLevel)
	case "error":
		log.SetLevel(logrus.ErrorLevel)
	default:
		log.SetLevel(logrus.InfoLevel)
	}

	log.Infof("Logger initialized with level: %s", log.GetLevel())
}

// Info logs info messages
func Info(args ...interface{}) {
	log.Info(args...)
}

// Warn logs warning messages
func Warn(args ...interface{}) {
	log.Warn(args...)
}

// Error logs error messages
func Error(args ...interface{}) {
	log.Error(args...)
}

// Debug logs debug messages
func Debug(args ...interface{}) {
	log.Debug(args...)
}

// WithFields logs with structured fields
func WithFields(fields logrus.Fields) *logrus.Entry {
	return log.WithFields(fields)
}
