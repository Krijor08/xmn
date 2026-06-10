package models

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

type User struct {
	ID    int
	Name  string
	Role  string
	Email string
	Phone string
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type BasePage struct {
	CurrentPage string
}

type HomeData struct {
	BasePage
	User User
}

type AboutData struct {
	BasePage
	Version string
}

type LoginData struct {
	BasePage
}
