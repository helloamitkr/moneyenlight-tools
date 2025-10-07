package logger

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func Init() {
	Log.Out = os.Stdout
	Log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	level := strings.ToLower(os.Getenv("LOG_LEVEL"))
	switch level {
	case "debug":
		Log.SetLevel(logrus.DebugLevel)
	case "warn":
		Log.SetLevel(logrus.WarnLevel)
	case "error":
		Log.SetLevel(logrus.ErrorLevel)
	default:
		Log.SetLevel(logrus.InfoLevel)
	}

	Log.Infof("Logger initialized with level: %s", Log.GetLevel())
}
