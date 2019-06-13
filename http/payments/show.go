package payments

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/hermesdt/form3-challenge/db"
	"go.mongodb.org/mongo-driver/bson"
)

func Show(db db.DBHolder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payment, err := getPayment(db, chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"payment": payment,
		})
	}
}

func getPayment(db db.DBHolder, id string) (interface{}, error) {
	filter := map[string]interface{}{
		"id": id,
	}
	singleResult := db.GetDB().Collection("payments").FindOne(nil, filter)

	var result bson.M
	err := singleResult.Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
