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

func main() {
	godotenv.Load() // Support for .env file
	db := database.InitMySQL()

	r := mux.NewRouter()
	handle := middleware.NewHandler(db)

	r.HandleFunc("/about", handle.About).Methods("GET")
	r.HandleFunc("/login", handle.LoginPage)

	r.HandleFunc("/", handle.Home).Methods("GET")

	r.PathPrefix("/static/").Handler(
		http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))),
	)

	r.HandleFunc("/api/login", handle.Login).Methods("POST")

	address := os.Getenv("ADDRESS")

	r.Use(middleware.Logger)
	fmt.Printf("Server is running on %s\n", address)
	log.Fatal(http.ListenAndServe(address, r))
}
