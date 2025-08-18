package models

type GoalsPoints struct {
	PointsPerGoal int `json:"points_per_goal"`
	StageID       int `json:"stage_id"`
}

type GoalsPointsStageName struct {
	PointsPerGoal int    `json:"points_per_goal"`
	Stage         string `json:"stage"`
}
