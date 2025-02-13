package models

type TokenResponse struct {
	Token string `json:"token"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type UserResponse struct {
	Message string `json:"message"`
	User    User   `json:"user"`
}
