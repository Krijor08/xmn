package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"example.com/go/backend/handle"
	. "example.com/go/backend/middleware"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	r := mux.NewRouter()

	r.HandleFunc("/about", handle.About).Methods("GET")
	r.HandleFunc("/", handle.Home).Methods("GET")

	r.PathPrefix("/static/").Handler(
		http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))),
	)

	address := os.Getenv("ADDRESS")

	r.Use(Logger)
	fmt.Printf("Server is running on %s\n", address)
	log.Fatal(http.ListenAndServe(address, r))
}
