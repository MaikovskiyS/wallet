package app

import (
	"fmt"
	"log"
	"time"
	"wallet/internal/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// Run migrations
func RunMigrations(cfg config.Postgres) error {
	dr, err := migrate.New(cfg.MigrationPath, cfg.ConnString())
	if err != nil {
		time.Sleep(1 * time.Second)
		fmt.Println(err.Error())
		return fmt.Errorf("migrate Err: %w", err)
	}
	defer dr.Close()

	if err := dr.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	log.Println("migrations done")
	return nil
}
