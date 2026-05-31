package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"example.com/go/backend/middleware"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var templates = template.Must(
	template.ParseGlob("templates/*.html"),
)

func Home(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "base.html", "home.html")

}

func About(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "base.html", "about.html")
}

func main() {
	godotenv.Load()

	r := mux.NewRouter()

	homeHandler := http.HandlerFunc(Home)
	aboutHandler := http.HandlerFunc(About)

	r.HandleFunc("/", homeHandler.ServeHTTP).Methods("GET")
	r.HandleFunc("/about", aboutHandler.ServeHTTP).Methods("GET")

	r.PathPrefix("/static/").Handler(
		http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))),
	)

	address := os.Getenv("ADDRESS")

	fmt.Println(address)

	r.Use(middleware.Logger)
	fmt.Printf("Server is running on %s\n", address)
	log.Fatal(http.ListenAndServe(address, r))
}
