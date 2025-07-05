package models

type Scorer struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	TeamID   int    `json:"team_id"`
	Position string `json:"position"`
}
