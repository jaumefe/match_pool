package models

type Scorer struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Team     string `json:"team"`
	Position string `json:"position"`
}
