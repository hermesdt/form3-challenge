package payments

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/hermesdt/form3-challenge/db"
)

func SetupRoutes(db db.DBHolder, router chi.Router) http.Handler {
	router.Get("/", Index(db))
	router.Get("/{id}", Show(db))
	router.Post("/", Create(db))
	router.Delete("/{id}", Delete(db))

	return router
}
