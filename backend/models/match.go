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

type MatchById struct {
	TieID           int  `json:"tie_id"`
	Leg             int  `json:"leg"`
	StageID         int  `json:"stage_id"`
	TeamHomeID      int  `json:"team_home_id"`
	TeamAwayID      int  `json:"team_away_id"`
	TeamHomeScore   int  `json:"team_home_score"`
	TeamAwayScore   int  `json:"team_away_score"`
	PenaltyWinnerID *int `json:"penalty_winner_id"`
}
