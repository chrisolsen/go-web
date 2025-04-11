package services

import (
	"log/slog"
	"os"
)

type LoggingServicer interface {
	LogInfo(message string)
	LogError(message string)
	LogDebug(message string)
}

func NewLoggingService() LoggingServicer {
	// Create a new logger with a custom handler
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	return &LoggingService{
		logger: logger,
	}
}

type LoggingService struct {
	logger *slog.Logger
}

func (ls *LoggingService) LogInfo(message string) {
	ls.logger.Info(message)
}

func (ls *LoggingService) LogError(message string) {
	ls.logger.Error(message)
}

func (ls *LoggingService) LogDebug(message string) {
	ls.logger.Debug(message)
}
