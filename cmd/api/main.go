package main

import (
	"log"

	"github.com/Alwin18/golang-modular-template/config"
	"github.com/Alwin18/golang-modular-template/internal/app"
)

func main() {
	cfg := config.LoadConfig()

	container, err := app.NewContainer(cfg)
	if err != nil {
		log.Fatal(err)
	}

	app := app.NewApp(container)
	log.Fatal(app.Listen(":" + cfg.AppPort))
}
