package logger

import "github.com/sirupsen/logrus"

type Logger interface {
	Info(args ...any)
	Error(args ...any)
}

type logrusLogger struct {
	*logrus.Logger
}

func New() Logger {
	l := logrus.New()
	l.SetFormatter(&logrus.JSONFormatter{})
	return &logrusLogger{l}
}
