package helpers

import (
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/hermesdt/form3-challenge/http/payloads"
)

func WriteJSON(w http.ResponseWriter, data interface{}, status int) {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

var documentNotFound = payloads.ErrorResponse{Error: "document not found"}

func WriteNotFoundErrorJSON(w http.ResponseWriter) {
	WriteJSON(w, documentNotFound, 404)
}

func WriteErrorJSON(w http.ResponseWriter, err error, status int) {
	if err == mongo.ErrNoDocuments {
		WriteNotFoundErrorJSON(w)
		return
	}

	WriteJSON(w, payloads.ErrorResponse{
		Error: err.Error(),
	}, status)
}
