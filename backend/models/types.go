package models

import (
	"github.com/golang-jwt/jwt/v5"
)

type ApiResponse struct {
	Data    any            `json:"data,omitempty"`
	Error   *ErrorResponse `json:"error,omitempty"`
	Message string         `json:"message,omitempty"`
	Status  string         `json:"status"`
}

type ErrorResponse struct {
	Code    string         `json:"code"`
	Message string         `json:"message"`
	Details map[string]any `json:"details,omitempty"`
}

type HomeUser struct {
	ID    int
	Name  string
	Role  string
	Email string
	Phone string
}

type User struct {
	ID       int
	Name     string
	Role     string
	Email    string
	Phone    string
	Password string
}

type Role struct {
	ID   int
	Role string
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type SignupRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    int    `json:"phone"`
	Role_ID  int    `json:"role_id"`
}

type Claims struct {
	UserID int    `json:user_id`
	Email  string `json:email`
	jwt.RegisteredClaims
}

type BasePage struct {
	CurrentPage     string
	IsAuthenticated bool
}

type HomeData struct {
	BasePage
	User HomeUser
}

type AboutData struct {
	BasePage
	Version string
}

type LoginData struct {
	BasePage
}
