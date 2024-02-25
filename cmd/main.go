package main

import (
	"log"
	"os"
	"urlShortener/internal/app"

	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("CFG_PATH") == "" {
		if err := godotenv.Load(); err != nil {
			log.Fatal(err.Error())
		}
	}

	app.Run()
}
