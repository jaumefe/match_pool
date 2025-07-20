package controllers

const (
	MAX_POINTS             = 33
	PRAGMA_FOREIGN_KEYS_ON = "PRAGMA foreign_keys = ON"
	SCORERS_QUERY          = `SELECT
						scorers.id,
						scorers.name,
						scorers.position,
						teams.name AS team_name
						FROM scorers
						JOIN teams ON scorers.team_id = teams.id`
	SCORERS_BY_GROUP_QUERY = `SELECT 
								scorers.id,
								scorers.name,
								scorers.position,
								teams.name AS team_name
								FROM scorers
								JOIN teams ON scorers.team_id = teams.id
								WHERE teams.group_name = ?`
	SCORERS_BY_GROUP_AND_POSITION_QUERY = `SELECT 
											scorers.id,
											scorers.name,
											scorers.position,
											teams.name AS team_name
											FROM scorers
											JOIN teams ON scorers.team_id = teams.id
											WHERE teams.group_name = ? AND scorers.position = ?`
	TEAMS_QUERY                 = `SELECT * FROM teams`
	TEAMS_BY_GROUP_QUERY        = `SELECT * FROM teams WHERE group_name = ?`
	USERS_TOKEN_BY_NAME_QUERY   = `SELECT id, token, role FROM users WHERE name = ?`
	GET_VALUES_BY_TEAM_ID_QUERY = `SELECT value FROM teams WHERE id = ?`
	GET_USER_TEAMS_QUERY        = `SELECT 
									teams.id,
									teams.name,
									teams.value
									FROM teams
									JOIN user_teams ON teams.id = user_teams.team_id
									WHERE user_teams.user_id = ?`
)
