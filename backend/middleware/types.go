package middleware

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
