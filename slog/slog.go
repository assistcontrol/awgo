// slog with a few extra methods
//
//   - slog.Fatal("msg", "name", "Foo Bar")
//
//   - slog.FatalError("msg", err)
//
//   - slog.Verbose() // debug
//
//   - slog.Quiet()   // warn
package slog

import (
	"log/slog"
	"os"
)

// Just use slog directly for these four
var (
	Debug = slog.Debug
	Info  = slog.Info
	Warn  = slog.Warn
	Error = slog.Error
)

func SetLevel(level slog.Level) {
	slog.SetLogLoggerLevel(level)
}

func Verbose() {
	SetLevel(slog.LevelDebug)
}

func Quiet() {
	SetLevel(slog.LevelWarn)
}

// Fatal wraps slog.Error and then calls os.Exit(1).
func Fatal(msg string, args ...any) {
	Error(msg, args...)
	os.Exit(1)
}

func FatalError(msg string, err error) {
	Fatal(msg, "error", err)
}
