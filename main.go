package main

import (
	"fmt"
	"net/http"
	"os"

	"example.com/go/backend/handlers"
	"example.com/go/backend/logger"
)

func main() {
	mux := http.NewServeMux()

	homeHandler := http.HandlerFunc(handlers.Home)
	wrappedHome := logger.Logger(homeHandler)

	mux.Handle("GET /", wrappedHome)

	mux.HandleFunc("GET /api/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})

	port := os.Getenv("PORT")
	fmt.Println(port)

	fmt.Printf("Server is running on http://localhost:%s", port)
	http.ListenAndServe(port, mux)
}
