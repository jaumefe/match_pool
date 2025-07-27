package models

type Goal struct {
	Player  string `json:"player"`
	MatchID int    `json:"match_id"`
	Goals   int    `json:"goals"`
}
