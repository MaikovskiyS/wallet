package main

import (
	"log"
	"wallet/internal/app"
	"wallet/internal/config"
)

// init config and migrations, running app
func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}
	err = app.RunMigrations(cfg.Psql)
	if err != nil {
		log.Fatal(err)
	}
	err = app.Run(cfg)
	if err != nil {
		log.Fatal(err)
	}
}
