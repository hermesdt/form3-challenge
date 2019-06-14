package http

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/hermesdt/form3-challenge/db"
	"github.com/hermesdt/form3-challenge/http/payments"
)

func SetupRoutes(db db.DBHolder) http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Route("/payments", func(r chi.Router) {
		payments.SetupRoutes(db, r)
	})

	return router
}
