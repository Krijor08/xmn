package handlers

import (
	"encoding/json"
	"html/template"
	"net/http"

	. "example.com/go/backend/middleware"
)

var homeTemplate = template.Must(template.ParseFiles("templates/index.html"))

func Home(w http.ResponseWriter, r *http.Request) {
	homeTemplate.Execute(w, nil)
}

func respondJSON(w http.ResponseWriter, status int, payload ApiResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}
