package models

type Match struct {
	TieID         int    `json:"tie_id"`
	Leg           int    `json:"leg"`
	Stage         string `json:"stage"`
	TeamHome      string `json:"team_home"`
	TeamAway      string `json:"team_away"`
	TeamHomeScore int    `json:"team_home_score"`
	TeamAwayScore int    `json:"team_away_score"`
	PenaltyWinner string `json:"penalty_winner"`
}
