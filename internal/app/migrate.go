package app

import (
	"fmt"
	"log"
	"wallet/internal/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// Run migrations
func RunMigrations(cfg config.Postgres) error {
	log.Printf("cfg.MigrationPath: %v\n", cfg.MigrationPath)
	fmt.Printf("cfg.ConnString(): %v\n", cfg.ConnString())
	dr, err := migrate.New(cfg.MigrationPath, cfg.ConnString())
	if err != nil {
		fmt.Println(err.Error())
		return fmt.Errorf("migrate Err: %w", err)
	}
	defer dr.Close()

	if err := dr.Up(); err != nil && err != migrate.ErrNoChange {
		log.Println(err.Error())
		fmt.Println(err)
		// err = dr.Drop()
		// if err != nil {
		// 	return err
		// }
		return err
	}
	//dr.Drop()
	log.Println("migrations done")
	return nil
}
