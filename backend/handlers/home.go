package handlers

import (
	"encoding/json"
	"net/http"

	. "example.com/go/backend/middleware"
)

func Home(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, ApiResponse{
		Message: "Welcome to the home page",
		Status:  "success",
	})
}

func respondJSON(w http.ResponseWriter, status int, payload ApiResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}
