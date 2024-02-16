package logger

import (
	"log/slog"
	"os"
	"sync"
)

var onceSetBuild sync.Once
var build string

var instance *slog.Logger
var once sync.Once

func SetBuild(buildType string) {
	onceSetBuild.Do(func() {
		build = buildType
	})
}

func GetInstance() *slog.Logger {
	if build == "" {
		onceSetBuild.Do(func() {
			build = "dev"
		})
	}

	once.Do(create)

	return instance
}

func create() {
	switch build {
	case "dev":
		instance = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level:     slog.LevelDebug,
			AddSource: true,
		}))
	case "prod":
		instance = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level:     slog.LevelInfo,
			AddSource: false,
		}))
	}
}
