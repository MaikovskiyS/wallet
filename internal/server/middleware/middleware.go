package middleware

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
	"wallet/internal/apperrors"
)

var (
	ErrBadRequest   = apperrors.New(apperrors.ErrBadRequest, "Server-Middleware-")
	ErrUnauthorized = apperrors.New(apperrors.ErrUnauthorized, "Server-Middleware-")
)

type middleware struct {
}
type AppHandler func(w http.ResponseWriter, r *http.Request) error

func New() *middleware {
	return &middleware{}
}

// Log request letency and answer from server
func (m *middleware) Logging(h AppHandler) AppHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		t := time.Now()
		err := h(w, r)

		log.Println(time.Since(t))

		if err != nil {
			var er *apperrors.AppErr
			if errors.As(err, &er) {
				log.Println(er.Log())
				return err
			}
			log.Printf("unknown error: %s", err)
			return err
		}
		log.Println("success")
		return nil
	}
}

// Handle Error and sending answer to client
func (m *middleware) ErrorHandle(h AppHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := h(w, r)
		if err != nil {
			var er *apperrors.AppErr
			if errors.As(err, &er) {
				apperrors.ErrResponse(w, er)
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("unknown error:  %s", err.Error())))
			return

		}
	}
}
