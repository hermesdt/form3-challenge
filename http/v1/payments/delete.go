package payments

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/hermesdt/form3-challenge/db"
	"github.com/hermesdt/form3-challenge/http/helpers"
)

func Delete(db db.DBHolder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		filter := map[string]interface{}{"id": id}
		result, err := db.GetDB().Collection("payments").DeleteOne(nil, filter)
		if err != nil {
			helpers.WriteErrorJSON(w, err, http.StatusInternalServerError)
			return
		}

		if result.DeletedCount != 1 {
			helpers.WriteNotFoundErrorJSON(w)
			return
		}

		data := map[string]interface{}{}
		helpers.WriteJSON(w, data, http.StatusOK)
	}
}
