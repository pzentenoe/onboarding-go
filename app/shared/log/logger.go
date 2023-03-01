package log

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func GetLogger() *log.Entry {
	logger := log.WithFields(log.Fields{
		"application_name": "cars-api",
	})
	return logger
}

func InitLogger() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
}
