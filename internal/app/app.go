package app

import (
	"fmt"
	"log"
	"os"
	"urlShortener/internal/cfg"
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

	// storage = storage.New()
	// service = service.New()
	// transport = transport.New()
}
