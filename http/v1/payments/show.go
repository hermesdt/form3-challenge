package payments

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/hermesdt/form3-challenge/db"
	"github.com/hermesdt/form3-challenge/http/helpers"
	"github.com/hermesdt/form3-challenge/http/payloads"
)

func Show(db db.DBHolder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payment, err := getPayment(db, chi.URLParam(r, "id"))
		if err != nil {
			helpers.WriteErrorJSON(w, err, http.StatusInternalServerError)
			return
		}

		data := map[string]interface{}{
			"payment": payment,
		}
		helpers.WriteJSON(w, data, http.StatusOK)
	}
}

func getPayment(db db.DBHolder, id string) (*payloads.Payment, error) {
	filter := map[string]interface{}{
		"id": id,
	}
	singleResult := db.GetDB().Collection("payments").FindOne(nil, filter)
	if singleResult.Err() != nil {
		return nil, singleResult.Err()
	}

	var payment payloads.Payment
	err := singleResult.Decode(&payment)
	if err != nil {
		return nil, err
	}

	return &payment, nil
}
