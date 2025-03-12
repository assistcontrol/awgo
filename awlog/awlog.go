package awlog

import (
	"log/slog"
	"os"
)

type Level int

type Config struct {
	Level Level
}

const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	Verbose
)

func Setup(c Config) slog.Level {
	var level slog.Level

	switch c.Level {
	case Verbose, DebugLevel:
		level = slog.LevelDebug
	case WarnLevel:
		level = slog.LevelWarn
	case ErrorLevel:
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	slog.SetLogLoggerLevel(level)
	return level
}

// Debug, Info, Warn, and Error are direct wrappers for the slog equivalents.
func Debug(msg string, args ...any) {
	slog.Debug(msg, args...)
}

func Info(msg string, args ...any) {
	slog.Info(msg, args...)
}

func Warn(msg string, args ...any) {
	slog.Warn(msg, args...)
}

func Error(msg string, args ...any) {
	slog.Error(msg, args...)
}

// Fatal is a wrapper for slog.Error that also calls os.Exit(1).
func Fatal(msg string, args ...any) {
	slog.Error(msg, args...)
	os.Exit(1)
}

func FatalError(msg string, err error) {
	Fatal(msg, "error", err)
}
