package utils

import (
	"encoding/json"
	"net/http"
)

func WriteJson(data any, status int, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
