package app

import (
	"context"
	"wallet/internal/config"
	"wallet/internal/server"
	"wallet/internal/server/router"
	"wallet/internal/service/wallet"

	"github.com/jmoiron/sqlx"
)

// init app deps and running http server
func Run(cfg *config.Config) error {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//init database client
	db, err := sqlx.ConnectContext(ctx, cfg.Psql.Driver(), cfg.Psql.ConnString())
	if err != nil {
		return err
	}

	//init http(default mux) router
	r := router.New()

	// init wallet service
	wallet.NewService(r, db)
	s := server.New(cfg)
	s.SetHandler(r)
	err = s.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}
