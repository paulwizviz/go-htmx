package logger

import (
	"log/slog"
	"os"
)

// InitLogger initializes slog with a configurable log level
func InitLogger() {
	logLevel := getLogLevelFromEnv()

	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: logLevel, // Set logging level dynamically
	})

	// Set the default logger
	slog.SetDefault(slog.New(handler))
}

// getLogLevelFromEnv reads the log level from an environment variable
func getLogLevelFromEnv() slog.Level {
	level := os.Getenv("LOG_LEVEL") // Read the log level from environment
	switch level {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	default:
		return slog.LevelInfo // Default to Info if not specified
	}
}
