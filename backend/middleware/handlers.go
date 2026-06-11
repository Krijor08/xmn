package middleware

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"example.com/go/backend/database"
	"example.com/go/backend/models"
	"github.com/rs/zerolog/log"
)

var tmpl = make(map[string]*template.Template)
var isAuthenticated bool

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
		BasePage: models.BasePage{CurrentPage: "home", IsAuthenticated: isAuthenticated},
		User:     models.HomeUser{ID: 1, Name: "Username", Role: "placeholder", Email: "placeholder@example.com"},
	}

	err := tmpl["home.html"].Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Printf("Error executing home: %v\n", err)
	}
}

func (h *Handler) About(w http.ResponseWriter, r *http.Request) {
	if !isAuthenticated {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

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

	User, err := database.AuthenticateUser(loginReq.Username, loginReq.Password, h.db)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Printf("Error occurred while authenticating user: %s\n", err.Error())
		return
	}

	if User.Name != "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		http.Error(w, fmt.Sprintf("User %s authenticated successfully!", loginReq.Username), http.StatusOK)
		isAuthenticated = true
	} else {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		http.Error(w, fmt.Sprintf("Authentication failed for user %s", loginReq.Username), http.StatusUnauthorized)
	}
}

func (h *Handler) SignupPage(w http.ResponseWriter, r *http.Request) {
	tmpl["signup.html"] = template.Must(template.ParseFiles("templates/base.html", "templates/signup.html"))

	data := models.SignupData{
		BasePage: models.BasePage{CurrentPage: "signup"},
		Roles: []models.Roles{
			{ID: "a", Role: "Guest"},
			{ID: "b", Role: "User"},
			{ID: "c", Role: "Tester"},
			{ID: "c", Role: "Something else"},
		},
	}

	err := tmpl["signup.html"].Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Printf("Error executing signup: %v\n", err)
	}
}

func (h *Handler) Signup(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	phone, err := strconv.Atoi(r.FormValue("phone"))
	if err != nil {
		phone = 0
	}

	signupReq := models.SignupRequest{
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
		Email:    r.FormValue("email"),
		Phone:    phone,
	}

	result, err := database.CreateUser(signupReq, h.db)
	if err != nil {
		log.Warn().Msg(fmt.Sprintf("Could not insert data: %s", err))
		http.Redirect(w, r, "/signup", http.StatusSeeOther)
		http.Error(w, fmt.Sprintf("Could not create user named %s", signupReq.Username), http.StatusUnauthorized)
		return
	}

	rows, _ := result.RowsAffected()

	if rows == 1 {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		http.Error(w, fmt.Sprintf("User %s created successfully!", signupReq.Username), http.StatusOK)
		isAuthenticated = true
	} else {
		log.Warn().Msg(fmt.Sprintf("Could not insert data: %s", err))
		http.Redirect(w, r, "/signup", http.StatusSeeOther)
		http.Error(w, fmt.Sprintf("Could not create user named %s", signupReq.Username), http.StatusUnauthorized)
	}
}

func (h *Handler) NoSQL(w http.ResponseWriter, r *http.Request) {
	tmpl["login.html"] = template.Must(template.ParseFiles("templates/base.html", "templates/nosql.html"))

	err := tmpl["nosql.html"].Execute(w, "base.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Printf("Error executing NoSQL: %v\n", err)
	}
}
