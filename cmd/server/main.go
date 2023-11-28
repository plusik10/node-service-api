package main

import (
	"log"

	"github.com/plusik10/note-service-api/config"
	"github.com/plusik10/note-service-api/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("config Err: %s", err.Error())
	}
	app.Run(cfg)
}
