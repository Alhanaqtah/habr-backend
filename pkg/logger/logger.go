package logger

import (
	"log/slog"
	"os"
)

const (
	local = "local"
	dev   = "dev"
	prod  = "prod"
)

func New(env string) *slog.Logger {
	var log *slog.Logger

	if env == "" {
		env = local
	}

	switch env {
	case local:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
	case dev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
	case prod:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}))
	}

	return log
}

// Err Custom error handler
func Err(log *slog.Logger, op, msg string, err error) {
	log.Error(msg,
		slog.String("err", err.Error()),
		slog.Attr{
			Key:   "op",
			Value: slog.StringValue(op),
		})
}
