package handle

import (
	"html/template"
	"net/http"

	. "example.com/go/backend/middleware"
)

var templates = template.Must(
	template.ParseGlob("templates/*.html"),
)

func Home(w http.ResponseWriter, r *http.Request) {
	data := HomeData{
		BasePage: BasePage{CurrentPage: "home"},
		Users:    []User{{ID: 1, Name: "placeholder", Role: "admin", Email: "placeholder@example.com"}},
	}

	templates.ExecuteTemplate(w, "home.html", data)

}

func About(w http.ResponseWriter, r *http.Request) {
	data := BasePage{CurrentPage: "about"}

	templates.ExecuteTemplate(w, "about.html", data)
}
