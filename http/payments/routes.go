package payments

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/hermesdt/form3-challenge/db"
)

func SetupRoutes(db db.DBHolder, router *chi.Mux) http.Handler {
	router.Route("/payments", func(r chi.Router) {
		r.Get("/", Index(db))
		r.Get("/{id}", Show(db))
	})

	return router
}
