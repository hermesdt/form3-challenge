package payments

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/hermesdt/form3-challenge/db"
	"github.com/hermesdt/form3-challenge/http/helpers"
	"go.mongodb.org/mongo-driver/bson"
)

func Show(db db.DBHolder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payment, err := getPayment(db, chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		data := map[string]interface{}{
			"payment": payment,
		}
		helpers.WriteJSON(w, data, 200)
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
