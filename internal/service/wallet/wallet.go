package wallet

import (
	"wallet/internal/server/router"
	"wallet/internal/service/wallet/ports/repository"
	"wallet/internal/service/wallet/service"
	"wallet/internal/service/wallet/transport/api"
	"wallet/pkg/converter"

	"github.com/jmoiron/sqlx"
)

func NewService(r *router.Router, dbConn *sqlx.DB) *service.Service {
	repo := repository.New(dbConn)
	c := converter.New()
	svc := service.New(repo, c)
	api := api.New(svc)

	api.RegisterRoutes(r)

	return svc
}
