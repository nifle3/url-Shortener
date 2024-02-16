package app

import (
	"fmt"
	"os"
	"urlShortener/internal/cfg"
	"urlShortener/internal/service"
	"urlShortener/internal/storage/inMemory"
	"urlShortener/internal/transport/http"
	"urlShortener/pkg/logger"
)

func Run() {
	config := cfg.GetInstance()

	log := logger.MustCreate(config.Build)
	log.Info(fmt.Sprintf("app is starting with cfg: %#v", config))

	storage := inmemory.New()
	urlShortener := service.New(storage)

	transport := http.New(&urlShortener)
	if err := transport.Listen(config.Port); err != nil {
		os.Exit(1)
	}
}
