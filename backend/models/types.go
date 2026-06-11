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

type HomeUser struct {
	ID    int
	Name  string
	Role  string
	Email string
	Phone int
}

type User struct {
	ID       int
	Name     string
	Role     string
	Email    string
	Phone    int
	Password string
}

type Roles struct {
	ID   string
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

type SignupData struct {
	BasePage
	Roles []Roles
}
