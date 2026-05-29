package handlers

import (
	"html/template"
	"net/http"
)

var aboutTemplate = template.Must(template.ParseFiles("templates/about.html"))

func About(w http.ResponseWriter, r *http.Request) {
	aboutTemplate.Execute(w, nil)
}
