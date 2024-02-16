package app

import (
	"fmt"
	"os"
	"urlShortener/internal/cfg"
	"urlShortener/internal/service"
	"urlShortener/internal/storage"
	"urlShortener/internal/transport/http"
	"urlShortener/pkg/logger"
)

func Run() {
	config := cfg.GetInstance()

	log := logger.MustCreate(config.Build)
	log.Info(fmt.Sprintf("app is starting with cfg: %#v", config))

	storageBuilder := storage.NewRedisBuilder()

	storageBuilder.SetPassword(config.Cache.Password)
	redis := storageBuilder.Build()

	urlShortener := service.New(redis)

	transport := http.New(&urlShortener)
	if err := transport.Listen(config.Port); err != nil {
		os.Exit(1)
	}
}
