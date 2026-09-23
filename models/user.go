package models

type User struct {
	ID int `json:"id"`
	Username string `json:"username"`
	Email string `json:"string"`
	Password string `json:"password"`
	Created string `json:"created"`
	Updated string `json:"updated"`
}