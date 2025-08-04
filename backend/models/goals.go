package models

type Goal struct {
	Player  string `json:"player"`
	MatchID int    `json:"match_id"`
	Goals   int    `json:"goals"`
}

type GoalByPlayerID struct {
	PlayerID int `json:"player_id"`
	MatchID  int `json:"match_id"`
	Goals    int `json:"goals"`
}
