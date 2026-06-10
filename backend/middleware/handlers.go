package middleware

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"example.com/go/backend/database"
	"example.com/go/backend/models"
)

var tmpl = make(map[string]*template.Template)

var isAuthenticated bool = false

type Handler struct {
	db *sql.DB
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	if !isAuthenticated {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

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
	defer http.Redirect(w, r, "/login", http.StatusSeeOther)
	r.ParseForm()

	loginReq := models.LoginRequest{
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	}

	fmt.Printf("Received login attempt for username: %s\n", loginReq.Username)

	var err error
	isAuthenticated, err = database.AuthenticateUser(loginReq.Username, loginReq.Password, h.db)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Printf("Error occurred while authenticating user: %s\n", err.Error())
		return
	}

	if isAuthenticated {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		http.Error(w, fmt.Sprintf("User %s authenticated successfully!", loginReq.Username), http.StatusOK)
	} else {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		http.Error(w, fmt.Sprintf("Authentication failed for user %s", loginReq.Username), http.StatusUnauthorized)
	}
}

func (h *Handler) SignupPage(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	phone, err := strconv.Atoi(r.FormValue("phone"))
	if err != nil {
		http.Error(w, fmt.Sprintf("%v is not a valid phone number", r.FormValue("phone")), http.StatusBadRequest)
	}

	role_id, err := database.GetRoleId(r.FormValue("role"), h.db)

	signupReq := models.SignupRequest{
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
		Role_ID:  role_id,
		Email:    r.FormValue("email"),
		Phone:    phone,
	}

	database.CreateUser(signupReq, h.db)
}
