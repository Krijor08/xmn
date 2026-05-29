package handlers

import (
	// "encoding/json"
	"html/template"
	"net/http"
	// . "example.com/go/backend/middleware"
)

var templates = template.Must(
	template.ParseGlob("templates/*.html"),
)

func Home(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "base.html", nil)

}

// func respondJSON(w http.ResponseWriter, status int, payload ApiResponse) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(status)
// 	json.NewEncoder(w).Encode(payload)
// }
