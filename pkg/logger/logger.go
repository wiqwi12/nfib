package logger

import (
	"log/slog"
	"os"
	"sync"
)

type MyLogger interface {
	Info(msg string, any ...interface{})
	Error(msg string, any ...interface{})
}

type Lgr struct {
	logger *slog.Logger
}

var (
	once     sync.Once
	instance *Lgr
)

// NewSlogLogger инициализирует логгер только один раз
func NewSlogLogger() *Lgr {
	once.Do(func() {
		instance = &Lgr{
			logger: slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{})),
		}
	})
	return instance
}

func (l *Lgr) Info(msg string, any ...interface{}) {

	l.logger.Info(msg, any...)

}

func (l *Lgr) Error(msg string, any ...interface{}) {

	l.logger.Error(msg, any...)

}
