package payments

import (
	"encoding/json"
	"net/http"

	"github.com/hermesdt/form3-challenge/http/helpers"
	"github.com/hermesdt/form3-challenge/http/payloads"

	"github.com/hermesdt/form3-challenge/db"
	uuid "github.com/satori/go.uuid"
)

func Create(db db.DBHolder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		var payment payloads.Payment
		json.NewDecoder(r.Body).Decode(&payment)
		payment.ID = uuid.NewV4().String()

		_, err := db.GetDB().Collection("payments").InsertOne(nil, payment)
		if err != nil {
			helpers.WriteErrorJSON(w, err, 500)
			return
		}

		helpers.WriteJSON(w, &payloads.IDResponse{
			ID: payment.ID,
		}, 201)
	}
}

// func paymentExists(db db.DBHolder, payment payloads.Payment) (bool, error) {
// 	filter := map[string]interface{}{
// 		"id": payment.ID,
// 	}
// 	count, err := db.GetDB().Collection("payments").CountDocuments(nil, filter)
// 	if err != nil {
// 		return false, err
// 	}

// 	return count > 0, nil
// }
