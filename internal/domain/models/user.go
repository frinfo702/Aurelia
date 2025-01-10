package models

type User struct {
	UserID       int    `json:"user_id"`
	UserName     string `json:"user_name"`
	UserEmail    string `json:"user_email"`
	UserPassword string `json:"user_password"`
	UserAddress  string `json:"user_address"`
}
