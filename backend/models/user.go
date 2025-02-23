package models

type User struct {
	ID       string `json:"username"`
	Password string `json:"password"`
}