package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"wallet/internal/config"
	"wallet/internal/server"
	"wallet/internal/server/router"
	"wallet/internal/service/wallet"

	"github.com/jmoiron/sqlx"
)

// init app deps and running http server
func Run(cfg *config.Config) error {

	//init database client
	db, err := sqlx.Connect(cfg.Psql.Driver(), cfg.Psql.ConnString())
	if err != nil {
		return fmt.Errorf("postgres connection Err: %w", err)
	}
	//init http(default mux) router
	r := router.New()

	// init wallet service
	wallet.NewService(r, db)

	//init http server
	s := server.New(cfg)
	s.SetHandler(r)

	// starting http server
	go func() {
		log.Printf("Starting server on %s", cfg.HttpServer.Port)
		err = s.ListenAndServe()
		if err != nil {
			if err == http.ErrServerClosed {
				log.Println("HTTP server stopped")
			}
			return
		}

	}()

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)

	<-sigint
	err = db.Close()
	if err != nil {
		return err
	}
	err = s.Shutdown()
	if err != nil {
		return err
	}
	log.Println("app closed")
	return nil
}
