package payments

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/hermesdt/form3-challenge/db"
	paymentsV1 "github.com/hermesdt/form3-challenge/http/v1/payments"
)

func SetupRoutes(db db.DBHolder, router chi.Router) http.Handler {
	router.Get("/", paymentsV1.Index(db))
	router.Get("/{id}", paymentsV1.Show(db))

	router.Post("/", paymentsV1.Create(db))
	router.Delete("/{id}", paymentsV1.Delete(db))
	router.Put("/{id}", paymentsV1.Update(db))

	return router
}
