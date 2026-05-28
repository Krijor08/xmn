package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"example.com/go/backend/handlers"
	"example.com/go/backend/middleware"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	r := mux.NewRouter()

	homeHandler := http.HandlerFunc(handlers.Home)
	aboutHandler := http.HandlerFunc(handlers.About)

	r.HandleFunc("/", homeHandler.ServeHTTP).Methods("GET")
	r.HandleFunc("/about", aboutHandler.ServeHTTP).Methods("GET")

	r.PathPrefix("/static/").Handler(
		http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))),
	)

	port := os.Getenv("PORT")
	address := fmt.Sprintf("localhost:%s", port)

	fmt.Println(address)

	r.Use(middleware.Logger)
	fmt.Printf("Server is running on %s\n", address)
	log.Fatal(http.ListenAndServe(address, r))
}
