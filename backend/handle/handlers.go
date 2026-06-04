package handle

import (
	"fmt"
	"html/template"
	"net/http"

	. "example.com/go/backend/middleware"
)

var tmpl = make(map[string]*template.Template)

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl["home.html"] = template.Must(template.ParseFiles("templates/base.html", "templates/home.html"))

	data := HomeData{
		BasePage: BasePage{CurrentPage: "home"},
		User:     User{ID: 1, Name: "Username", Role: "placeholder", Email: "placeholder@example.com"},
	}

	err := tmpl["home.html"].Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Printf("Error executing home: %v\n", err)
	}
}

func About(w http.ResponseWriter, r *http.Request) {
	tmpl["about.html"] = template.Must(template.ParseFiles("templates/base.html", "templates/about.html"))

	data := AboutData{BasePage: BasePage{CurrentPage: "about"}, Version: "0.0.1"}

	fmt.Printf("About page accessed {%s}\n", data)

	err := tmpl["about.html"].Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Printf("Error executing about: %v\n", err)
	}
}
