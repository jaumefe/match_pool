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
	GET_USER_TEAM_ID_BY_USER_ID_QUERY = `SELECT team_id
										FROM user_teams
										WHERE user_id = ?`
	UPDATE_USER_TEAMS = `UPDATE user_teams
						SET team_id = ?
						WHERE user_id = ? AND team_id = ?`
	GET_USER_SCORERS_QUERY = `SELECT
							scorers.id,
							scorers.name
							FROM scorers
							JOIN user_scorers ON scorers.id = user_scorers.player_id
							WHERE user_scorers.user_id = ?`
	GET_USER_SCORERS_ID_BY_USER_ID_QUERY = `SELECT player_id
											FROM user_scorers
											WHERE user_id = ?`
	UPDATE_USER_SCORERS = `UPDATE user_scorers
						SET player_id = ?
						WHERE user_id = ? AND player_id = ?`
)
