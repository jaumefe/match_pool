package controllers

const (
	//MAX_POINTS          = 33
	CREATE_DEFAULT_USER = `INSERT INTO users (
							name, token, role
							)
							VALUES (?, ?, ?);`
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
	SCORERS_BY_TEAM_ID_QUERY = `SELECT
								id,
								name
								FROM scorers
								WHERE team_id = ?`
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
	GET_USER_TEAMS_QUERY_NO_VALUE = `SELECT 
									teams.id,
									teams.name
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
	USERS_QUERY            = `SELECT id, name, score FROM users`
	GET_MATCH_RESULT_QUERY = `SELECT
							stage_id,
							team_home_score,
							team_away_score,
							team_home_id,
							team_away_id,
							penalty_winner_id
							FROM match
							WHERE team_home_id = ? OR team_away_id = ?`
	POINT_PER_STAGE_QUERY        = `SELECT id, points_win FROM stage`
	REGISTER_MATCH_BY_TEAM_NAMES = `WITH ids AS (
									SELECT
										(SELECT id FROM teams WHERE name = ?) AS home_id,
										(SELECT id FROM teams WHERE name = ?) AS away_id,
										(SELECT id FROM stage WHERE name = ?) AS stage_id,
										(SELECT id FROM teams WHERE name = ?) AS penalty_winner_id
									)
									INSERT INTO match (
										leg, stage_id, team_home_id, team_away_id, team_home_score, team_away_score, penalty_winner_id
									)
									SELECT
										?, stage_id, home_id, away_id, ?, ?, penalty_winner_id
									FROM ids;`
	REGISTER_MATCH_BY_ID = `INSERT INTO match (
								leg, stage_id, team_home_id, team_away_id, team_home_score, team_away_score, penalty_winner_id
							)
							VALUES (?, ?, ?, ?, ?, ?, ?);`
	REGISTER_GOAL_BY_PLAYER_NAME_AND_MATCH_ID = `WITH ids AS (
												SELECT
													(SELECT id FROM scorers WHERE name = ?) AS player_id											)
												INSERT INTO player_goals (
													player_id, match_id, goals
												)
												SELECT
													player_id, ?, ?
												FROM ids;`
	REGISTER_GOAL_BY_PLAYER_AND_MATCH_ID = `INSERT INTO player_goals (
													player_id, match_id, goals
												)
											VALUES (?, ?, ?);`
	POINTS_PER_GOAL_QUERY   = `SELECT pointsPerGoal, stage_id FROM points_goals`
	GET_POINTS_SCORER_QUERY = `SELECT
							player_goals.goals,
							match.stage_id
							FROM player_goals
							JOIN match ON player_goals.match_id = match.id
							WHERE player_goals.player_id = ?;`
	GET_ID_NAME_STAGES_QUERY = `SELECT id, name FROM stage`
	GET_MATCH_ID_QUERY       = `SELECT
								id
								FROM match
								WHERE team_home_id = ? AND team_away_id = ? AND stage_id = ?`
	SET_POOL_POSITION_TEAM = `UPDATE teams 
								SET pool_position = ?
								WHERE id = ?;`
	GET_POOL_POSITION_TEAM_ID = `SELECT pool_position
								FROM teams
								WHERE id = ?;`
	INSERT_TEAMS_FROM_CSV = `INSERT INTO teams (
							name, group_name, value, pool_group
							)
							VALUES (?, ?, ?, ?);`
	INSERT_SCORERS_FROM_CSV = `INSERT INTO scorers (
								name, team_id, position
								)
								VALUES (?, ?, ?);`
	GET_TEAM_ID_BY_NAME = `SELECT id FROM teams WHERE name = ?;`
	UPDATE_USER_NAME    = `UPDATE users
							SET name = ?
							WHERE id = ?;`
	GET_USER_NAME    = `SELECT name FROM users WHERE id = ?;`
	UPDATE_USER_PASS = `UPDATE users
						SET token = ?
						WHERE id = ?;`
	GET_STAGES_NAME_POINTS      = `SELECT id, name, points_win FROM stage;`
	SET_STAGES_WIN_POINTS_STAGE = `INSERT INTO stage (
									name, points_win
									)
									VALUES (?, ?);`
	GET_POINTS_GOAL_PER_STAGE = `SELECT
								points_goals.pointsPerGoal, stage.name
								FROM points_goals
								JOIN stage ON points_goals.stage_id = stage.id;`
	SET_POINTS_GOAL_PER_STAGE = `INSERT INTO points_goals (
									pointsPerGoal, stage_id
								)
								VALUES (?, ?);`
	GET_CONFIGURATION = `SELECT key, value FROM config;`
	SET_CONFIGURATION = `UPDATE config
						SET value = ?
						WHERE key = ?;`
	GET_MAX_POINTS = `SELECT value FROM config WHERE key = 'max_points';`
)
