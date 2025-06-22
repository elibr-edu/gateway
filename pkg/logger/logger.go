package logger

import (
	"fmt"
	"log/slog"
	"os"
)

func MustInitLogger(level string) *slog.Logger {
	const op = "logger.MustInitLogger"

	var log *slog.Logger

	out := os.Stdout

	switch level {
	case "debug":
		log = slog.New(slog.NewJSONHandler(out, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
	case "info":
		log = slog.New(slog.NewJSONHandler(out, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}))
	case "test":
		log = slog.New(slog.NewTextHandler(out, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
	default:
		panic(fmt.Sprintf("%s: level - {%s} not in {debug, test, info}", op, level))
	}
	return log
}
