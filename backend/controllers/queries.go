package controllers

const (
	SELECT_SCORERS_QUERY = `SELECT
							scorers.id,
							scorers.name,
							scorers.position,
							teams.name AS team_name
							FROM scorers
							JOIN teams ON scorers.team_id = teams.id`
	SELECT_SCORERS_BY_GROUP_QUERY = `SELECT 
									scorers.id,
									scorers.name,
									scorers.position,
									teams.name AS team_name
									FROM scorers
									JOIN teams ON scorers.team_id = teams.id
									WHERE teams.group_name = ?`
	SELECT_SCORERS_BY_GROUP_AND_POSITION_QUERY = `SELECT 
									scorers.id,
									scorers.name,
									scorers.position,
									teams.name AS team_name
									FROM scorers
									JOIN teams ON scorers.team_id = teams.id
									WHERE teams.group_name = ? AND scorers.position = ?`
	SELECT_TEAMS_QUERY          = `SELECT * FROM teams`
	SELECT_TEAMS_BY_GROUP_QUERY = `SELECT * FROM teams WHERE group_name = ?`
)
