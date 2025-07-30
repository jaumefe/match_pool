package models

type Stage struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	PointsWin int    `json:"points_win"`
}
