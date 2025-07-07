package models

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Token string `json:"token"`
	Role  string `json:"role"`
	Score int    `json:"score"`
}
