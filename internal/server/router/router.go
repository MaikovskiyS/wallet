package router

import (
	"net/http"
	"wallet/internal/server/middleware"
)

type Middle interface {
	Logging(h middleware.AppHandler) middleware.AppHandler
	ErrorHandle(h middleware.AppHandler) http.HandlerFunc
}
type Router struct {
	Middle
	*http.ServeMux
}

func New() *Router {
	mux := http.NewServeMux()
	m := middleware.New()

	return &Router{m, mux}
}
