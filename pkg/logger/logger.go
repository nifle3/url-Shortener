package logger

import (
	"log"
	"log/slog"
	"os"
)

func MustCreate(build string) *slog.Logger {
	var logger *slog.Logger

	switch build {
	case "dev":
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level:     slog.LevelDebug,
			AddSource: true,
		}))
	case "prod":
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level:     slog.LevelInfo,
			AddSource: false,
		}))
	default:
		log.Fatal("build type is not define")
	}

	return logger
}
