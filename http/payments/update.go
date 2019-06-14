package payments

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/hermesdt/form3-challenge/db"
	"github.com/hermesdt/form3-challenge/http/helpers"
	"github.com/hermesdt/form3-challenge/http/payloads"
)

func Update(db db.DBHolder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		filter := map[string]interface{}{"id": id}

		var payment payloads.Payment
		json.NewDecoder(r.Body).Decode(&payment)
		update := bson.M{
			"$set": payment,
		}
		upsert := false
		opts := options.UpdateOptions{Upsert: &upsert}
		result, err := db.GetDB().Collection("payments").UpdateOne(nil, filter, update, &opts)
		if err != nil {
			helpers.WriteErrorJSON(w, err, http.StatusInternalServerError)
			return
		}

		if result.ModifiedCount != 1 {
			helpers.WriteNotFoundErrorJSON(w)
			return
		}

		data := map[string]interface{}{}
		helpers.WriteJSON(w, data, http.StatusOK)
	}
}
