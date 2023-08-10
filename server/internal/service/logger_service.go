package service

import (
	"fmt"

	"github.com/phamphihungbk/go-graphql-api/internal/config"
	"github.com/phamphihungbk/go-graphql-api/internal/util"
	"github.com/sirupsen/logrus"
)

type Logger struct {
	instance *logrus.Logger
	config   config.Server
}

type ILogger interface {
	Info(message string)
	Error(message string)
	Fatal(message string)
}

func NewLogger(config config.Server) *Logger {
	logger := logrus.New()
	logLocation := util.OpenFile("logs", "app.log")
	logger.SetOutput(logLocation)

	if config.Env == "production" {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}

	return &Logger{instance: logger, config: config}
}

func (l *Logger) Info(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.instance.Info(msg)
}
func (l *Logger) Error(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.instance.Error(msg)
}

func (l *Logger) Fatal(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.instance.Fatal(msg)
}
