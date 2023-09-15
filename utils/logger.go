package utils

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

func loggerConfig() {
	logLevel := os.Getenv("LOG_LEVEL")
	logLevel = strings.ToLower(logLevel)

	switch logLevel {
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	default:
		logrus.SetLevel(logrus.InfoLevel)
	}
}
