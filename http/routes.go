package http

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/hermesdt/form3-challenge/db"
	paymentsV1 "github.com/hermesdt/form3-challenge/http/v1/payments"
	paymentsV2 "github.com/hermesdt/form3-challenge/http/v2/payments"
)

func SetupRoutes(db db.DBHolder) http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Route("/v1/payments", func(r chi.Router) {
		paymentsV1.SetupRoutes(db, r)
	})

	router.Route("/v2/payments", func(r chi.Router) {
		paymentsV2.SetupRoutes(db, r)
	})

	return router
}
