package models

type Team struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	GroupName   string  `json:"group_name"`
	Value       int     `json:"value"`
	PoolPostion *int    `json:"pool_position"`
	PoolGroup   *string `json:"pool_group"`
}
