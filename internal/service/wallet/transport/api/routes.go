package api

import "wallet/internal/server/router"

func (a *api) RegisterRoutes(r *router.Router) {

	r.HandleFunc("/wallet/create", r.ErrorHandle(r.Logging(a.CreateWallet)))
	r.HandleFunc("/wallet/update", r.ErrorHandle(r.Logging(a.UpdateBalance)))
	r.HandleFunc("/wallet/transfer", r.ErrorHandle(r.Logging(a.TransferAmount)))
}
