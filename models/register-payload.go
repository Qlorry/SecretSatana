package models

// LoginPayload represents the expected login form data

type RegisterPayload struct {
	Username       string `form:"username"`
	Password       string `form:"password"`
	RepeatPassword string `form:"repeat-password"`
}
