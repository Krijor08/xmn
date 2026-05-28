package handlers

import (
	"net/http"

	"example.com/go/backend/middleware"
)

func About(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, middleware.ApiResponse{
		Message: "This is the about page",
		Status:  "success",
	})
}
