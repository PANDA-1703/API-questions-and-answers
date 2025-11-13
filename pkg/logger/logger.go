package logger

import (
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"
)

type Logger interface {
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
	With(args ...any) *slog.Logger
}

func New() Logger {
	w := os.Stdout
	// create new logger
	logger := slog.New(tint.NewHandler(
		w,
		&tint.Options{
			Level: slog.LevelDebug, // minimum log level
		}),
	)

	// set global logger
	slog.SetDefault(slog.New(
		tint.NewHandler(w, &tint.Options{
			TimeFormat: time.Kitchen,
		}),
	))

	return logger
}
