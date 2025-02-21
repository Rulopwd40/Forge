package models

type TokenResponse struct {
	Token string `json:"token"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type UserResponse struct {
	Message  string `json:"message"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Level    int    `json:"level"`
}
