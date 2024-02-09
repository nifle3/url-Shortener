package app

import (
	"fmt"
	"log"
	"os"
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
	cfgPath := os.Getenv(cfgPath)
	if cfgPath == "" {
		log.Fatal("cfgPath is not set")
	}
	cfg := cfg.MustLoad(cfgPath)

	log := logger.MustCreate(cfg.Build)
	log.Info(fmt.Sprintf("app is starting with cfg: %#v", cfg))

	storage := inmemory.New()
	urlShortener := service.New(storage)

	transport := http.New(&urlShortener)
	transport.Listen(cfg.Port)
}
