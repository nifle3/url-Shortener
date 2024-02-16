package app

import (
	"fmt"
	"urlShortener/internal/cfg"
	"urlShortener/internal/service"
	inmemory "urlShortener/internal/storage/inMemory"
	"urlShortener/internal/transport/http"
	"urlShortener/pkg/logger"
)

const (
	cfgPath = "CFG_PATH"
)

func Run() {
	cfg := cfg.GetInstance()

	log := logger.MustCreate(cfg.Build)
	log.Info(fmt.Sprintf("app is starting with cfg: %#v", cfg))

	storage := inmemory.New()
	urlShortener := service.New(storage)

	transport := http.New(&urlShortener)
	transport.Listen(cfg.Port)
}
