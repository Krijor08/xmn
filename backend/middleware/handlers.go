package middleware

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	"example.com/go/backend/database"
	"example.com/go/backend/models"
)

var tmpl = make(map[string]*template.Template)

type Handler struct {
	db *sql.DB
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	tmpl["home.html"] = template.Must(template.ParseFiles("templates/base.html", "templates/home.html"))

	data := models.HomeData{
		BasePage: models.BasePage{CurrentPage: "home"},
		User:     models.User{ID: 1, Name: "Username", Role: "placeholder", Email: "placeholder@example.com"},
	}

	err := tmpl["home.html"].Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Printf("Error executing home: %v\n", err)
	}
}

func (h *Handler) About(w http.ResponseWriter, r *http.Request) {
	tmpl["about.html"] = template.Must(template.ParseFiles("templates/base.html", "templates/about.html"))

	data := models.AboutData{BasePage: models.BasePage{CurrentPage: "about"}, Version: "0.0.1"}

	err := tmpl["about.html"].Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Printf("Error executing about: %v\n", err)
	}
}

func (h *Handler) LoginPage(w http.ResponseWriter, r *http.Request) {
	tmpl["login.html"] = template.Must(template.ParseFiles("templates/base.html", "templates/login.html"))

	data := models.LoginData{BasePage: models.BasePage{CurrentPage: "login"}}

	err := tmpl["login.html"].Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Printf("Error executing login: %v\n", err)
	}

}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	loginReq := models.LoginRequest{
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	}

	fmt.Printf("Received login attempt for username: %s\n", loginReq.Username)

	isAuthenticated, err := database.AuthenticateUser(loginReq.Username, loginReq.Password, h.db)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Printf("Error occurred while authenticating user: %s\n", err.Error())
		return
	}

	if isAuthenticated {
		fmt.Printf("User %s authenticated successfully!\n", loginReq.Username)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		fmt.Printf("Authentication failed for user %s\n", loginReq.Username)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

}
