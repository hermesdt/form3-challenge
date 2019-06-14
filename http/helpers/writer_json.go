package helpers

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, data map[string]interface{}, status int) {
	w.WriteHeader(200)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
