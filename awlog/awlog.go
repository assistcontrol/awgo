package awlog

import (
	"log/slog"
	"os"
)

type Config struct {
	Level Level
}

type Level int

const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	Verbose
)

func Setup(c Config) slog.Level {
	level := logLevel(c.Level)
	slog.SetLogLoggerLevel(level)
	return level
}

func logLevel(level Level) slog.Level {
	switch level {
	case Verbose, DebugLevel:
		return slog.LevelDebug
	case WarnLevel:
		return slog.LevelWarn
	case ErrorLevel:
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
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
