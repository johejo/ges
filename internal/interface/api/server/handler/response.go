package handler

import (
	"encoding/json"
	"net/http"
)

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		respondJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
}

func respondError(w http.ResponseWriter, status int, err error) {
	respondJSON(w, status, map[string]interface{}{"error": err.Error()})
}
