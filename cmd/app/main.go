package main

import (
	"github.com/imdong/torrent-tools/internal/app"
	"github.com/imdong/torrent-tools/internal/config"
	"log"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
