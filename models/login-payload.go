package models

// LoginPayload represents the expected login form data

type LoginPayload struct {
	Username string `form:"username"`
	Password string `form:"password"`
}
