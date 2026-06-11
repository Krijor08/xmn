package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"example.com/go/backend/database"
	"example.com/go/backend/middleware"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// Load .env when not in docker container
func init() {
	if os.Getenv("DOCKER_ENV") != "true" {
		_ = godotenv.Load()
	}
}

var sql bool

func main() {
	db, err := database.InitMySQL()
	if err == nil {
		sql = true
	} else {
		sql = false
	}

	defer db.Close()

	r := mux.NewRouter()
	handle := middleware.NewHandler(db)

	r.HandleFunc("/about", handle.About).Methods("GET")

	r.HandleFunc("/", handle.Home).Methods("GET")

	r.PathPrefix("/static/").Handler(
		http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))),
	)

	if sql {
		r.HandleFunc("/login", handle.LoginPage).Methods("GET")
		r.HandleFunc("/signup", handle.SignupPage).Methods("GET")

		r.HandleFunc("/api/login", handle.Login).Methods("POST")
		r.HandleFunc("/api/signup", handle.Signup).Methods("POST")
	} else {
		r.HandleFunc("/login", handle.NoSQL).Methods("GET")
		r.HandleFunc("/signup", handle.NoSQL).Methods("GET")
	}

	address := os.Getenv("ADDRESS")

	r.Use(middleware.Logger)
	fmt.Printf("Server is running on %s\n", address)
	log.Fatal(http.ListenAndServe(address, r))
}
