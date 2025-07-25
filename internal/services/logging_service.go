package services

import (
	"log/slog"
	"os"
)

type Logger interface {
	LogInfo(message string)
	LogError(message string)
	LogDebug(message string)
}

func NewLogger() Logger {
	// Create a new logger with a custom handler
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	return &Log{
		logger: logger,
	}
}

type Log struct {
	logger *slog.Logger
}

func (ls *Log) LogInfo(message string) {
	ls.logger.Info(message)
}

func (ls *Log) LogError(message string) {
	ls.logger.Error(message)
}

func (ls *Log) LogDebug(message string) {
	ls.logger.Debug(message)
}
