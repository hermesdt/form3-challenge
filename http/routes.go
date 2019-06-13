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

	payments.SetupRoutes(db, router)

	return router
}
