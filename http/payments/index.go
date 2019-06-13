package payments

import (
	"encoding/json"
	"net/http"

	"github.com/hermesdt/form3-challenge/db"
	"go.mongodb.org/mongo-driver/bson"
)

func Index(db db.DBHolder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payments, err := getAllPayments(db)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"payments": payments,
		})
	}
}

func getAllPayments(db db.DBHolder) ([]interface{}, error) {
	filter := make(map[string]interface{})
	cursor, err := db.GetDB().Collection("payments").Find(nil, filter)
	if err != nil {
		return nil, err
	}

	var payments = []interface{}{}
	for cursor.Next(nil) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			return nil, err
		}

		payments = append(payments, result)
	}

	return payments, nil
}
